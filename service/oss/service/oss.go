package service

import (
	"strings"

	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/pkg/oss"
	"github.com/txchat/dtalk/service/oss/model"
)

func (s *Service) AssumeRole(appId, ossType string) (*oss.AssumeRoleResp, error) {
	invoker, err := s.getInvoker(appId, ossType)
	if err != nil {
		return nil, err
	}

	assume, err := s.load(appId, ossType, invoker)
	if err == nil {
		return assume, nil
	}

	assume, err = invoker.AssumeRole()
	if err != nil {
		return nil, xerror.NewError(xerror.CodeInnerError)
	}
	err = s.save(appId, ossType, invoker, assume)
	if err != nil {
		s.log.Warn().Err(err).Interface("assume", assume).Msg("can not save assumeRole")
	}
	return assume, nil
}

func (s *Service) load(appId, ossType string, invoker oss.Oss) (*oss.AssumeRoleResp, error) {
	return s.dao.GetAssumeRole(appId, ossType, invoker.Config())
}

func (s *Service) save(appId, ossType string, invoker oss.Oss, data *oss.AssumeRoleResp) error {
	return s.dao.SaveAssumeRole(appId, ossType, invoker.Config(), data)
}

// Upload 上传文件
func (s *Service) Upload(req *model.UploadRequest) (res *model.UploadResponse, err error) {
	var fullErr error
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Err(fullErr).Interface("req", req).Str("PersonIs", req.PersonId).Msg("Upload")
		} else {
			s.log.Info().Interface("req", req).Interface("res", res).Str("PersonIs", req.PersonId).Msg("Upload")
		}
	}()

	if err := checkKey(req.Key); err != nil {
		return nil, err
	}

	invoker, err := s.getInvoker(req.AppId, req.OssType)
	if err != nil {
		return nil, err
	}

	url, uri, err := invoker.Upload(req.Key, req.Body, req.Size)
	if err != nil {
		fullErr = err
		return &model.UploadResponse{}, xerror.NewError(xerror.CodeInnerError)
	}
	return &model.UploadResponse{
		Url: url,
		Uri: uri,
	}, nil
}

// InitiateMultipartUpload 初始化分段上传任务
func (s *Service) InitiateMultipartUpload(req *model.InitMultipartUploadRequest) (res *model.InitMultipartUploadResponse, err error) {
	var fullErr error
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Err(fullErr).Interface("req", req).Str("PersonIs", req.PersonId).Msg("InitiateMultipartUpload")
		} else {
			s.log.Info().Interface("req", req).Interface("res", res).Str("PersonIs", req.PersonId).Msg("InitiateMultipartUpload")
		}
	}()

	if err := checkKey(req.Key); err != nil {
		return nil, err
	}

	invoker, err := s.getInvoker(req.AppId, req.OssType)
	if err != nil {
		return nil, err
	}

	uploadId, err := invoker.InitiateMultipartUpload(req.Key)
	if err != nil {
		fullErr = err
		return nil, xerror.NewError(xerror.CodeInnerError)
	}
	return &model.InitMultipartUploadResponse{
		UploadId: uploadId,
		Key:      req.Key,
	}, nil
}

// UploadPart 上传段
func (s *Service) UploadPart(req *model.UploadPartRequest) (res *model.UploadPartResponse, err error) {
	var fullErr error
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Err(fullErr).Interface("req", req).Str("PersonIs", req.PersonId).Msg("UploadPart")
		} else {
			s.log.Info().Interface("req", req).Interface("res", res).Str("PersonIs", req.PersonId).Msg("UploadPart")
		}
	}()

	if err := checkKey(req.Key); err != nil {
		return nil, err
	}

	invoker, err := s.getInvoker(req.AppId, req.OssType)
	if err != nil {
		return nil, err
	}

	ETag, err := invoker.UploadPart(req.Key, req.UploadId, req.Body, req.PartNumber, req.Offset, req.PartSize)
	if err != nil {
		fullErr = err
		return nil, xerror.NewError(xerror.CodeInnerError)
	}

	res = &model.UploadPartResponse{}
	res.UploadId = req.UploadId
	res.Key = req.Key
	res.ETag = ETag
	res.PartNumber = req.PartNumber
	return res, nil
}

