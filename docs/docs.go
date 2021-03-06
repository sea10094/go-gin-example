// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-13 16:25:07.196503018 +0800 CST m=+0.100263410

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
        "termsOfService": "https://github.com/EDDYCJY/go-gin-example",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/assetSearch": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get asset details of an agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task id",
                        "name": "taskid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "agent id",
                        "name": "agentid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/getAgentList": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get agent list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/getAgentMonitor": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get agent monitor info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "agent id",
                        "name": "agentid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/getProgress": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get progress of a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task id",
                        "name": "taskid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "agent id",
                        "name": "agentid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/getTaskAttr": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get task attr",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task id",
                        "name": "taskid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/getTaskList": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get task list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/rebootAgent": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "reboot the specialized agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "agent id",
                        "name": "agentid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/startScan": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Start scan task to the pointed agent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "agent id",
                        "name": "agentid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ip segment",
                        "name": "ips",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin API",
	Description: "An example of gin",
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
