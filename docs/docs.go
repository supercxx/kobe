// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-05 18:43:37.744527 +0800 CST m=+0.072703130

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/inventory/": {
            "get": {
                "description": "List Inventory",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory"
                ],
                "summary": "List Inventory",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Inventory"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create Inventory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory"
                ],
                "summary": "Create Inventory",
                "parameters": [
                    {
                        "description": "create inventory",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Inventory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Inventory"
                        }
                    }
                }
            }
        },
        "/inventory/{name}": {
            "get": {
                "description": "Get Inventory",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory"
                ],
                "summary": "Get Inventory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Inventory"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Inventory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory"
                ],
                "summary": "Update Inventory",
                "parameters": [
                    {
                        "description": "update inventory",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Inventory"
                        }
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/models.Inventory"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Inventory",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "inventory"
                ],
                "summary": "Delete Inventory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/playbooks/": {
            "get": {
                "description": "List all playbooks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playbook"
                ],
                "summary": "List playbooks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Playbook"
                            }
                        }
                    }
                }
            }
        },
        "/playbooks/{dir}": {
            "get": {
                "description": "List all playbooks under dir",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playbook"
                ],
                "summary": "List playbooks under dir",
                "parameters": [
                    {
                        "type": "string",
                        "description": "dir",
                        "name": "dir",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Playbook"
                            }
                        }
                    }
                }
            }
        },
        "/result/{uid}": {
            "get": {
                "description": "Get task result by task id when task finished",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "result"
                ],
                "summary": "Get Task Result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task_uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/tasks/": {
            "get": {
                "description": "List task info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "List Task Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            }
        },
        "/tasks/adhoc/": {
            "post": {
                "description": "Create Run Adhoc Task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "runner"
                ],
                "summary": "RunAdhoc",
                "parameters": [
                    {
                        "description": "create adhoc task",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RunAdhocRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            }
        },
        "/tasks/playbook/": {
            "post": {
                "description": "Create Run Playbook Task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "runner"
                ],
                "summary": "RunPlaybook",
                "parameters": [
                    {
                        "description": "create playbook task",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RunPlaybookRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            }
        },
        "/tasks/{uid}": {
            "get": {
                "description": "Get task info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Get Task Info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task_uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Inventory": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Playbook": {
            "type": "object",
            "properties": {
                "dir": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Result": {
            "type": "object",
            "properties": {
                "end_time": {
                    "type": "string"
                },
                "logfile": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "stdout": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.RunAdhocRequest": {
            "type": "object",
            "properties": {
                "arg": {
                    "type": "string"
                },
                "inventory": {
                    "type": "string"
                },
                "module": {
                    "type": "string"
                },
                "pattern": {
                    "type": "string"
                }
            }
        },
        "models.RunPlaybookRequest": {
            "type": "object",
            "properties": {
                "dir": {
                    "type": "string"
                },
                "inventory": {
                    "type": "string"
                },
                "playbook": {
                    "type": "string"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "created_time": {
                    "type": "string"
                },
                "finished": {
                    "type": "boolean"
                },
                "state": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "uid": {
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "Kobe Restful API",
	Description: "This is RestAPI Client for ansible",
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
