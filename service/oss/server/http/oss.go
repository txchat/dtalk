package http

import (
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/oss/model"
)

func GetOssToken(c *gin.Context) {
	resp, err := svc.AssumeRole("dtalk", model.Oss_Aliyun)
	c.Set(api.ReqResult, resp)
	c.Set(api.ReqError, err)
}

func GetHuaweiyunOssToken(c *gin.Context) {
	resp, err := svc.AssumeRole("dtalk", model.Oss_Huaweiuyn)
	c.Set(api.ReqResult, resp)
	c.Set(api.ReqError, err)
}

// Upload
// @Summary 上传文件
// @Author chy@33.cn
// @Tags oss 普通上传
// @Accept multipart/form-data
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param file formData file true "file"
// @Param data formData model.UploadRequest true "body"
// @Success 200 {object} model.GeneralResponse{data=model.UploadResponse}
// @Router	/upload [post]
func Upload(c *gin.Context) {
	req := &model.UploadRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	fileOpen, err := file.Open()
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	defer fileOpen.Close()

	req.Size = file.Size
	req.Body = fileOpen

	if file.Size > model.MaxPretartSize {
		c.Set(api.ReqError, xerror.NewError(xerror.OssFileTooBig))
		return
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.Upload(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// InitMultipartUpload
// @Summary 初始化分段上传任务
// @Author chy@33.cn
// @Tags oss 分段上传
// @Accept application/json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.InitMultipartUploadRequest true "body"
// @Success 200 {object} model.GeneralResponse{data=model.InitMultipartUploadResponse}
// @Router	/init-multipart-upload [post]
func InitMultipartUpload(c *gin.Context) {
	req := &model.InitMultipartUploadRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.InitiateMultipartUpload(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// UploadPart
// @Summary 上传段
// @Author chy@33.cn
// @Tags oss 分段上传
// @Accept multipart/form-data
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param file formData file true "aliyun, huaweiyun 除了最后一段以外，其他段的大小范围是100KB~5GB；最后段大小范围是0~5GB。minio 除了最后一段以外，其他段的大小范围是5MB~5GB；最后段大小范围是0~5GB。"
// @Param data formData model.UploadPartRequest true "body"
// @Success 200 {object} model.GeneralResponse{data=model.UploadPartResponse}
// @Router	/upload-part [post]
func UploadPart(c *gin.Context) {
	req := &model.UploadPartRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}

	fileOpen, err := file.Open()
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	defer fileOpen.Close()

	req.PartSize = file.Size
	req.Body = fileOpen

	//if file.Size < model.MinPartSize {
	//	c.Set(api.ReqError, xerror.NewError(xerror.OssFileTooSmall))
	//	return
	//}

	if file.Size > model.MaxPartSize {
		c.Set(api.ReqError, xerror.NewError(xerror.OssFileTooBig))
		return
	}

	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.UploadPart(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// CompleteMultipartUpload
// @Summary 合并段
// @Author chy@33.cn
// @Tags oss 分段上传
// @Accept application/json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.CompleteMultipartUploadRequest true "body"
// @Success 200 {object} model.GeneralResponse{data=model.CompleteMultipartUploadResponse}
// @Router	/complete-multipart-upload [post]
func CompleteMultipartUpload(c *gin.Context) {
	req := &model.CompleteMultipartUploadRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.CompleteMultipartUpload(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// AbortMultipartUpload
// @Summary 取消分段上传任务
// @Author chy@33.cn
// @Tags oss 分段上传
// @Accept application/json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.AbortMultipartUploadRequest true "body"
// @Success 200 {object} model.GeneralResponse{data=model.AbortMultipartUploadResponse}
// @Router	/abort-multipart-upload [post]
func AbortMultipartUpload(c *gin.Context) {
	req := &model.AbortMultipartUploadRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.AbortMultipartUpload(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}

// GetHost
// @Summary 获得 oss Host
// @Author chy@33.cn
// @Tags oss
// @Accept application/json
// @Param FZM-SIGNATURE	header string true "MOCK"
// @Param data body model.GetHostReq true "body"
// @Success 200 {object} model.GeneralResponse{data=model.GetHostResp}
// @Router	/get-host [post]
func GetHost(c *gin.Context) {
	req := &model.GetHostReq{}
	err := c.ShouldBind(req)
	if err != nil {
		c.Set(api.ReqError, xerror.NewError(xerror.ParamsError).SetExtMessage(err.Error()))
		return
	}
	userId, ok := c.Get(api.Address)
	if !ok {
		c.Set(api.ReqError, xerror.NewError(xerror.SignatureInvalid))
		return
	}
	req.PersonId = userId.(string)

	res, err := svc.GetHost(req)
	c.Set(api.ReqResult, res)
	c.Set(api.ReqError, err)
}
