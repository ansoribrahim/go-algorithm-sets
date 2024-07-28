// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// AlgorithmBinarySearchRequest defines model for AlgorithmBinarySearchRequest.
type AlgorithmBinarySearchRequest struct {
	// Array An array of integers to search within, must be sorted.
	Array []int `json:"array"`

	// Target The target integer to search for in the array.
	Target int `json:"target"`
}

// AlgorithmBinarySearchResponse defines model for AlgorithmBinarySearchResponse.
type AlgorithmBinarySearchResponse struct {
	// Index The index of the target in the array, or -1 if not found.
	Index *int `json:"index,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error message
	Error *string `json:"error,omitempty"`
}

// AlgorithmBinarySearchJSONRequestBody defines body for AlgorithmBinarySearch for application/json ContentType.
type AlgorithmBinarySearchJSONRequestBody = AlgorithmBinarySearchRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Binary search using golang
	// (POST /algorithm/binary-search)
	AlgorithmBinarySearch(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AlgorithmBinarySearch operation middleware
func (siw *ServerInterfaceWrapper) AlgorithmBinarySearch(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AlgorithmBinarySearch(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/algorithm/binary-search", wrapper.AlgorithmBinarySearch)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xUTW/bMAz9KwK3oxun+wAK31pgKHrYZe1t6EG1aVuFLaok3S0o8t8Hykn6kWS7bLvZ",
	"osT3SL7HJ6hpTBQxqkD1BFL3OPr8eT50xEH78SJEz6tr9Fz33/BhQlGLJ6aErAHzbc/sV/bRoNQckgaK",
	"UMF5dDniqHUhKnbI4pSc5GzuR9A+xMKNk6i7QyfEis0CCgiKY06sq4RQweYxrIvtyYxo/5471H3smx7d",
	"HNtCv0BuiV2ITnucCRrmW6h1AYwPU2BsoPq+Adzh3e4e0N091mpUjvRMEkXB/aaF2ODPw8RzyLqmL6t4",
	"5ls4Yndy6kLrIqlraYrNkRr2WH5hJj7OCi28zyq/ciOK+A6fkUQ5xO4QkB2F2NJ+qmtUsdr8tl3iwpgG",
	"HDEqNlbnJUEBQ6hxwzD60VJ/vbrJEw862O8luV3HoYBHZJkBThfLxdJuUsLoU4AKPuajApLXPpdZ7tDL",
	"uzytk1kauSE0a9za4o30VWNiPjRdmEWCohfUZAfUFBVjfu9TGkKdM5T3YtS2FrOv94wtVPCufPZguTFg",
	"+Vv3rV9LU3nCfDBPNFf3Ybn811w2+slk3sx3qmsUaafBbUllq2f5/knZNrZPf5H9a7UfYHsVH/0QTHZp",
	"0sLholsUm61l1pp3ktltDCIhdi559iMqshjXz/+XqyJHPzhBfkR2s1ftnkzj6HkFFcxT2i66KVPuaPCx",
	"y8urk7zMDmr5NkPOue3aE0w8QAW9aqrKcqDaDz2JVmfLsyWsb9e/AgAA//84nBYxQAYAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
