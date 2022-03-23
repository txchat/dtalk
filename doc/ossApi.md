# 对象存储服务
### 获取token
URL: /oss/get-token

`get`|`post`

**请求参数：**  
无

**返回参数：**  

```json
{
    "result": 0,
    "message": "操作成功",
    "data": {
        "RequestId": "E7A855E3-0A3B-479A-BF22-416A2D5D27B1",
        "Credentials": {
            "AccessKeySecret": "BxbYXWGjKrVT4qSzi6ZoLtD1qjZLAfNaULreCMb7EsYm",
            "Expiration": "2021-02-24T07:48:21Z",
            "AccessKeyId": "STS.NUrf45xSfTX7EkJxoo3G7C9Cm",
            "SecurityToken": "CAIS8wF1q6Ft5B2yfSjIr5bHLY6BlYxH45rcR037nG86P8gbrPzojzz2IHhNfXlqBeket/g1n2hY5/oelqN1TIVATEiBZNNotm32E/M0Jdivgde8yJBZor/HcDHhJnyW9cvWZPqDP7G5U/yxalfCuzZuyL/hD1uLVECkNpv74vwOLK5gPG+CYCFBGc1dKyZ7tcYeLgGxD/u2NQPwiWeiZygB+CgE0DkhtfTvkp3Ht0OP1wOhk9V4/dqhfsKWCOB3J4p6XtuP2+h7S7HMyiY46WIRrP4v3PYYoG+e4ovGWwkAv0mcUezT6NhmKEpwfrAq7DDiSR5ZYpcagAEcYDFBCF3QY2aQ8Jcs9L4utLb7Z6JdPvyG/zF7l7P1Br0MX0ad6sgWEGWTvTuxemRtmKSWxMEivRzx+6aT515Y7OQdFyX1uLrR1YOq9R2tjysh2d54wu09qr7BhPavADSFHj8fCrZugoul6vHtoslRgL+KSpuI3buhZ08VDxh5rg=="
        },
        "AssumedRoleUser": {
            "AssumedRoleId": "301182410142308944:normal-app",
            "Arn": "acs:ram::1264888835193631:role/normal-app/normal-app"
        }
    }
}
```

### 获取华为云token
URL: /oss/get-huaweiyun-token

`get`|`post`

**请求参数：**  
无

**返回参数：**  

```json
{
     "result": 0,
     "message": "操作成功",
     "data": {
         "RequestId": "",
         "Credentials": {
             "AccessKeySecret": "BxbYXWGjKrVT4qSzi6ZoLtD1qjZLAfNaULreCMb7EsYm",
             "Expiration": "2021-02-24T07:48:21Z",
             "AccessKeyId": "NUrf45xSfTX7EkJxoo3G7C9Cm",
             "SecurityToken": "CAIS8wF1q6Ft5B2yfSjIr5bHLY6BlYxH45rcR037nG86P8gbrPzojzz2IHhNfXlqBeket/g1n2hY5/oelqN1TIVATEiBZNNotm32E/M0Jdivgde8yJBZor/HcDHhJnyW9cvWZPqDP7G5U/yxalfCuzZuyL/hD1uLVECkNpv74vwOLK5gPG+CYCFBGc1dKyZ7tcYeLgGxD/u2NQPwiWeiZygB+CgE0DkhtfTvkp3Ht0OP1wOhk9V4/dqhfsKWCOB3J4p6XtuP2+h7S7HMyiY46WIRrP4v3PYYoG+e4ovGWwkAv0mcUezT6NhmKEpwfrAq7DDiSR5ZYpcagAEcYDFBCF3QY2aQ8Jcs9L4utLb7Z6JdPvyG/zF7l7P1Br0MX0ad6sgWEGWTvTuxemRtmKSWxMEivRzx+6aT515Y7OQdFyX1uLrR1YOq9R2tjysh2d54wu09qr7BhPavADSFHj8fCrZugoul6vHtoslRgL+KSpuI3buhZ08VDxh5rg=="
         },
         "AssumedRoleUser": {
             "AssumedRoleId": "",
             "Arn": ""
         }
     }
 }
```