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

// AlgorithmBubbleSortRequest defines model for AlgorithmBubbleSortRequest.
type AlgorithmBubbleSortRequest struct {
	// Array An array of integers to be sorted.
	Array []int `json:"array"`
}

// AlgorithmBubbleSortResponse defines model for AlgorithmBubbleSortResponse.
type AlgorithmBubbleSortResponse struct {
	// SortedArray The sorted array of integers.
	SortedArray *[]int `json:"sortedArray,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Error message
	Error *string `json:"error,omitempty"`
}

// AlgorithmBinarySearchJSONRequestBody defines body for AlgorithmBinarySearch for application/json ContentType.
type AlgorithmBinarySearchJSONRequestBody = AlgorithmBinarySearchRequest

// PostAlgorithmBubleSortJSONRequestBody defines body for PostAlgorithmBubleSort for application/json ContentType.
type PostAlgorithmBubleSortJSONRequestBody = AlgorithmBubbleSortRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Binary search using golang
	// (POST /algorithm/binary-search)
	AlgorithmBinarySearch(c *gin.Context)
	// Sort an array using Bubble Sort Algorithm
	// (POST /algorithm/buble-sort)
	PostAlgorithmBubleSort(c *gin.Context)
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

// PostAlgorithmBubleSort operation middleware
func (siw *ServerInterfaceWrapper) PostAlgorithmBubleSort(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostAlgorithmBubleSort(c)
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
	router.POST(options.BaseURL+"/algorithm/buble-sort", wrapper.PostAlgorithmBubleSort)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8SVyW7bMBCGX4Vge1QspwsQ6JYARZBDgSLJrfCBlscSA4pkhsO0RuB3L4b0IkdyuiBp",
	"bhKXmW+2n4+ydp13FiwFWT3KULfQqfR5bhqHmtruQluFqxtQWLfXcB8hEO97dB6QNKTTClGt+GMBoUbt",
	"STsrK3luRdoRbim0JWgAgyAnQrImfmhqtS1EFwOJOYjgkGAxkYXUBF0yTCsPspKby3JdbFeyR/5X2AAN",
	"fd+2IPLe1nXP89Kh0FZQCxmQfT51tS4kwn3UCAtZfd843Pmb7S64+R3UxChHcha8swGGSdN2AT/HwdMW",
	"Z436Uex5C+FQnJwKvRTWkVi6aBdHYniGMs7nBm4c0kvV9d9qOJro2Z+SH8tuBjkfD4BznA8MA/lb9gHm",
	"F0SHx8GAt4dI6ZboIATVwL6UgVDbZswRL2m7dENTN0CBI1LbfAWhO2+gA8sRaysunSyk0TVsCK3q2PTX",
	"q9sUoibDv5dO7FIuC/kAGLKD08l0MuWTzoNVXstKfkxLhfSK2hRmufNeztM4nOTZSwlxudk4LYqhrxbc",
	"VWPjI3NzQKALt0iVrJ0lsOm+8t7oOlko7wKjbTWMv94jLGUl35V7kSs3Clc+K2/rw5YkjJAWckVTdB+m",
	"09dm2fRPgnlS31jXEMIyGrGFSlqa9OF30sFl+/SC9IfdPkJ7ZR+U0dx2PlIhYNJMis3UsXZtxpB7X4eg",
	"bSO8QtUBAQZm/fx/WQnQKiMC4AOgyLPK50LsOoUrWclcpe1LEhNy44yyTXodmpBEbLSXZ2ypPxhxbuCE",
	"M9CfikOka6CINqQC9iWLZepwfr65QH2JzAr56gM0eEVG8ppk+OCFkG8zYcOHY4Q2nxJ8jAcsGnrboXnS",
	"gIlLbZ/g3IB95L1mJ8O5lbkrH2VEIyvZEvmqLI2rlWldoOpsejaV69n6VwAAAP//t4hc3RAKAAA=",
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
