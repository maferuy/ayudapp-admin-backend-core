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
        "contact": {
            "name": "Martin Fernandez",
            "url": "https://mafer.dev",
            "email": "maramal@outlook.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/appointments": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene todos las citas",
                "operationId": "get-appointments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetAppointmentsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetAppointmentsResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Crea una cita",
                "operationId": "create-appointment",
                "parameters": [
                    {
                        "description": "Datos de la cita",
                        "name": "CreateAppointmentRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateAppointmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.CreateAppointmentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.CreateAppointmentResponse"
                        }
                    }
                }
            }
        },
        "/admin/appointments/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene una cita",
                "operationId": "get-appointment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la cita",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetAppointmentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetAppointmentResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Actualiza una cita",
                "operationId": "update-appointment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la cita",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos de la cita",
                        "name": "UpdateAppointmentRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateAppointmentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateAppointmentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateAppointmentResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Elimina una cita",
                "operationId": "delete-appointment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la cita",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/categories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene las categor??as",
                "operationId": "get-categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetCategoriesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetCategoriesResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Crea una categor??a",
                "operationId": "create-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.CreateCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.CreateCategoryResponse"
                        }
                    }
                }
            }
        },
        "/admin/categories/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene una categor??a por su ID",
                "operationId": "get-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetCategoryResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Actualiza una categor??a",
                "operationId": "update-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateCategoryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateCategoryResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Elimina una categor??a",
                "operationId": "delete-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene todos los usuarios",
                "operationId": "get-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetUsersResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Crea un usuario",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "Datos del usuario",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.CreateUserResponse"
                        }
                    }
                }
            }
        },
        "/admin/users/email/{email}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene un usuario por su correo electr??nico",
                "operationId": "get-user-by-email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Correo electr??nico del usuario",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetUserResponse"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Obtiene un usuario",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.GetUserResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Actualiza un usuario",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos del usuario",
                        "name": "UpdateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateUserResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Elimina un usuario",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}/password": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cambia la contrase??a de un usuario",
                "operationId": "change-password",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Datos del usuario",
                        "name": "ChangePasswordRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.ChangePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}/set-super-admin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Configura un usuario como super administrador",
                "operationId": "set-super-admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}/unset-super-admin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Configura un super administrador como usuario",
                "operationId": "unset-super-admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Ingresa un usuario",
                "operationId": "login-user",
                "parameters": [
                    {
                        "description": "Datos del usuario",
                        "name": "loginUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Respuesta del login",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Error en la solicitud",
                        "schema": {
                            "$ref": "#/definitions/handlers.loginUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.loginUserRequest": {
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
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "handlers.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_expires_at": {
                    "type": "string"
                },
                "session_id": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/handlers.userResponse"
                }
            }
        },
        "handlers.userResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password_changed_at": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                }
            }
        },
        "models.Appointment": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "helper": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "password_changed_at": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "services.ChangePasswordRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "password_confirmation": {
                    "type": "string"
                }
            }
        },
        "services.CreateAppointmentRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "helper": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "services.CreateAppointmentResponse": {
            "type": "object",
            "properties": {
                "appointment_id": {
                    "type": "string"
                }
            }
        },
        "services.CreateCategoryResponse": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                }
            }
        },
        "services.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "services.CreateUserResponse": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "services.GetAppointmentResponse": {
            "type": "object",
            "properties": {
                "appointment": {
                    "$ref": "#/definitions/models.Appointment"
                }
            }
        },
        "services.GetAppointmentsResponse": {
            "type": "object",
            "properties": {
                "appointments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Appointment"
                    }
                }
            }
        },
        "services.GetCategoriesResponse": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Category"
                    }
                }
            }
        },
        "services.GetCategoryResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.Category"
                }
            }
        },
        "services.GetUserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "services.GetUsersResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "services.UpdateAppointmentRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "helper": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "services.UpdateAppointmentResponse": {
            "type": "object",
            "properties": {
                "appointment": {
                    "$ref": "#/definitions/models.Appointment"
                }
            }
        },
        "services.UpdateCategoryResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.Category"
                }
            }
        },
        "services.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "profile_image": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "services.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "API de Administraci??n de Ayud App",
	Description:      "API para administradores de Ayud App",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
