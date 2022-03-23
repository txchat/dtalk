package huaweiyun

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iamModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/txchat/dtalk/pkg/oss"
)

var _ oss.Oss = (*Huaweiyun)(nil)

type Policy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Action []string `json:"Action"`
	Effect string   `json:"Effect"`
}

type Huaweiyun struct {
	cfg    *oss.Config
	client *obs.ObsClient
}

// New 创建一个 Huaweiyun 对象
func New(cfg *oss.Config) *Huaweiyun {
	huawei := &Huaweiyun{
		cfg:    cfg,
		client: nil,
	}

	err := huawei.createClient()
	if err != nil {
		panic(err)
	}

	return huawei
}

// Config 返回 Huaweiyun Config 信息
func (huawei *Huaweiyun) Config() *oss.Config {
	return huawei.cfg
}

// AssumeRole 创建临时授权角色
func (huawei *Huaweiyun) AssumeRole() (*oss.AssumeRoleResp, error) {
	ak := huawei.cfg.AccessKeyId
	sk := huawei.cfg.AccessKeySecret

	//构建一个华为云客户端, 用于发起请求。
	//构建华为云客户端时，需要设置AccessKey ID, AccessKey Secret和区域代码。
	auth := global.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(region.ValueOf(huawei.cfg.RegionId)).
			WithCredential(auth).
			Build())

	// 构建请求对象
	request := &iamModel.CreateTemporaryAccessKeyByTokenRequest{}

	// 设置参数, 具体参考https://support.huaweicloud.com/api-iam/iam_04_0002.html
	// 设置 policy 策略
	var policy Policy
	err := json.Unmarshal([]byte(huawei.cfg.Policy), &policy) //第二个参数要地址传递
	if err != nil {
		return nil, err
	}

	var listStatementPolicy []iamModel.ServiceStatement
	for _, statement := range policy.Statement {
		var listActionStatement []string
		for _, action := range statement.Action {
			listActionStatement = append(listActionStatement, action)
		}

		var serviceStatement iamModel.ServiceStatement
		serviceStatement.Action = listActionStatement
		if statement.Effect == "Allow" {
			serviceStatement.Effect = iamModel.GetServiceStatementEffectEnum().ALLOW
		} else {
			serviceStatement.Effect = iamModel.GetServiceStatementEffectEnum().DENY
		}
		listStatementPolicy = append(listStatementPolicy, serviceStatement)
	}

	policyIdentity := &iamModel.ServicePolicy{
		Version:   policy.Version,
		Statement: listStatementPolicy,
	}

	// 设置有效时间
	durationSecondsTokenIdentityToken := int32(huawei.cfg.DurationSeconds)
	tokenIdentity := &iamModel.IdentityToken{
		DurationSeconds: &durationSecondsTokenIdentityToken,
	}

	// 设置Token
	var listMethodsIdentity = []iamModel.TokenAuthIdentityMethods{
		iamModel.GetTokenAuthIdentityMethodsEnum().TOKEN,
	}

	// 认证
	identityAuth := &iamModel.TokenAuthIdentity{
		Methods: listMethodsIdentity,
		Token:   tokenIdentity,
		Policy:  policyIdentity,
	}
	authbody := &iamModel.TokenAuth{
		Identity: identityAuth,
	}

	request.Body = &iamModel.CreateTemporaryAccessKeyByTokenRequestBody{
		Auth: authbody,
	}

	//发起请求，并得到响应。
	response, err := client.CreateTemporaryAccessKeyByToken(request)
	if err != nil {
		return nil, err
	}

	return &oss.AssumeRoleResp{
		RequestId: "", /*response.RequestId*/
		Credentials: oss.Credentials{
			AccessKeySecret: response.Credential.Secret,
			Expiration:      response.Credential.ExpiresAt,
			AccessKeyId:     response.Credential.Access,
			SecurityToken:   response.Credential.Securitytoken,
		},
		AssumedRoleUser: oss.AssumedRoleUser{
			AssumedRoleId: "", /*response.AssumedRoleUser.AssumedRoleId*/
			Arn:           "", /*response.AssumedRoleUser.Arn*/
		},
	}, err

}

// Upload 上传文件
func (huawei *Huaweiyun) Upload(key string, body io.Reader, size int64) (url, uri string, err error) {
	input := &obs.PutObjectInput{}
	input.Bucket = huawei.cfg.Bucket
	input.Key = key
	input.Body = body
	input.ContentLength = size
	_, err = huawei.client.PutObject(input)
	if err != nil {
		return "", "", errors.WithMessage(err, "huawei.client.PutObject()")
	}
	return huawei.parseUrl(key), huawei.parseUri(key), nil
}

// watchClient 定时更新临时客户端
func (huawei *Huaweiyun) watchClient() {
	// 创建一个计时器
	durationTime := time.Second * time.Duration(huawei.cfg.DurationSeconds)
	timeTickerChan := time.Tick(durationTime)

	// 更新临时客户端
	for {
		if err := huawei.createTemporaryClient(); err != nil {
			log.Error().Err(err).Msg("create huawei Client err")
		} else {
			log.Debug().Msg("create huawei Client success")
		}
		<-timeTickerChan
	}
}

