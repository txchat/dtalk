package aliyun

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/oss"
)

var _ oss.Oss = (*Aliyun)(nil)

type Role struct {
	RoleSessionName string `json:"roleSessionName"`
	RoleArn         string `json:"roleArn"`
}

type Aliyun struct {
	cfg    *oss.Config
	client *alioss.Client
}

func New(cfg *oss.Config) *Aliyun {
	ali := &Aliyun{
		cfg:    cfg,
		client: nil,
	}

	// 创建并维护客户端
	//go ali.watchClient()
	err := ali.createClient()
	if err != nil {
		panic(err)
	}

	return ali
}

// Config 返回配置信息
func (ali *Aliyun) Config() *oss.Config {
	return ali.cfg
}

// AssumeRole 返回临时授权角色
func (ali *Aliyun) AssumeRole() (*oss.AssumeRoleResp, error) {
	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	/*
		regionId官方示例就是写死的
		https://help.aliyun.com/document_detail/184381.htm?spm=a2c4g.11186623.2.19.25eb4bceQK29Sb#concept-1955433
	*/
	client, err := sts.NewClientWithAccessKey(ali.cfg.RegionId, ali.cfg.AccessKeyId, ali.cfg.AccessKeySecret)
	if err != nil {
		return nil, errors.WithMessage(err, "sts.NewClientWithAccessKey()")
	}
	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	var role Role
	err = json.Unmarshal([]byte(ali.cfg.Role), &role)
	if err != nil {
		return nil, errors.WithMessagef(err, "json.Unmarshal role=%+v", role)
	}
	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = role.RoleArn
	request.RoleSessionName = role.RoleSessionName
	// 设置过期时间
	request.DurationSeconds = requests.NewInteger(ali.cfg.DurationSeconds)

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, errors.WithMessage(err, "client.AssumeRole()")
	}
	return &oss.AssumeRoleResp{
		RequestId: response.RequestId,
		Credentials: oss.Credentials{
			AccessKeySecret: response.Credentials.AccessKeySecret,
			Expiration:      response.Credentials.Expiration,
			AccessKeyId:     response.Credentials.AccessKeyId,
			SecurityToken:   response.Credentials.SecurityToken,
		},
		AssumedRoleUser: oss.AssumedRoleUser{
			AssumedRoleId: response.AssumedRoleUser.AssumedRoleId,
			Arn:           response.AssumedRoleUser.Arn,
		},
	}, nil
}

// Upload   上传文件
// key string 文件名, 包括路径 "dtalk/2021-07-01/1.jpg"
// body io.Reader 文件内容.
// url string 文件资源链接.
// error it's nil if no error, otherwise it's an error object.
func (ali *Aliyun) Upload(key string, body io.Reader, size int64) (url, uri string, err error) {
	bucket, err := ali.client.Bucket(ali.cfg.Bucket)
	if err != nil {
		return "", "", errors.WithMessage(err, "client.Bucket()")
	}

	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := alioss.ObjectStorageClass(alioss.StorageStandard)

	// 指定存储类型为归档存储。
	// storageType := alioss.ObjectStorageClass(alioss.StorageArchive)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := alioss.ObjectACL(alioss.ACLPublicRead)

	// 上传字符串。
	err = bucket.PutObject(key, body, storageType, objectAcl)
	if err != nil {
		return "", "", errors.WithMessage(err, "bucket.PutObject()")
	}

	return ali.parseUrl(key), ali.parseUri(key), nil
}

func (ali *Aliyun) createClient() error {
	client, err := alioss.New(
		ali.cfg.EndPoint,
		ali.cfg.AccessKeyId,
		ali.cfg.AccessKeySecret,
	)
	if err != nil {
		return errors.WithMessage(err, "alioss.New")
	}

	ali.client = client

	return nil
}

// watchClient 定时更新临时客户端
func (ali *Aliyun) watchClient() {
	// 创建一个计时器
	durationTime := time.Second * time.Duration(ali.cfg.DurationSeconds)
	timeTickerChan := time.Tick(durationTime)

	// 更新客户端
	for {
		if err := ali.createTemporaryClient(); err != nil {
			log.Error().Err(err).Msg("create ali temporary Client err")
		} else {
			log.Debug().Msg("create ali temporary Client success")
		}
		<-timeTickerChan
	}
}

// createTemporaryClient 创建临时客户端
func (ali *Aliyun) createTemporaryClient() (err error) {
	// 获得临时角色
	assumeRoleResp, err := ali.AssumeRole()
	if err != nil {
		return errors.WithMessage(err, "ali.AssumeRole()")
	}

	// 创建临时客户端
	ali.client, err = alioss.New(ali.cfg.EndPoint,
		assumeRoleResp.Credentials.AccessKeyId,
		assumeRoleResp.Credentials.AccessKeySecret,
		alioss.SecurityToken(assumeRoleResp.Credentials.SecurityToken))
	if err != nil {
		return errors.WithMessage(err, "alioss.New()")
	}

	return nil
}

// parseUrl 获取对象在阿里云OSS上的完整访问URL
func (ali *Aliyun) parseUrl(key string) string {
	return fmt.Sprintf("https://%s.%s/%s", ali.cfg.Bucket, ali.cfg.EndPoint, key)
}

// parseUri ...
func (ali *Aliyun) parseUri(key string) string {
	return fmt.Sprintf("/%s", key)
}

