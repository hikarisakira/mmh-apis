// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "李承洋/K861",
            "email": "hsakira.k861@mmh.org.tw"
        },
        "license": {
            "name": "Private/馬偕紀念醫院 員工保密協定"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/search/del/{pno}": {
            "get": {
                "description": "輸入病歷號碼來獲取取消掛號紀錄",
                "produces": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Search PatientService"
                ],
                "summary": "GetDelRecord",
                "parameters": [
                    {
                        "type": "string",
                        "description": "病歷號碼",
                        "name": "pno",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "24.02b7f3",
	Host:             "10.8.41.142:8080",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "testApi",
	Description:      "健檢中心表單API，目前僅使用於勞工體檢並已保留未來擴展可能。本機網域位置位於:10.8.41.142，使用時請使用時請將localhost部分改為前述網域位置。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}