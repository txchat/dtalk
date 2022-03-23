// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/abort-multipart-upload": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "oss 分段上传"
                ],
                "summary": "取消分段上传任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MOCK",
                        "name": "FZM-SIGNATURE",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AbortMultipartUploadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.AbortMultipartUploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/complete-multipart-upload": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "oss 分段上传"
                ],
                "summary": "合并段",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MOCK",
                        "name": "FZM-SIGNATURE",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CompleteMultipartUploadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.CompleteMultipartUploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/get-host": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "oss"
                ],
                "summary": "获得 oss Host",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MOCK",
                        "name": "FZM-SIGNATURE",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetHostReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.GetHostResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/init-multipart-upload": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "oss 分段上传"
                ],
                "summary": "初始化分段上传任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MOCK",
                        "name": "FZM-SIGNATURE",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InitMultipartUploadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.InitMultipartUploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "oss 普通上传"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MOCK",
                        "name": "FZM-SIGNATURE",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用 ID",
                        "name": "appId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件名(包含路径)",
                        "name": "key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)",
                        "name": "ossType",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.UploadResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/upload-part": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "oss 分段上传"
                ],
                "summary": "上传段",
                "parameters": [
                    {
                        "type": "string",
                        "description": "MOCK",
                        "name": "FZM-SIGNATURE",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "aliyun, huaweiyun 除了最后一段以外，其他段的大小范围是100KB~5GB；最后段大小范围是0~5GB。minio 除了最后一段以外，其他段的大小范围是5MB~5GB；最后段大小范围是0~5GB。",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "应用 ID",
                        "name": "appId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文件名(包含路径)",
                        "name": "key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)",
                        "name": "ossType",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "分段序号, 范围是1~10000",
                        "name": "partNumber",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "分段上传任务全局唯一标识",
                        "name": "uploadId",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.UploadPartResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AbortMultipartUploadRequest": {
            "type": "object",
            "required": [
                "appId",
                "key",
                "uploadId"
            ],
            "properties": {
                "appId": {
                    "description": "应用 ID",
                    "type": "string"
                },
                "key": {
                    "description": "文件名(包含路径)",
                    "type": "string"
                },
                "ossType": {
                    "description": "云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)",
                    "type": "string"
                },
                "uploadId": {
                    "description": "分段上传任务全局唯一标识",
                    "type": "string"
                }
            }
        },
        "model.AbortMultipartUploadResponse": {
            "type": "object"
        },
        "model.CompleteMultipartUploadRequest": {
            "type": "object",
            "required": [
                "appId",
                "key",
                "parts",
                "uploadId"
            ],
            "properties": {
                "appId": {
                    "description": "应用 ID",
                    "type": "string"
                },
                "key": {
                    "description": "文件名(包含路径)",
                    "type": "string"
                },
                "ossType": {
                    "description": "云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)",
                    "type": "string"
                },
                "parts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Part"
                    }
                },
                "uploadId": {
                    "description": "分段上传任务全局唯一标识",
                    "type": "string"
                }
            }
        },
        "model.CompleteMultipartUploadResponse": {
            "type": "object",
            "properties": {
                "uri": {
                    "type": "string"
                },
                "url": {
                    "description": "资源链接",
                    "type": "string"
                }
            }
        },
        "model.GeneralResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "integer"
                },
                "result": {
                    "type": "integer"
                }
            }
        },
        "model.GetHostReq": {
            "type": "object",
            "required": [
                "appId"
            ],
            "properties": {
                "appId": {
                    "description": "应用 ID",
                    "type": "string"
                },
                "ossType": {
                    "description": "云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)",
                    "type": "string"
                }
            }
        },
        "model.GetHostResp": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                }
            }
        },
        "model.InitMultipartUploadRequest": {
            "type": "object",
            "required": [
                "appId",
                "key"
            ],
            "properties": {
                "appId": {
                    "description": "应用 ID",
                    "type": "string"
                },
                "key": {
                    "description": "文件名(包含路径)",
                    "type": "string"
                },
                "ossType": {
                    "description": "云服务商, 可不填, 会选择默认服务商, 目前可选huaweiyun,aliyun,minio(不支持临时角色)",
                    "type": "string"
                }
            }
        },
        "model.InitMultipartUploadResponse": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "文件名(包含路径)",
                    "type": "string"
                },
                "uploadId": {
                    "description": "分段上传任务全局唯一标识",
                    "type": "string"
                }
            }
        },
        "model.Part": {
            "type": "object",
            "required": [
                "ETag",
                "partNumber"
            ],
            "properties": {
                "ETag": {
                    "description": "段数据的MD5值",
                    "type": "string"
                },
                "partNumber": {
                    "description": "分段序号, 范围是1~10000",
                    "type": "integer"
                }
            }
        },
        "model.UploadPartResponse": {
            "type": "object",
            "required": [
                "ETag",
                "partNumber"
            ],
            "properties": {
                "ETag": {
                    "description": "段数据的MD5值",
                    "type": "string"
                },
                "key": {
                    "description": "文件名(包含路径)",
                    "type": "string"
                },
                "partNumber": {
                    "description": "分段序号, 范围是1~10000",
                    "type": "integer"
                },
                "uploadId": {
                    "description": "分段上传任务全局唯一标识",
                    "type": "string"
                }
            }
        },
        "model.UploadResponse": {
            "type": "object",
            "properties": {
                "uri": {
                    "type": "string"
                },
                "url": {
                    "description": "资源链接",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:18005",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "云存储服务接口",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
