package service

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	publicUtil "github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/auth/model"
	"gopkg.in/yaml.v2"
)

func (s *Service) Auth(request *model.AuthRequest) (res *model.AuthResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("Auth", "err", err, "req", request)
		} else {
			s.log.Info("Auth", "req", request)
		}
	}()

	//解析请求中的token
	digest, createTime, token, err := s.ParseTokenInRequest(&request.Token)

	if err != nil {
		return &model.AuthResponse{}, model.ErrInvalidToken
	}

	if !time.Now().Before(time.Unix(0, createTime).Add(s.cfg.Key.KeyExpireDuration)) {
		return &model.AuthResponse{}, model.ErrInvalidRequest
	}

	authInfo := model.AuthInfo{
		AppId:      request.AppId,
		Token:      token,
		Digest:     digest,
		CreateTime: createTime,
	}

	//获取key
	key, err := s.dao.GetKey(authInfo.AppId)
	if err != nil {
		return &model.AuthResponse{}, model.ErrQueryKey
	}
	if key == "" && err == nil {
		return &model.AuthResponse{}, model.ErrKeyIsNonexistent
	}

	//检验digest
	result, err := s.checkDigest(&authInfo, key)
	if err != nil {
		return &model.AuthResponse{}, model.ErrCheckDigest
	}
	if result == false && err == nil {
		return &model.AuthResponse{}, model.ErrInvalidRequest
	}

	//加载配置
	configString, err := s.dao.LoadConfig(authInfo.AppId)
	if err != nil {
		return &model.AuthResponse{}, model.ErrLoadConfig
	}
	if configString == "" && err == nil {
		return &model.AuthResponse{}, model.ErrConfigurationIsNonexistent
	}

	//分析配置
	appConfig, err := s.ParseConfigString(configString)
	if err != nil {
		return &model.AuthResponse{}, model.ErrParseConfig
	}

	//发送请求
	var bodyOfResponse []byte
	bodyOfResponse, err = s.issueHTTPCommunication(appConfig.BasicConfig.Method, appConfig.BasicConfig.Url, appConfig.Request.TokenName, authInfo.Token)
	if err != nil {
		return &model.AuthResponse{}, model.ErrHTTPCommunication
	}
	dataOfResponse := make(map[string]interface{})
	err = json.Unmarshal(bodyOfResponse, &dataOfResponse)
	if err != nil {
		return &model.AuthResponse{}, model.ErrUnmarshalBodyOfResponse
	}

	//比较返回结果
	successValue := appConfig.Response.SuccessResult
	result, err = s.CompareValues(successValue, dataOfResponse)
	if err != nil {
		return &model.AuthResponse{}, err
	}
	if result != true {
		return &model.AuthResponse{}, model.ErrAuthFiled
	}

	//获取uid
	uidName := appConfig.Response.UidName
	uid, found := s.FindUid(uidName, dataOfResponse)
	if found {
		return &model.AuthResponse{
			Uid: uid,
		}, nil
	}

	return &model.AuthResponse{}, model.ErrFindUid
}

func (s *Service) ParseConfigString(configString string) (*model.AppConfig, error) {
	appConfig := model.AppConfig{}
	err := yaml.Unmarshal([]byte(configString), &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}

func (s *Service) issueHTTPCommunication(method string, url string, tokenName string, token string) ([]byte, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set(tokenName, token)

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyOfResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bodyOfResponse, nil
}

func (s *Service) CompareValues(srcValue interface{}, destValue interface{}) (bool, error) {
	result := false
	//判断srcValue是否是映射
	if subValueOfSrcValue, ok := srcValue.(map[interface{}]interface{}); ok {
		//判断destValue与srcValue格式是否一致
		subValueOfDestValue, ok := destValue.(map[string]interface{})
		if !ok {
			result = false
			return result, model.ErrWrongDataFormat
		}
		//遍历映射，对值进行比较
		for name, grandValueOfSrcValue := range subValueOfSrcValue {
			stringOfName := publicUtil.MustToString(name)
			//判断是否存在该键值对
			grandValueOfDestValue, exist := subValueOfDestValue[stringOfName]
			if !exist {
				result = false
				return result, model.ErrFieldIsNonexistentInCompareValues
			}
			//对值进行递归比较
			result, err := s.CompareValues(grandValueOfSrcValue, grandValueOfDestValue)
			if result == false {
				return result, err
			}
		}
		result = true
		return result, nil
	}
	//将值转换成字符串比较
	srcValueString := publicUtil.MustToString(srcValue)
	destValueString := publicUtil.MustToString(destValue)

	if srcValueString == destValueString {
		result = true
	}

	return result, nil
}

func (s *Service) FindUid(uidName string, result interface{}) (string, bool) {
	uid := ""
	found := false
	//递归查找uid
	if subResult, ok := result.(map[string]interface{}); ok {
		for name, value := range subResult {
			if grandResult, ok := value.(map[string]interface{}); ok {
				uid, found = s.FindUid(uidName, grandResult)
				if found {
					return uid, found
				}
			} else if name == uidName {
				uid = publicUtil.MustToString(value)
				found = true
				return uid, found
			}
		}
	}

	return uid, found
}

func (s *Service) checkDigest(authInfo *model.AuthInfo, key string) (bool, error) {
	str := authInfo.AppId + authInfo.Token + publicUtil.MustToString(authInfo.CreateTime) + key
	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return false, err
	}

	src := hash.Sum(nil)
	//转化为十六进制字符串
	digest := hex.EncodeToString(src)

	if digest != authInfo.Digest {
		return false, nil
	}

	return true, nil
}

func (s *Service) ParseTokenInRequest(tokenInRequest *string) (digest string, createTime int64, token string, err error) {
	stringArray := strings.SplitN(*tokenInRequest, "$", 3)
	if len(stringArray) < 3 {
		return "", 0, "", model.ErrInvalidToken
	}
	digest = stringArray[0]
	createTime = publicUtil.MustToInt64(stringArray[1])
	token = stringArray[2]
	return digest, createTime, token, nil
}
