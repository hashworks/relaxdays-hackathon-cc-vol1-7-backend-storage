// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {
            "name": "Justin Kromlinger",
            "url": "https://hashworks.net",
            "email": "justin.kromlinger@stud.htwk-leipzig.de"
        },
        "license": {
            "name": "GNU Affero General Public License v3",
            "url": "https://gnu.org/licenses/agpl.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/storagePlace": {
            "get": {
                "tags": [
                    "V0Storage"
                ],
                "summary": "Returns a specifc storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to retrieve",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.V0Storage"
                        }
                    },
                    "404": {
                        "description": "Storage not found",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "V0Storage"
                ],
                "summary": "Delete a storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to delete",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/storagePlace/": {
            "put": {
                "tags": [
                    "V0Storage"
                ],
                "summary": "Save a storage",
                "parameters": [
                    {
                        "description": "Storage to save",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.V0Storage"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "V0Storage"
                ],
                "summary": "Update a storage",
                "parameters": [
                    {
                        "description": "Storage to Update",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.V0Storage"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/storagesPlaces": {
            "get": {
                "tags": [
                    "V0Storage"
                ],
                "summary": "Returns \"n\" storages lexicographically after storage \"name\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of storages after the named one",
                        "name": "n",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Storage name where we should start the cursor",
                        "name": "x",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.V0Storage"
                            }
                        }
                    }
                }
            }
        },
        "/v1/storagePlace": {
            "get": {
                "tags": [
                    "V1Storage"
                ],
                "summary": "Returns a specifc storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to retrieve",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.V1Storage"
                        }
                    },
                    "404": {
                        "description": "Storage not found",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "V1Storage"
                ],
                "summary": "Delete a storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to delete",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/storagePlace/": {
            "put": {
                "tags": [
                    "V1Storage"
                ],
                "summary": "Save a storage",
                "parameters": [
                    {
                        "description": "Storage to save",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.V1Storage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "V1Storage"
                ],
                "summary": "Update a storage",
                "parameters": [
                    {
                        "description": "Storage to Update",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.V1Storage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/v1/storagesPlaces": {
            "get": {
                "tags": [
                    "V1Storage"
                ],
                "summary": "Returns \"n\" storages lexicographically after storage \"name\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of storages after the named one",
                        "name": "n",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Storage name where we should start the cursor",
                        "name": "x",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.V1Storage"
                            }
                        }
                    }
                }
            }
        },
        "/v2/storagePlace": {
            "get": {
                "tags": [
                    "V2Storage"
                ],
                "summary": "Returns a specifc storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to retrieve",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.V2Storage"
                        }
                    },
                    "404": {
                        "description": "Storage not found",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "V2Storage"
                ],
                "summary": "Delete a storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to delete",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/v2/storagePlace/": {
            "put": {
                "tags": [
                    "V2Storage"
                ],
                "summary": "Save a storage",
                "parameters": [
                    {
                        "description": "Storage to save",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.V2Storage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "V2Storage"
                ],
                "summary": "Update a storage",
                "parameters": [
                    {
                        "description": "Storage to Update",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.V2Storage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/v2/storagePlacesForArticleID": {
            "get": {
                "tags": [
                    "V2Storage"
                ],
                "summary": "Returns all storages that contain an article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "x",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.V2Storage"
                            }
                        }
                    }
                }
            }
        },
        "/v2/storagesPlaces": {
            "get": {
                "tags": [
                    "V2Storage"
                ],
                "summary": "Returns \"n\" storages lexicographically after storage \"name\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of storages after the named one",
                        "name": "n",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Storage name where we should start the cursor",
                        "name": "x",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.V2Storage"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.V0Storage": {
            "type": "object",
            "properties": {
                "articleID": {
                    "type": "integer"
                },
                "bestand": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.V1Storage": {
            "type": "object",
            "properties": {
                "articleID": {
                    "type": "integer"
                },
                "bestand": {
                    "type": "integer"
                },
                "hoehe": {
                    "type": "integer"
                },
                "lagerabschnitt": {
                    "type": "integer"
                },
                "platz": {
                    "type": "integer"
                },
                "reihe": {
                    "type": "integer"
                },
                "standort": {
                    "type": "string"
                }
            }
        },
        "models.V2Storage": {
            "type": "object",
            "properties": {
                "articleID": {
                    "type": "integer"
                },
                "bestand": {
                    "type": "integer"
                },
                "hoehe": {
                    "type": "integer"
                },
                "kapazitaet": {
                    "type": "integer"
                },
                "lagerabschnitt": {
                    "type": "integer"
                },
                "platz": {
                    "type": "integer"
                },
                "reihe": {
                    "type": "integer"
                },
                "standort": {
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
	Host:        "127.0.0.1:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Storage Backend Task",
	Description: "Solution for 'Lager' backend task of https://sites.google.com/relaxdays.de/hackathon-relaxdays/startseite#h.n7504a3wvsj5",
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
