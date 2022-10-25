// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Greeter Server",
    "version": "1.0.0"
  },
  "paths": {
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getHello",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains the actual greeting as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/time": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getTime",
        "responses": {
          "200": {
            "description": "returns the current time",
            "schema": {
              "description": "contains current timestamp as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "name of the person",
            "name": "name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "name of the person",
            "name": "pass",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns all users within criteria",
            "schema": {
              "description": "contains JSON of user name and boolean if login was successful",
              "type": "object",
              "required": [
                "name",
                "result"
              ],
              "properties": {
                "name": {
                  "type": "string"
                },
                "pass": {
                  "type": "boolean"
                }
              }
            }
          },
          "500": {
            "description": "error",
            "schema": {
              "type": "object",
              "required": [
                "error"
              ],
              "properties": {
                "error": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "Greeter Server",
    "version": "1.0.0"
  },
  "paths": {
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getHello",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains the actual greeting as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/time": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getTime",
        "responses": {
          "200": {
            "description": "returns the current time",
            "schema": {
              "description": "contains current timestamp as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "name of the person",
            "name": "name",
            "in": "query"
          },
          {
            "type": "string",
            "description": "name of the person",
            "name": "pass",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns all users within criteria",
            "schema": {
              "description": "contains JSON of user name and boolean if login was successful",
              "type": "object",
              "required": [
                "name",
                "result"
              ],
              "properties": {
                "name": {
                  "type": "string"
                },
                "pass": {
                  "type": "boolean"
                }
              }
            }
          },
          "500": {
            "description": "error",
            "schema": {
              "type": "object",
              "required": [
                "error"
              ],
              "properties": {
                "error": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}
