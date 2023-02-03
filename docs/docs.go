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
        "/order/": {
            "post": {
                "description": "Create a Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Create",
                "operationId": "do-post",
                "parameters": [
                    {
                        "description": "Set up order",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/raffle/": {
            "get": {
                "description": "Show all available raffles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "raffle"
                ],
                "summary": "Show raffles",
                "operationId": "getAll",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.availableResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a Raffle",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "raffle"
                ],
                "summary": "Create",
                "operationId": "do-create",
                "parameters": [
                    {
                        "description": "Set up raffle",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/raffle.Request"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/raffle/:id": {
            "get": {
                "description": "Show raffle by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "raffle"
                ],
                "summary": "Show raffles",
                "operationId": "getbyID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.availableResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/steam/do-player-inventory": {
            "post": {
                "description": "Create a Player Inventory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "steam"
                ],
                "summary": "Create",
                "operationId": "do-player-inventory",
                "parameters": [
                    {
                        "description": "set up steam",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doSteamRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/skin.Skin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "order.Request": {
            "type": "object",
            "required": [
                "numbers",
                "productID"
            ],
            "properties": {
                "numbers": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "example": [
                        1,
                        2
                    ]
                },
                "productID": {
                    "type": "string",
                    "example": "30dd879c-ee2f-11db-8314-0800200c9a66"
                }
            }
        },
        "raffle.Numbers": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "61f0c143ad06223fa03910b0"
                },
                "name": {
                    "type": "string",
                    "example": "Number"
                },
                "number": {
                    "type": "integer",
                    "example": 5
                },
                "status": {
                    "type": "string",
                    "example": "paid"
                }
            }
        },
        "raffle.Raffle": {
            "type": "object",
            "properties": {
                "PrizeDrawNumber": {
                    "type": "integer",
                    "example": 10
                },
                "description": {
                    "type": "string",
                    "example": "Rifa description"
                },
                "id": {
                    "type": "string",
                    "example": "61f0c143ad06223fa03910b0"
                },
                "imageURL": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "Rifa"
                },
                "numbers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/raffle.Numbers"
                    }
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                },
                "slug": {
                    "type": "string",
                    "example": "butterfly-32"
                },
                "status": {
                    "type": "string",
                    "example": "open"
                },
                "unitPrice": {
                    "type": "integer",
                    "example": 5
                },
                "userLimit": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "raffle.Request": {
            "type": "object",
            "required": [
                "description",
                "imageURL",
                "name",
                "quantity",
                "unitPrice"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Rifa"
                },
                "imageURL": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "Rifa"
                },
                "quantity": {
                    "type": "integer",
                    "example": 10
                },
                "unitPrice": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "skin.Item": {
            "type": "object",
            "properties": {
                "icon_url": {
                    "type": "string"
                },
                "market_hash_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "skin.Skin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/skin.Item"
                    }
                },
                "steam_id": {
                    "type": "string"
                }
            }
        },
        "v1.availableResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/raffle.Raffle"
                    }
                }
            }
        },
        "v1.doSteamRequest": {
            "type": "object",
            "required": [
                "steam_id"
            ],
            "properties": {
                "steam_id": {
                    "type": "string",
                    "example": "894012849024820948209"
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go Rifa Microservice",
	Description:      "Microservice for Rifa",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
