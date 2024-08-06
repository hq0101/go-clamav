// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/allmatchscan": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "AllMatchScan scan a file or directory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File path to scan",
                        "name": "file",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cli.ScanResult"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/contscan": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Continuously scan a file or directory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File path to scan",
                        "name": "file",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cli.ScanResult"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/instream": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ClamAV"
                ],
                "summary": "Scan data stream",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cli.ScanResult"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/multiscan": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Multithreaded scan of a file or directory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File path to scan",
                        "name": "file",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cli.ScanResult"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Ping Check the server's state. It should reply with \"PONG\".",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reload": {
            "post": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Reload the virus database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/scan": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Scan a file or directory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File path to scan",
                        "name": "file",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/cli.ScanResult"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shutdown": {
            "post": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Shut down the ClamAV server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/stats": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Get ClamAV stats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/cli.ClamdStats"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Get the ClamAV version",
                "responses": {
                    "200": {
                        "description": "ClamAV 0.103.11/27353/Wed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/versioncommands": {
            "get": {
                "tags": [
                    "ClamAV"
                ],
                "summary": "Get the ClamAV version",
                "responses": {
                    "200": {
                        "description": "ClamAV 0.103.11/27353/Wed",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cli.ClamdStats": {
            "type": "object",
            "properties": {
                "free": {
                    "type": "number"
                },
                "heap": {
                    "type": "number"
                },
                "idleTimeout": {
                    "type": "integer"
                },
                "mmap": {
                    "type": "number"
                },
                "pools": {
                    "type": "integer"
                },
                "poolsTotal": {
                    "type": "number"
                },
                "poolsUsed": {
                    "type": "number"
                },
                "queueItems": {
                    "type": "integer"
                },
                "releasable": {
                    "type": "number"
                },
                "state": {
                    "type": "string"
                },
                "stats": {
                    "type": "number"
                },
                "threadsIdle": {
                    "type": "integer"
                },
                "threadsLive": {
                    "type": "integer"
                },
                "threadsMax": {
                    "type": "integer"
                },
                "used": {
                    "type": "number"
                }
            }
        },
        "cli.ScanResult": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "virus": {
                    "type": "string"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
