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
        "/ins/ins": {
            "post": {
                "description": "輸入通報資料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/x-www-form-urlencoded"
                ],
                "tags": [
                    "Insert"
                ],
                "summary": "MannouInsert",
                "parameters": [
                    {
                        "description": "由上而下值為：問卷代號(請全部大寫)、項目代碼、答案、病歷號碼",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FeedBackFormat"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.FeedBackFormat": {
            "type": "object",
            "required": [
                "an_Code",
                "an_Cout1",
                "an_Cout2",
                "an_Cout3",
                "an_Cout4",
                "an_Cout5",
                "an_Cout6",
                "an_Name"
            ],
            "properties": {
                "an_Code": {
                    "type": "string"
                },
                "an_Cout1": {
                    "type": "string"
                },
                "an_Cout2": {
                    "type": "string"
                },
                "an_Cout3": {
                    "type": "string"
                },
                "an_Cout4": {
                    "type": "string"
                },
                "an_Cout5": {
                    "type": "string"
                },
                "an_Cout6": {
                    "type": "string"
                },
                "an_Name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "24.07b1",
	Host:             "10.8.41.142:6000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "toilet-feedback",
	Description:      "公共區域通報系統",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
