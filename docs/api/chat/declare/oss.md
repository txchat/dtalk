### 1. "获取token"

1. route definition

- Url: /oss/get-token
- Method: POST
- Request: `GetTokenReq`
- Response: `GetTokenResp`

2. request definition



```golang
type GetTokenReq struct {
}
```


3. response definition



```golang
type GetTokenResp struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	Credentials Credentials `json:"Credentials" xml:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

type Credentials struct {
	AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	Expiration string `json:"Expiration" xml:"Expiration"`
	AccessKeyId string `json:"AccessKeyId" xml:"AccessKeyId"`
	SecurityToken string `json:"SecurityToken" xml:"SecurityToken"`
}

type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId" xml:"AssumedRoleId"`
	Arn string `json:"Arn" xml:"Arn"`
}
```

### 2. "获取华为云token"

1. route definition

- Url: /oss/get-huaweiyun-token
- Method: POST
- Request: `GetHWCloudTokenReq`
- Response: `GetHWCloudTokenResp`

2. request definition



```golang
type GetHWCloudTokenReq struct {
}
```


3. response definition



```golang
type GetHWCloudTokenResp struct {
	RequestId string `json:"RequestId" xml:"RequestId"`
	Credentials Credentials `json:"Credentials" xml:"Credentials"`
	AssumedRoleUser AssumedRoleUser `json:"AssumedRoleUser" xml:"AssumedRoleUser"`
}

type Credentials struct {
	AccessKeySecret string `json:"AccessKeySecret" xml:"AccessKeySecret"`
	Expiration string `json:"Expiration" xml:"Expiration"`
	AccessKeyId string `json:"AccessKeyId" xml:"AccessKeyId"`
	SecurityToken string `json:"SecurityToken" xml:"SecurityToken"`
}

type AssumedRoleUser struct {
	AssumedRoleId string `json:"AssumedRoleId" xml:"AssumedRoleId"`
	Arn string `json:"Arn" xml:"Arn"`
}
```

### 3. "上传"

1. route definition

- Url: /oss/upload
- Method: POST
- Request: `UploadReq`
- Response: `UploadResp`

2. request definition



```golang
type UploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type UploadResp struct {
	Url string `json:"url"`
	Uri string `json:"uri"`
}
```

### 4. "初始化分段上传"

1. route definition

- Url: /oss/init-multipart-upload
- Method: POST
- Request: `InitMultiUploadReq`
- Response: `InitMultiUploadResp`

2. request definition



```golang
type InitMultiUploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type InitMultiUploadResp struct {
	UploadId string `json:"uploadId"` // 分段上传任务全局唯一标识
	Key string `json:"key"` // 文件名(包含路径)
}
```

### 5. "上传某一段"

1. route definition

- Url: /oss/upload-part
- Method: POST
- Request: `UploadPartReq`
- Response: `UploadPartResp`

2. request definition



```golang
type UploadPartReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
	UploadId string `json:"uploadId" form:"uploadId"` // 分段上传任务全局唯一标识
	PartNumber int32 `json:"partNumber,range=[1:10000]" form:"partNumber,range=[1:10000]"` // 分段序号, 范围是1~10000
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type UploadPartResp struct {
	ETag string `json:"ETag" form:"ETag"` // 段数据的MD5值
	PartNumber int32 `json:"partNumber,range=[1:10000]" form:"partNumber,range=[1:10000]"` //分段序号, 范围是1~10000
	UploadId string `json:"uploadId"` // 分段上传任务全局唯一标识
	Key string `json:"key"` // 文件名(包含路径)
}

type Part struct {
	ETag string `json:"ETag" form:"ETag"` // 段数据的MD5值
	PartNumber int32 `json:"partNumber,range=[1:10000]" form:"partNumber,range=[1:10000]"` //分段序号, 范围是1~10000
}
```

### 6. "完成分段上传"

1. route definition

- Url: /oss/complete-multipart-upload
- Method: POST
- Request: `CompleteMultiUploadReq`
- Response: `CompleteMultiUploadResp`

2. request definition



```golang
type CompleteMultiUploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
	UploadId string `json:"uploadId" form:"uploadId"` // 分段上传任务全局唯一标识
	Parts []Part `json:"parts"`
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type CompleteMultiUploadResp struct {
	Url string `json:"url"`
	Uri string `json:"uri"`
}
```

### 7. "终止分段上传"

1. route definition

- Url: /oss/abort-multipart-upload
- Method: POST
- Request: `AbortMultiUploadReq`
- Response: `AbortMultiUploadResp`

2. request definition



```golang
type AbortMultiUploadReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
	Key string `json:"key" form:"key"` // 文件名(包含路径)
	UploadId string `json:"uploadId" form:"uploadId"` // 分段上传任务全局唯一标识
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type AbortMultiUploadResp struct {
}
```

### 8. "获取主机地址"

1. route definition

- Url: /oss/get-host
- Method: POST
- Request: `GetHostReq`
- Response: `GetHostResp`

2. request definition



```golang
type GetHostReq struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}

type OssBase struct {
	AppId string `json:"appId" form:"appId"`
	OssType string `json:"ossType,optional,options=huaweiyun|aliyun|minio" form:"ossType,optional,options=huaweiyun|aliyun|minio"` //云服务商（选填）: 自动选择默认服务商, 或者指定
}
```


3. response definition



```golang
type GetHostResp struct {
	Host string `json:"host"`
}
```