// GetHost ...
func (ali *Aliyun) GetHost() string {
	return fmt.Sprintf("https://%s.%s", ali.cfg.Bucket, ali.cfg.EndPoint)
}

// InitiateMultipartUpload 初始化分段上传任务
// 使用分段上传方式传输数据前，必须先通知OBS初始化一个分段上传任务。
// 该操作会返回一个OBS服务端创建的全局唯一标识（Upload ID），用于标识本次分段上传任务。
// 您可以根据这个唯一标识来发起相关的操作，如取消分段上传任务、列举分段上传任务、列举已上传的段等。
//
// key string
//
// uploadId string
// err 		error
func (ali *Aliyun) InitiateMultipartUpload(key string) (uploadId string, err error) {
	bucket, err := ali.client.Bucket(ali.cfg.Bucket)
	if err != nil {
		return "", errors.WithMessage(err, "client.Bucket()")
	}

	imur, err := bucket.InitiateMultipartUpload(key)
	if err != nil {
		return "", errors.WithMessage(err, "bucket.InitiateMultipartUpload")
	}

	return imur.UploadID, nil
}

// UploadPart 上传段
// 初始化一个分段上传任务之后，可以根据指定的对象名和Upload ID来分段上传数据。
// 每一个上传的段都有一个标识它的号码——分段号（Part Number，范围是1~10000）。
// 对于同一个Upload ID，该分段号不但唯一标识这一段数据，也标识了这段数据在整个对象内的相对位置。
// 如果您用同一个分段号上传了新的数据，那么OBS上已有的这个段号的数据将被覆盖。
// 除了最后一段以外，其他段的大小范围是100KB~5GB；最后段大小范围是0~5GB。
// 每个段不需要按顺序上传，甚至可以在不同进程、不同机器上上传，OBS会按照分段号排序组成最终对象。
//
// key 			string
// uploadId 	string
// body 		io.Reader
// partNumber 	int32
// offset 		int64
// partSize 	int64
//
// ETag	string
// err	error
func (ali *Aliyun) UploadPart(key, uploadId string, body io.Reader, partNumber int32, offset,
	partSize int64) (ETag string, err error) {
	// TODO offset, partSize 好像是不需要的参数, 如果 body 是 io.Reader 的话
	bucket, err := ali.client.Bucket(ali.cfg.Bucket)
	if err != nil {
		return "", errors.WithMessage(err, "client.Bucket")
	}

	imur := alioss.InitiateMultipartUploadResult{
		Key:      key,
		UploadID: uploadId,
		Bucket:   ali.cfg.Bucket,
	}

	part, err := bucket.UploadPart(imur, body, partSize, int(partNumber))
	if err != nil {
		return "", errors.WithMessage(err, "bucket.UploadPart")
	}
	return part.ETag, nil
}

// CompleteMultipartUpload 合并段
// 所有分段上传完成后，需要调用合并段接口来在OBS服务端生成最终对象。
// 在执行该操作时，需要提供所有有效的分段列表（包括分段号和分段ETag值）；
// OBS收到提交的分段列表后，会逐一验证每个段的有效性。当所有段验证通过后，OBS将把这些分段组合成最终的对象。
//
// key 		string
// uploadId string
// parts 	[]model.Part
//
// url string
// err error
func (ali *Aliyun) CompleteMultipartUpload(key, uploadId string, parts []oss.Part) (url, uri string, err error) {
	bucket, err := ali.client.Bucket(ali.cfg.Bucket)
	if err != nil {
		return "", "", errors.WithMessage(err, "client.Bucket")
	}

	imur := alioss.InitiateMultipartUploadResult{
		Key:      key,
		UploadID: uploadId,
		Bucket:   ali.cfg.Bucket,
	}

	ossParts := make([]alioss.UploadPart, len(parts), len(parts))
	sort.Sort(oss.Parts(parts))
	for i := range parts {
		ossParts[i] = alioss.UploadPart{
			PartNumber: int(parts[i].PartNumber),
			ETag:       parts[i].ETag,
		}
	}

	_, err = bucket.CompleteMultipartUpload(imur, ossParts)
	if err != nil {
		return "", "", errors.WithMessage(err, "bucket.CompleteMultipartUpload")
	}
	return ali.parseUrl(key), ali.parseUri(key), nil
}

// AbortMultipartUpload 取消分段上传任务
// 分段上传任务可以被取消，当一个分段上传任务被取消后，就不能再使用其Upload ID做任何操作，已经上传段也会被OBS删除。
// 采用分段上传方式上传对象过程中或上传对象失败后会在桶内产生段，这些段会占用您的存储空间，您可以通过取消该分段上传任务来清理掉不需要的段，节约存储空间。
//
// key 		string
// uploadId string
//
// err error
func (ali *Aliyun) AbortMultipartUpload(key, uploadId string) error {
	bucket, err := ali.client.Bucket(ali.cfg.Bucket)
	if err != nil {
		return errors.WithMessage(err, "client.Bucket")
	}

	imur := alioss.InitiateMultipartUploadResult{
		Key:      key,
		UploadID: uploadId,
		Bucket:   ali.cfg.Bucket,
	}

	err = bucket.AbortMultipartUpload(imur)
	if err != nil {
		return errors.WithMessage(err, "bucket.AbortMultipartUpload")
	}
	return nil
}

func (ali *Aliyun) test() {
	//ali.client.
}
