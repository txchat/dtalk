package model

import (
	"io"

	"github.com/txchat/dtalk/pkg/oss"
)

type GeneralResponse struct {
	Result  int         `json:"result"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type ossBase struct {
	// 应用 ID
	AppId string `json:"appId" form:"appId" binding:"required"`

	// 云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)
	OssType string `json:"ossType" form:"ossType"`

	// 上传者
	PersonId string `json:"-" from:"-"`
}

type Part struct {
	// 段数据的MD5值
	ETag string `json:"ETag" form:"ETag" binding:"required"`
	// 分段序号, 范围是1~10000
	PartNumber int32 `json:"partNumber" form:"partNumber" binding:"required"`
}

func (p *Part) Convert2OssPart() oss.Part {
	return oss.Part{
		ETag:       p.ETag,
		PartNumber: p.PartNumber,
	}
}

type UploadRequest struct {
	ossBase

	// 文件名(包含路径)
	Key string `json:"key" form:"key" binding:"required"`

	// 桶名
	//Bucket  string `json:"bucket" form:"bucket" binding:"required"`

	// 上传文件的人
	//Person string `json:"person" form:"person" binding:"required"`

	// 文件内容
	Body io.Reader `json:"-" form:"-"`
	Size int64     `json:"-" from:"-"`
}

type UploadResponse struct {
	// 资源链接
	Url string `json:"url"`
	Uri string `json:"uri"`
}

type InitMultipartUploadRequest struct {
	ossBase

	// 文件名(包含路径)
	Key string `json:"key" form:"key" binding:"required"`
}

type InitMultipartUploadResponse struct {
	// 分段上传任务全局唯一标识
	UploadId string `json:"uploadId"`
	// 文件名(包含路径)
	Key string `json:"key"`
}

type UploadPartRequest struct {
	ossBase

	// 文件名(包含路径)
	Key string `json:"key" form:"key" binding:"required"`
	// 分段上传任务全局唯一标识
	UploadId string `json:"uploadId" form:"uploadId" binding:"required"`
	// 分段序号, 范围是1~10000
	PartNumber int32 `json:"partNumber" form:"partNumber" binding:"required"`

	// 偏移量, 可不填
	Offset int64 `json:"-" form:"-"`

	// 分段文件大小, 可不填
	PartSize int64 `json:"-" form:"-"`

	// 文件内容
	// aliyun, huaweiyun 除了最后一段以外，其他段的大小范围是100KB~5GB；最后段大小范围是0~5GB。
	// minio 除了最后一段以外，其他段的大小范围是5MB~5GB；最后段大小范围是0~5GB。
	Body io.Reader `json:"-" form:"-"`
}

type UploadPartResponse struct {
	// 分段上传任务全局唯一标识
	UploadId string `json:"uploadId"`
	// 文件名(包含路径)
	Key string `json:"key"`
	Part
}

type CompleteMultipartUploadRequest struct {
	ossBase

	// 文件名(包含路径)
	Key string `json:"key" form:"key" binding:"required"`
	// 分段上传任务全局唯一标识
	UploadId string `json:"uploadId" form:"uploadId" binding:"required"`
	//
	Parts []Part `json:"parts" binding:"required"`
}

type CompleteMultipartUploadResponse struct {
	// 资源链接
	Url string `json:"url"`
	Uri string `json:"uri"`
}

type AbortMultipartUploadRequest struct {
	ossBase

	// 文件名(包含路径)
	Key string `json:"key" form:"key" binding:"required"`
	// 分段上传任务全局唯一标识
	UploadId string `json:"uploadId" form:"uploadId" binding:"required"`
}

type AbortMultipartUploadResponse struct {
}

type GetHostReq struct {
	ossBase
}

type GetHostResp struct {
	Host string `json:"host"`
}
