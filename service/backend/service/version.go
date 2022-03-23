package service

import (
	"github.com/dgrijalva/jwt-go"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backend/config"
	"github.com/txchat/dtalk/service/backend/model"
	"time"
)

func (s *Service) CreateVersion(request *model.VersionCreateRequest) (res *model.VersionCreateResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("CreateVersion", "err", err, "req", request)
		} else {
			s.log.Info("CreateVersion", "req", request)
		}
	}()

	versionForm := model.VersionForm{
		Platform:    request.Platform,
		DeviceType:  request.DeviceType,
		VersionName: request.VersionName,
		VersionCode: request.VersionCode,
		Url:         request.Url,
		Size:        request.Size,
		Md5:         request.Md5,
		Description: request.Description,
		Force:       request.Force,
		OpeUser:     request.OpeUser,
		CreateTime:  time.Now().UnixNano() / 1e6,
		UpdateTime:  time.Now().UnixNano() / 1e6,
	}
	_, id, err := s.dao.InsertVersion(&versionForm)
	if err != nil {
		return nil, err
	}
	versionForm.Id = id

	return &model.VersionCreateResponse{
		Version: versionForm,
	}, err
}

func (s *Service) UpdateVersion(request *model.VersionUpdateRequest) (res *model.VersionUpdateResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("UpdateVersion", "err", err, "req", request)
		} else {
			s.log.Info("UpdateVersion", "req", request)
		}
	}()
	versionForm := model.VersionForm{
		Id:          request.Id,
		VersionName: request.VersionName,
		VersionCode: request.VersionCode,
		Url:         request.Url,
		Description: request.Description,
		UpdateTime:  time.Now().UnixNano() / 1e6,
		Force:       request.Force,
		OpeUser:     request.OpeUser,
		Md5:         request.Md5,
		Size:        request.Size,
	}

	result, err := s.dao.UpdateVersion(&versionForm)
	if err != nil {
		return nil, err
	}
	response := &model.VersionUpdateResponse{
		Version: *result,
	}
	return response, err
}

func (s *Service) ChangeVersionStatus(request *model.VersionChangeStatusRequest) (err error) {
	defer func() {
		if err != nil {
			s.log.Error("ChangeVersionStatus", "err", err, "req", request)
		} else {
			s.log.Info("ChangeVersionStatus", "req", request)
		}
	}()
	versionForm := model.VersionForm{
		Id:         request.Id,
		UpdateTime: time.Now().UnixNano() / 1e6,
		OpeUser:    request.OpeUser,
	}
	err = s.dao.ChangeVersionStatus(&versionForm)
	return err
}

func (s *Service) GetVersionList(request *model.GetVersionListRequest) (res *model.GetVersionListResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("GetVersionList", "err", err, "req", request)
		} else {
			s.log.Info("GetVersionList", "req", request)
		}
	}()
	versionForm := model.VersionForm{
		Platform:   request.Platform,
		DeviceType: request.DeviceType,
	}
	versionList, totalElements, totalPages, err := s.dao.GetVersionList(&versionForm, request.Page, model.Size)
	if err != nil {
		return nil, err
	}
	response := &model.GetVersionListResponse{
		TotalElements: totalElements,
		TotalPages:    totalPages,
		VersionList:   *versionList,
	}
	return response, err
}

func (s *Service) CheckAndUpdateVersion(request *model.VersionCheckAndUpdateRequest) (res *model.VersionCheckAndUpdateResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("CheckAndUpdateVersion", "err", err, "req", request)
		} else {
			s.log.Info("CheckAndUpdateVersion", "req", request)
		}
	}()
	versionForm := model.VersionForm{
		VersionCode: request.VersionCode,
		DeviceType:  request.DeviceType,
		Platform:    s.Platform,
	}
	result, err := s.dao.CheckAndUpdateVersion(&versionForm)
	if err != nil {
		return nil, err
	}
	response := &model.VersionCheckAndUpdateResponse{}

	response.Id = result.Id
	response.Platform = result.Platform
	response.Status = result.Status
	response.DeviceType = result.DeviceType
	response.VersionName = result.VersionName
	response.VersionCode = result.VersionCode
	response.Url = result.Url
	response.Force = result.Force
	response.Description = result.Description
	response.OpeUser = result.OpeUser
	response.Md5 = result.Md5
	response.Size = result.Size
	response.UpdateTime = result.UpdateTime
	response.CreateTime = result.CreateTime

	return response, err
}

func (s *Service) GetToken(request *model.GetTokenRequest) (res *model.GetTokenResponse, err error) {
	defer func() {
		if err != nil {
			s.log.Error("GetToken", "err", err, "req", request)
		} else {
			s.log.Info("GetToken", "req", request)
		}
	}()
	userInfo := model.UserInfo{
		UserName: request.UserName,
		Password: request.Password,
	}
	if userInfo.UserName != config.Conf.Release.UserName || userInfo.Password != config.Conf.Release.Password {
		return nil, xerror.NewError(xerror.ParamsError).SetExtMessage("用户名或密码错误")
	}
	token, err := s.createToken(userInfo.UserName)
	if err != nil {
		return nil, err
	}

	return &model.GetTokenResponse{
		UserInfo: model.UserInfoResponse{
			UserName: userInfo.UserName,
			Token:    token,
		},
	}, nil
}

func (s *Service) createToken(username string) (string, error) {
	c := model.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.Conf.Release.TokenExpireDuration).Unix(),
			Issuer:    config.Conf.Release.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Conf.Release.Key))
}
