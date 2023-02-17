module github.com/txchat/dtalk

go 1.15

require (
	github.com/33cn/chain33 v1.65.3
	github.com/BurntSushi/toml v0.3.1
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.948
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/glog v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/haltingstate/secp256k1-go v0.0.0-20151224084235-572209b26df6
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.21.1+incompatible
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.35-rc
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac
	github.com/jinzhu/gorm v1.9.16
	github.com/minio/minio-go/v7 v7.0.12
	github.com/oofpgDLD/u-push v0.0.2
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.21.0
	github.com/stretchr/testify v1.8.1
	github.com/tencentyun/tls-sig-api-v2-golang v1.2.0
	github.com/txchat/im v0.1.0
	github.com/txchat/pkg v0.0.1
	github.com/zeromicro/go-zero v1.4.3
	go.etcd.io/etcd/api/v3 v3.5.5
	go.etcd.io/etcd/client/v3 v3.5.5
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
)

replace (
	github.com/txchat/im => ../im
)