// createTemporaryClient 创建临时客户端
func (huawei *Huaweiyun) createTemporaryClient() error {
	assumeRoleResp, err := huawei.AssumeRole()
	if err != nil {
		return errors.WithMessage(err, "huawei.AssumeRole()")
	}
	// 创建临时客户端
	huawei.client, err = obs.New(
		assumeRoleResp.Credentials.AccessKeyId,
		assumeRoleResp.Credentials.AccessKeySecret,
		huawei.cfg.EndPoint,
		obs.WithSecurityToken(assumeRoleResp.Credentials.SecurityToken))
	if err != nil {
		return errors.WithMessage(err, "huawei.obs.New()")
	}
	return nil
}

// createClient 创建永久客户端
func (huawei *Huaweiyun) createClient() (err error) {
	// 创建临时客户端
	huawei.client, err = obs.New(
		huawei.cfg.AccessKeyId,
		huawei.cfg.AccessKeySecret,
		huawei.cfg.EndPoint,
		obs.WithMaxRetryCount(5),
	)
	if err != nil {
		return errors.WithMessage(err, "huawei.obs.New()")
	}
	return nil
}

// parseUrl 获取对象在 Huaweiyun 上的完整访问URL
func (huawei *Huaweiyun) parseUrl(key string) string {
	return fmt.Sprintf("https://%s.%s/%s", huawei.cfg.Bucket, huawei.cfg.EndPoint, key)
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
func (huawei *Huaweiyun) InitiateMultipartUpload(key string) (uploadId string, err error) {
	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = huawei.cfg.Bucket
	input.Key = key

	output, err := huawei.client.InitiateMultipartUpload(input)
	if err != nil {
		return "", errors.WithMessage(err, "huawei.client.InitiateMultipartUpload")
	}
	return output.UploadId, nil
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
func (huawei *Huaweiyun) UploadPart(key, uploadId string, body io.Reader, partNumber int32, offset,
	partSize int64) (ETag string, err error) {
	// TODO offset, partSize 好像是不需要的参数, 如果 body 是 io.Reader 的话
	input := &obs.UploadPartInput{}
	input.Bucket = huawei.cfg.Bucket
	input.Key = key
	input.UploadId = uploadId
	input.Body = body
	input.PartNumber = int(partNumber)
	input.Offset = offset
	input.PartSize = partSize

	output, err := huawei.client.UploadPart(input)
	if err != nil {
		return "", errors.WithMessage(err, "huawei.client.UploadPart")
	}
	return output.ETag, nil
}

// CompleteMultipartUpload 合并段
// 所有分段上传完成后，需要调用合并段接口来在OBS服务端生成最终对象。
// 在执行该操作时，需要提供所有有效的分段列表（包括分段号和分段ETag值）；
// OBS收到提交的分段列表后，会逐一验证每个段的有效性。当所有段验证通过后，OBS将把这些分段组合成最终的对象。
//
// key 		string
// uploadId string
// parts 	[]oss.Part
//
// url string
// err error
func (huawei *Huaweiyun) CompleteMultipartUpload(key, uploadId string, parts []oss.Part) (url, uri string, err error) {
	input := &obs.CompleteMultipartUploadInput{}
	input.Bucket = huawei.cfg.Bucket
	input.Key = key
	input.UploadId = uploadId
	obsParts := make([]obs.Part, len(parts), len(parts))
	sort.Sort(oss.Parts(parts))
	for i := range parts {
		obsParts[i] = obs.Part{
			PartNumber: int(parts[i].PartNumber),
			ETag:       parts[i].ETag,
		}
	}
	input.Parts = obsParts

	_, err = huawei.client.CompleteMultipartUpload(input)
	if err != nil {
		return "", "", errors.WithMessage(err, "huawei.client.CompleteMultipartUpload")
	}
	return huawei.parseUrl(key), huawei.parseUri(key), nil
}

// AbortMultipartUpload 取消分段上传任务
// 分段上传任务可以被取消，当一个分段上传任务被取消后，就不能再使用其Upload ID做任何操作，已经上传段也会被OBS删除。
// 采用分段上传方式上传对象过程中或上传对象失败后会在桶内产生段，这些段会占用您的存储空间，您可以通过取消该分段上传任务来清理掉不需要的段，节约存储空间。
//
// key 		string
// uploadId string
//
// err error
func (huawei *Huaweiyun) AbortMultipartUpload(key, uploadId string) error {
	input := &obs.AbortMultipartUploadInput{}
	input.Bucket = huawei.cfg.Bucket
	input.Key = key
	input.UploadId = uploadId

	_, err := huawei.client.AbortMultipartUpload(input)
	if err != nil {
		return errors.WithMessage(err, "huawei.client.AbortMultipartUpload")
	}
	return nil
}

func (huawei *Huaweiyun) test() {
	input := &obs.InitiateMultipartUploadInput{}
	input.Bucket = huawei.cfg.Bucket
	input.Key = "objectname"
	input.ContentType = "text/plain"
	//input.Metadata = map[string]string{"property1": "property-value1", "property2": "property-value2"}
	output, err := huawei.client.InitiateMultipartUpload(input)
	if err == nil {
		fmt.Printf("UploadId:%s\n", output.UploadId)
	} else if obsError, ok := err.(obs.ObsError); ok {
		fmt.Printf("Code:%s\n", obsError.Code)
		fmt.Printf("Message:%s\n", obsError.Message)
	}
}

// parseUri ...
func (huawei *Huaweiyun) parseUri(key string) string {
	return fmt.Sprintf("/%s", key)
}

// GetHost ...
func (huawei *Huaweiyun) GetHost() string {
	return fmt.Sprintf("https://%s.%s", huawei.cfg.Bucket, huawei.cfg.EndPoint)
}
