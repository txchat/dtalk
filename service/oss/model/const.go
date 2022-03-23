package model

const Oss_Aliyun = "aliyun"
const Oss_Huaweiuyn = "huaweiyun"
const Oss_Minio = "minio"

// MinPartSize - absolute minimum part size (5 MiB) below which
// a part in a multipart upload may not be uploaded.
const MinPartSize = 1024 * 1024 * 5

// MaxPartSize - maximum part size 5GiB for a single multipart upload
// operation.
const MaxPartSize = 1024 * 1024 * 1024 * 5
