// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/reports": {
            "get": {
                "description": "get MyReports",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "List reports",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ReportForm"
                        }
                    }
                }
            }
        },
        "/reports/upload": {
            "post": {
                "description": "에러 신고에 이미지, 장소, 내용이 담긴다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reports"
                ],
                "summary": "에러 신고 업로드",
                "parameters": [
                    {
                        "type": "string",
                        "name": "content",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "example": 37.55327189658719,
                        "name": "latitude",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "example": 126.97232991836047,
                        "name": "longtitude",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ReportForm"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ReportForm": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number",
                    "example": 37.55327189658719
                },
                "longtitude": {
                    "type": "number",
                    "example": 126.97232991836047
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
