// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "credential"
                ],
                "summary": "create Token",
                "parameters": [
                    {
                        "description": "login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Data",
                        "schema": {
                            "$ref": "#/definitions/auth.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPError"
                        }
                    }
                }
            }
        },
        "/spaceship": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaceship"
                ],
                "summary": "get list spaceship",
                "parameters": [
                    {
                        "type": "string",
                        "name": "class",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "$ref": "#/definitions/spaceship.FindSpaceshipResponse"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaceship"
                ],
                "summary": "create spaceship",
                "parameters": [
                    {
                        "description": "login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spaceship.CreateSpaceshipRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "success true",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "success false",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPResponse"
                        }
                    }
                }
            }
        },
        "/spaceship/id/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaceship"
                ],
                "summary": "get spaceship",
                "parameters": [
                    {
                        "type": "string",
                        "description": "spaceship id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spaceship.CreateSpaceshipRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data",
                        "schema": {
                            "$ref": "#/definitions/spaceship.SpaceshipResponse"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaceship"
                ],
                "summary": "update spaceship",
                "parameters": [
                    {
                        "type": "string",
                        "description": "spaceship id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/spaceship.CreateSpaceshipRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success true",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPResponse"
                        }
                    },
                    "400": {
                        "description": "success false",
                        "schema": {
                            "$ref": "#/definitions/http.HTTPResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "auth.LoginResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "token": {
                    "$ref": "#/definitions/auth.Token"
                }
            }
        },
        "auth.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "http.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "http.HTTPResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "spaceship.Armament": {
            "type": "object",
            "properties": {
                "qty": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "spaceship.CreateSpaceshipRequest": {
            "type": "object",
            "required": [
                "armament",
                "class",
                "name",
                "status"
            ],
            "properties": {
                "armament": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spaceship.Armament"
                    }
                },
                "class": {
                    "type": "string"
                },
                "crew": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "spaceship.DataSpaceship": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "spaceship.FindSpaceshipResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spaceship.DataSpaceship"
                    }
                }
            }
        },
        "spaceship.SpaceshipResponse": {
            "type": "object",
            "properties": {
                "armament": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/spaceship.Armament"
                    }
                },
                "class": {
                    "type": "string"
                },
                "crew": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "alchemy",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
