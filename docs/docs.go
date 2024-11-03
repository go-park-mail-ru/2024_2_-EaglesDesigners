// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addchat": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Add new chat",
                "parameters": [
                    {
                        "description": "Chat info",
                        "name": "chat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChatDTOInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Чат создан"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "500": {
                        "description": "Не удалось добавить чат / группу"
                    }
                }
            }
        },
        "/auth": {
            "get": {
                "description": "Retrieve user data based on the JWT token present in the cookies.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate a user",
                "responses": {
                    "200": {
                        "description": "User data retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/delivery.AuthResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized: token is invalid",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chat/{chatId}/addusers": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Добавить пользователей в чат",
                "parameters": [
                    {
                        "description": "Пользователи на добавление",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddUsersIntoChatDTO"
                        }
                    },
                    {
                        "maxLength": 36,
                        "minLength": 36,
                        "type": "string",
                        "example": "\"123e4567-e89b-12d3-a456-426614174000\"",
                        "description": "Chat ID (UUID)",
                        "name": "chatId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пользователи добавлены"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "500": {
                        "description": "Не удалось добавить пользователей"
                    }
                }
            }
        },
        "/chat/{chatId}/delete": {
            "delete": {
                "tags": [
                    "chat"
                ],
                "summary": "Удаличть чат или группу",
                "parameters": [
                    {
                        "maxLength": 36,
                        "minLength": 36,
                        "type": "string",
                        "example": "\"123e4567-e89b-12d3-a456-426614174000\"",
                        "description": "Chat ID (UUID)",
                        "name": "chatId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Чат удалён"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "403": {
                        "description": "Нет полномочий"
                    },
                    "500": {
                        "description": "Не удалось удалить чат"
                    }
                }
            }
        },
        "/chat/{chatId}/messages": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "summary": "Add new message",
                "parameters": [
                    {
                        "maxLength": 36,
                        "minLength": 36,
                        "type": "string",
                        "example": "\"123e4567-e89b-12d3-a456-426614174000\"",
                        "description": "Chat ID (UUID)",
                        "name": "chatId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Message info",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Сообщение успешно добавлено"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "500": {
                        "description": "Не удалось добавить сообщение"
                    }
                }
            }
        },
        "/chats": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Get chats of user",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ChatsDTO"
                        }
                    },
                    "500": {
                        "description": "Не удалось получить сообщения"
                    }
                }
            }
        },
        "/contacts": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all contacts of user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Get all contacts",
                "responses": {
                    "200": {
                        "description": "Contacts found",
                        "schema": {
                            "$ref": "#/definitions/models.GetContactsRespDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
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
                "description": "Create a new contact for the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contacts"
                ],
                "summary": "Add new contact",
                "parameters": [
                    {
                        "description": "Credentials for create a new contact",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddContactReqDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Contact created",
                        "schema": {
                            "$ref": "#/definitions/models.ContactDTO"
                        }
                    },
                    "400": {
                        "description": "Failed to create contact",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate a user with username and password.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Credentials for login, including username and password",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.AuthCredentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Authentication successful",
                        "schema": {
                            "$ref": "#/definitions/responser.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid format JSON",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Incorrect login or password",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Invalidate the user's session by clearing the access token cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Log out a user",
                "responses": {
                    "200": {
                        "description": "Logout successful",
                        "schema": {
                            "$ref": "#/definitions/responser.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "No access token found",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get bio, avatar and birthdate of user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Get self profile data",
                "responses": {
                    "200": {
                        "description": "Profile data found",
                        "schema": {
                            "$ref": "#/definitions/models.GetProfileResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid format JSON",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update bio, avatar, name or birthdate of user.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Update profile data",
                "parameters": [
                    {
                        "description": "JSON representation of profile data",
                        "name": "profile_data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfileRequestDTO"
                        }
                    },
                    {
                        "type": "file",
                        "description": "User avatar image",
                        "name": "avatar",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile updated",
                        "schema": {
                            "$ref": "#/definitions/responser.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Failed to update profile",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/profile/{userid}": {
            "get": {
                "description": "Get bio, avatar and birthdate of user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Get profile data",
                "responses": {
                    "200": {
                        "description": "Profile data found",
                        "schema": {
                            "$ref": "#/definitions/models.GetProfileResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid format JSON",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Creates a new user with the provided credentials.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Registration information",
                        "name": "creds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delivery.RegisterCredentials"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Registration successful",
                        "schema": {
                            "$ref": "#/definitions/delivery.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "User failed to create",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "A user with that username already exists",
                        "schema": {
                            "$ref": "#/definitions/responser.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/uploads/{folder}/{name}": {
            "get": {
                "description": "Fetches an image from the specified folder and by filename",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uploads"
                ],
                "summary": "Retrieve an image",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"avatar\"",
                        "description": "Folder name",
                        "name": "folder",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"642c5a57-ebc7-49d0-ac2d-f2f1f474bee7.png\"",
                        "description": "File name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful image retrieval",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "404": {
                        "description": "File not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "delivery.AuthCredentials": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "12345678"
                },
                "username": {
                    "type": "string",
                    "example": "user11"
                }
            }
        },
        "delivery.AuthResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/delivery.UserData"
                }
            }
        },
        "delivery.RegisterCredentials": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Vincent Vega"
                },
                "password": {
                    "type": "string",
                    "example": "go_do_a_crime"
                },
                "username": {
                    "type": "string",
                    "example": "killer1994"
                }
            }
        },
        "delivery.RegisterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Registration successful"
                },
                "user": {
                    "$ref": "#/definitions/delivery.UserData"
                }
            }
        },
        "delivery.UserData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "2"
                },
                "name": {
                    "type": "string",
                    "example": "Dr Peper"
                },
                "username": {
                    "type": "string",
                    "example": "user12"
                }
            }
        },
        "model.AddUsersIntoChatDTO": {
            "type": "object",
            "properties": {
                "usersId": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "uuid1",
                        "uuid2"
                    ]
                }
            }
        },
        "model.ChatDTOInput": {
            "type": "object",
            "properties": {
                "avatarBase64": {
                    "type": "string"
                },
                "chatName": {
                    "type": "string",
                    "example": "Чат с пользователем 2"
                },
                "chatType": {
                    "type": "string",
                    "example": "personalMessages"
                },
                "usersToAdd": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "uuid1",
                        "uuid2"
                    ]
                }
            }
        },
        "model.ChatDTOOutput": {
            "type": "object",
            "properties": {
                "avatarBase64": {
                    "type": "string"
                },
                "chatId": {
                    "type": "string"
                },
                "chatName": {
                    "type": "string",
                    "example": "Чат с пользователем 2"
                },
                "chatType": {
                    "type": "string",
                    "example": "personal"
                },
                "countOfUsers": {
                    "type": "integer",
                    "example": 52
                },
                "lastMessage": {
                    "$ref": "#/definitions/models.Message"
                }
            }
        },
        "model.ChatsDTO": {
            "type": "object",
            "properties": {
                "chats": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ChatDTOOutput"
                    }
                }
            }
        },
        "models.AddContactReqDTO": {
            "type": "object",
            "properties": {
                "contactUsername": {
                    "type": "string",
                    "example": "user11"
                }
            }
        },
        "models.ContactDTO": {
            "type": "object",
            "properties": {
                "avatarURL": {
                    "type": "string",
                    "example": "/uploads/avatar/642c5a57-ebc7-49d0-ac2d-f2f1f474bee7.png"
                },
                "id": {
                    "type": "string",
                    "example": "08a0f350-e122-467b-8ba8-524d2478b56e"
                },
                "name": {
                    "description": "can be nil",
                    "type": "string",
                    "example": "Витек"
                },
                "username": {
                    "type": "string",
                    "example": "user11"
                }
            }
        },
        "models.GetContactsRespDTO": {
            "type": "object",
            "properties": {
                "contacts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ContactDTO"
                    }
                }
            }
        },
        "models.GetProfileResponseDTO": {
            "type": "object",
            "properties": {
                "avatarURL": {
                    "type": "string",
                    "example": "/2024_2_eaglesDesigners/uploads/avatar/f0364477-bfd4-496d-b639-d825b009d509.png"
                },
                "bio": {
                    "type": "string",
                    "example": "Не люблю сети"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2024-04-13T08:30:00Z"
                },
                "name": {
                    "type": "string",
                    "example": "Vincent Vega"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "authorID": {
                    "type": "string"
                },
                "authorName": {
                    "type": "string"
                },
                "chatId": {
                    "type": "string"
                },
                "datetime": {
                    "type": "string",
                    "example": "2024-04-13T08:30:00Z"
                },
                "isRedacted": {
                    "type": "boolean"
                },
                "messageId": {
                    "type": "string",
                    "example": "1"
                },
                "text": {
                    "type": "string",
                    "example": "тут много текста"
                }
            }
        },
        "models.UpdateProfileRequestDTO": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string",
                    "example": "Не люблю сети"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2024-04-13T08:30:00Z"
                },
                "name": {
                    "type": "string",
                    "example": "Vincent Vega"
                }
            }
        },
        "responser.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "responser.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "212.233.98.59:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Patefon API",
	Description:      "This is a description of the Patefon server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
