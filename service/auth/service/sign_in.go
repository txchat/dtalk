package service

import (
	"crypto/sha256"
	"encoding/hex"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/auth/model"
	"io/ioutil"
	"math/rand"
	"time"
)

func (s *Service) SignIn(request *model.SignInRequest) (res *model.SignInResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("SignIn", "err", err, "req", request)
		} else {
			s.log.Info("SignIn", "req", request)
		}
	}()

	signInInfo := model.SignInInfo{
		AppId:      request.AppId,
		ConfigFile: request.ConfigFile,
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6,
	}

	//读取配置文件内容
	file, err := signInInfo.ConfigFile.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	configString := string(buf)
	//生成key
	key, err := s.createKey(signInInfo.AppId)
	if err != nil {
		return nil, err
	}

	num, err := s.dao.SignIn(signInInfo.AppId, &configString, key, signInInfo.CreateTime, signInInfo.UpdateTime)
	if err != nil {
		return nil, err
	}

	if num == 0 {
		return nil, xerror.NewError(xerror.ParamsError).SetExtMessage("AppId已存在")
	}

	return &model.SignInResponse{
		AppId:      signInInfo.AppId,
		Key:        key,
		CreateTime: signInInfo.CreateTime,
		UpdateTime: signInInfo.UpdateTime,
	}, nil
}

func (s *Service) createKey(appId string) (string, error) {
	rand.Seed(time.Now().Unix())
	salt := ""
	hash := sha256.New()

	for i := 0; i < model.Length; i++ {
		salt = salt + string(model.Chars[rand.Intn(len(model.Chars))])
	}
	finalString := appId + salt
	_, err := hash.Write([]byte(finalString))
	if err != nil {
		return "", err
	}

	srcKeyBytes := hash.Sum(nil)
	//转化为十六进制字符串
	dstKey := hex.EncodeToString(srcKeyBytes)

	return dstKey, nil
}
