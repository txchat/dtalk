// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/app/cdk/create-cdk-order": {
            "post": {
                "tags": [
                    "Cdk App"
                ],
                "summary": "创建兑换订单",
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
                            "$ref": "#/definitions/types.CreateCdkOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.CreateCdkOrderResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/cdk/deal-cdk-order": {
            "post": {
                "tags": [
                    "Cdk App"
                ],
                "summary": "处理兑换订单",
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
                            "$ref": "#/definitions/types.DealCdkOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.DealCdkOrderResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/cdk/get-cdk-type-by-coin-name": {
            "post": {
                "tags": [
                    "Cdk App"
                ],
                "summary": "查询一个票券对应的 cdkType",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.GetCdkTypeByCoinNameReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.GetCdkTypeByCoinNameResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/app/cdk/get-cdks-by-user-id": {
            "post": {
                "tags": [
                    "Cdk App"
                ],
                "summary": "分页获得一个人拥有的 cdks",
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
                            "$ref": "#/definitions/types.GetCdksByUserIdReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.GetCdksByUserIdResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/create-cdk-type": {
            "post": {
                "tags": [
                    "Cdk 后台"
                ],
                "summary": "创建 CdkType",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.CreateCdkTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.CreateCdkTypeResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/create-cdks": {
            "post": {
                "tags": [
                    "Cdk 后台"
                ],
                "summary": "创建 Cdks",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.CreateCdksReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.CreateCdksResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/delete-cdk-types": {
            "post": {
                "tags": [
                    "Cdk 后台"
                ],
                "summary": "删除 cdkTypes",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.DeleteCdkTypesReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.DeleteCdkTypesResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/delete-cdks": {
            "post": {
                "tags": [
                    "Cdk 后台"
                ],
                "summary": "删除 cdks",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.DeleteCdksReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.DeleteCdksResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/exchange-cdks": {
            "post": {
                "tags": [
                    "Cdk 后台"
                ],
                "summary": "兑换 cdks",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.ExchangeCdksReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.ExchangeCdksResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/get-cdk-types": {
            "post": {
                "tags": [
                    "Cdk 查询"
                ],
                "summary": "分页获得 cdkType",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.GetCdkTypesReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.GetCdkTypesResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/get-cdks": {
            "post": {
                "tags": [
                    "Cdk 查询"
                ],
                "summary": "分页获得 cdks",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.GetCdksReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.GetCdksResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/backend/cdk/update-cdk-type": {
            "post": {
                "tags": [
                    "Cdk 后台"
                ],
                "summary": "更新 cdkType",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.UpdateCdkTypeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/types.GeneralResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/types.UpdateCdkTypeResp"
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
        "types.Cdk": {
            "type": "object",
            "properties": {
                "cdkContent": {
                    "type": "string"
                },
                "cdkId": {
                    "type": "string"
                },
                "cdkName": {
                    "type": "string"
                },
                "cdkStatus": {
                    "type": "integer"
                },
                "createTime": {
                    "type": "string"
                },
                "exchangeTime": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "orderId": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "types.CdkType": {
            "type": "object",
            "properties": {
                "cdkAvailable": {
                    "description": "未发放的cdk数量",
                    "type": "integer"
                },
                "cdkFrozen": {
                    "description": "冻结状态中的cdk数量",
                    "type": "integer"
                },
                "cdkId": {
                    "type": "string"
                },
                "cdkInfo": {
                    "type": "string"
                },
                "cdkName": {
                    "type": "string"
                },
                "cdkUsed": {
                    "description": "已发放的cdk数量",
                    "type": "integer"
                },
                "coinName": {
                    "type": "string"
                },
                "exchangeRate": {
                    "type": "integer"
                }
            }
        },
        "types.CreateCdkOrderReq": {
            "type": "object",
            "required": [
                "cdkId",
                "number"
            ],
            "properties": {
                "cdkId": {
                    "description": "cdk 种类编号",
                    "type": "string"
                },
                "number": {
                    "description": "兑换数量",
                    "type": "integer"
                }
            }
        },
        "types.CreateCdkOrderResp": {
            "type": "object",
            "properties": {
                "orderId": {
                    "description": "订单编号",
                    "type": "string"
                }
            }
        },
        "types.CreateCdkTypeReq": {
            "type": "object",
            "required": [
                "coinName",
                "exchangeRate"
            ],
            "properties": {
                "cdkInfo": {
                    "type": "string"
                },
                "cdkName": {
                    "type": "string"
                },
                "coinName": {
                    "type": "string"
                },
                "exchangeRate": {
                    "type": "integer"
                }
            }
        },
        "types.CreateCdkTypeResp": {
            "type": "object",
            "properties": {
                "cdkId": {
                    "type": "string"
                }
            }
        },
        "types.CreateCdksReq": {
            "type": "object",
            "required": [
                "cdkContents",
                "cdkId"
            ],
            "properties": {
                "cdkContents": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "cdkId": {
                    "type": "string"
                }
            }
        },
        "types.CreateCdksResp": {
            "type": "object"
        },
        "types.DealCdkOrderReq": {
            "type": "object",
            "required": [
                "orderId",
                "transferHash"
            ],
            "properties": {
                "orderId": {
                    "description": "订单编号",
                    "type": "string"
                },
                "result": {
                    "description": "处理结果",
                    "type": "boolean"
                },
                "transferHash": {
                    "description": "转账记录 hash",
                    "type": "string"
                }
            }
        },
        "types.DealCdkOrderResp": {
            "type": "object"
        },
        "types.DeleteCdkTypesReq": {
            "type": "object",
            "required": [
                "cdkIds"
            ],
            "properties": {
                "cdkIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "types.DeleteCdkTypesResp": {
            "type": "object"
        },
        "types.DeleteCdksReq": {
            "type": "object",
            "required": [
                "ids"
            ],
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "types.DeleteCdksResp": {
            "type": "object"
        },
        "types.ExchangeCdksReq": {
            "type": "object",
            "required": [
                "ids"
            ],
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "types.ExchangeCdksResp": {
            "type": "object"
        },
        "types.GeneralResponse": {
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
        "types.GetCdkTypeByCoinNameReq": {
            "type": "object",
            "required": [
                "coinName"
            ],
            "properties": {
                "coinName": {
                    "type": "string"
                }
            }
        },
        "types.GetCdkTypeByCoinNameResp": {
            "type": "object",
            "properties": {
                "cdkAvailable": {
                    "description": "未发放的cdk数量",
                    "type": "integer"
                },
                "cdkFrozen": {
                    "description": "冻结状态中的cdk数量",
                    "type": "integer"
                },
                "cdkId": {
                    "type": "string"
                },
                "cdkInfo": {
                    "type": "string"
                },
                "cdkName": {
                    "type": "string"
                },
                "cdkUsed": {
                    "description": "已发放的cdk数量",
                    "type": "integer"
                },
                "coinName": {
                    "type": "string"
                },
                "exchangeRate": {
                    "type": "integer"
                }
            }
        },
        "types.GetCdkTypesReq": {
            "type": "object",
            "properties": {
                "coinName": {
                    "type": "string"
                },
                "page": {
                    "description": "页数",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "types.GetCdkTypesResp": {
            "type": "object",
            "properties": {
                "cdkTypes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.CdkType"
                    }
                },
                "totalElements": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        },
        "types.GetCdksByUserIdReq": {
            "type": "object",
            "properties": {
                "page": {
                    "description": "页数",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "types.GetCdksByUserIdResp": {
            "type": "object",
            "properties": {
                "cdks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Cdk"
                    }
                },
                "totalElements": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        },
        "types.GetCdksReq": {
            "type": "object",
            "required": [
                "cdkId"
            ],
            "properties": {
                "cdkContent": {
                    "type": "string"
                },
                "cdkId": {
                    "type": "string"
                },
                "page": {
                    "description": "页数",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                }
            }
        },
        "types.GetCdksResp": {
            "type": "object",
            "properties": {
                "cdks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Cdk"
                    }
                },
                "totalElements": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        },
        "types.UpdateCdkTypeReq": {
            "type": "object",
            "required": [
                "cdkId",
                "cdkName",
                "coinName",
                "exchangeRate"
            ],
            "properties": {
                "cdkId": {
                    "type": "string"
                },
                "cdkName": {
                    "type": "string"
                },
                "coinName": {
                    "type": "string"
                },
                "exchangeRate": {
                    "type": "integer"
                }
            }
        },
        "types.UpdateCdkTypeResp": {
            "type": "object"
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}