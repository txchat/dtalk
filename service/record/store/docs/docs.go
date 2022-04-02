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
        "/app/pri-chat-record": {
            "post": {
                "tags": [
                    "record 获得聊天记录"
                ],
                "summary": "获得聊天记录",
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
                        "schema": {
                            "$ref": "#/definitions/model.GetPriRecordsReq"
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
                                            "$ref": "#/definitions/model.GetPriRecordsResp"
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
        "model.GeneralResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "integer"
                },
                "result": {
                    "type": "integer"
                }
            }
        },
        "model.GetPriRecordsReq": {
            "type": "object",
            "required": [
                "count",
                "targetId"
            ],
            "properties": {
                "count": {
                    "description": "消息数量",
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "logId": {
                    "description": "消息 ID",
                    "type": "string"
                },
                "targetId": {
                    "description": "接受者 ID",
                    "type": "string"
                }
            }
        },
        "model.GetPriRecordsResp": {
            "type": "object",
            "properties": {
                "record_count": {
                    "description": "聊天记录数量",
                    "type": "integer"
                },
                "records": {
                    "description": "聊天记录",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Record"
                    }
                }
            }
        },
        "model.Record": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "消息内容"
                },
                "createTime": {
                    "description": "消息发送时间",
                    "type": "integer"
                },
                "fromId": {
                    "description": "发送者 id",
                    "type": "string"
                },
                "logId": {
                    "description": "log id",
                    "type": "string"
                },
                "msgId": {
                    "description": "msg id (uuid)",
                    "type": "string"
                },
                "msgType": {
                    "description": "消息类型",
                    "type": "integer"
                },
                "targetId": {
                    "description": "接收者 id",
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
	Host:        "127.0.0.1:18015",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "聊天记录",
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
	swag.Register("swagger", &s{})
}