// CompleteMultipartUpload 合并段
func (s *Service) CompleteMultipartUpload(req *model.CompleteMultipartUploadRequest) (res *model.CompleteMultipartUploadResponse, err error) {
	var fullErr error
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Err(fullErr).Interface("req", req).Str("PersonIs", req.PersonId).Msg("CompleteMultipartUpload")
		} else {
			s.log.Info().Interface("req", req).Interface("res", res).Str("PersonIs", req.PersonId).Msg("CompleteMultipartUpload")
		}
	}()

	if err := checkKey(req.Key); err != nil {
		return nil, err
	}

	invoker, err := s.getInvoker(req.AppId, req.OssType)
	if err != nil {
		return nil, err
	}

	ossParts := make([]oss.Part, len(req.Parts), len(req.Parts))
	for i := range ossParts {
		ossParts[i] = req.Parts[i].Convert2OssPart()
	}
	url, uri, err := invoker.CompleteMultipartUpload(req.Key, req.UploadId, ossParts)
	if err != nil {
		fullErr = err
		return nil, xerror.NewError(xerror.CodeInnerError)
	}
	return &model.CompleteMultipartUploadResponse{
		Url: url,
		Uri: uri,
	}, nil
}

// AbortMultipartUpload 取消分段上传任务
func (s *Service) AbortMultipartUpload(req *model.AbortMultipartUploadRequest) (res *model.AbortMultipartUploadResponse, err error) {
	var fullErr error
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Err(fullErr).Interface("req", req).Str("PersonIs", req.PersonId).Msg("AbortMultipartUpload")
		} else {
			s.log.Info().Interface("req", req).Interface("res", res).Str("PersonIs", req.PersonId).Msg("AbortMultipartUpload")
		}
	}()

	if err := checkKey(req.Key); err != nil {
		return nil, err
	}

	invoker, err := s.getInvoker(req.AppId, req.OssType)
	if err != nil {
		return nil, err
	}

	err = invoker.AbortMultipartUpload(req.Key, req.UploadId)
	if err != nil {
		fullErr = err
		return nil, xerror.NewError(xerror.CodeInnerError)
	}
	return &model.AbortMultipartUploadResponse{}, nil
}

func (s *Service) getInvoker(appId, ossType string) (oss.Oss, error) {
	a := s.GetEngine(appId)
	if a == nil {
		return nil, xerror.NewError(xerror.FeaturesUnSupported)
	}

	// 未选择服务商, 使用默认服务商
	if ossType == "" {
		ossType = a.DefaultOssType
	}
	invoker := a.GetInvoker(ossType)
	if invoker == nil {
		return nil, xerror.NewError(xerror.FeaturesUnSupported)
	}
	return invoker, nil
}

func (s *Service) GetHost(req *model.GetHostReq) (res *model.GetHostResp, err error) {
	var fullErr error
	defer func() {
		if err != nil {
			s.log.Error().Err(err).Err(fullErr).Interface("req", req).Str("PersonIs", req.PersonId).Msg("GetHost")
		} else {
			s.log.Info().Interface("req", req).Interface("res", res).Str("PersonIs", req.PersonId).Msg("GetHost")
		}
	}()

	invoker, err := s.getInvoker(req.AppId, req.OssType)
	if err != nil {
		return nil, err
	}

	host := invoker.GetHost()
	return &model.GetHostResp{
		Host: host,
	}, nil
}

func checkKey(key string) error {
	// key 非空
	if strings.TrimSpace(key) == "" {
		return xerror.NewError(xerror.OssKeyIllegal)
	}

	// key 不包含 ..
	if strings.Contains(key, "..") {
		return xerror.NewError(xerror.OssKeyIllegal)
	}

	return nil
}
