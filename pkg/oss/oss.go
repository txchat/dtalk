package oss

import (
	"io"
)

type Oss interface {
	Config() *Config
	AssumeRole() (*AssumeRoleResp, error)
	Upload(string, io.Reader, int64) (url, uri string, err error)
	InitiateMultipartUpload(key string) (uploadId string, err error)
	UploadPart(key, uploadId string, body io.Reader, partNumber int32, offset, partSize int64) (ETag string, err error)
	CompleteMultipartUpload(key, uploadId string, parts []Part) (url, uri string, err error)
	AbortMultipartUpload(key, uploadId string) error
	GetHost() string
}

type Config struct {
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	Role            string
	Policy          string
	DurationSeconds int
	Bucket          string
	EndPoint        string
	PublicUrl       string
}

type Part struct {
	// 段数据的MD5值
	ETag string `json:"ETag" form:"ETag" binding:"required"`
	// 分段序号, 范围是1~10000
	PartNumber int32 `json:"partNumber" form:"partNumber" binding:"required"`
}

// 结构体数组
type Parts []Part

// 下面的三个函数必须实现（获取长度函数，交换函数，比较函数）
func (p Parts) Len() int {
	return len(p)
}
func (p Parts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Parts) Less(i, j int) bool {
	return p[i].PartNumber < p[j].PartNumber
}

//{
//"Credentials": {
//"AccessKeyId": "STS.L4aBSCSJVMuKg5U1****",
//"AccessKeySecret": "wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pK****",
//"Expiration": "2015-04-09T11:52:19Z",
//"SecurityToken": "********"
//},
//"AssumedRoleUser": {
//"Arn": "acs:ram::123456789012****:role/adminrole/alice",
//"AssumedRoleId":"34458433936495****:alice"
//},
//"RequestId": "6894B13B-6D71-4EF5-88FA-F32781734A7F"
//}

type AssumeRoleResp struct {
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Credentials     Credentials     `json:"Credentials" xml:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

// Credentials is a nested struct in sts response
type Credentials struct {
	AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	Expiration      string `json:"Expiration" xml:"Expiration"`
	AccessKeyId     string `json:"AccessKeyId" xml:"AccessKeyId"`
	SecurityToken   string `json:"SecurityToken" xml:"SecurityToken"`
}

// AssumedRoleUser is a nested struct in sts response
type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId" xml:"AssumedRoleId"`
	Arn           string `json:"Arn" xml:"Arn"`
}
