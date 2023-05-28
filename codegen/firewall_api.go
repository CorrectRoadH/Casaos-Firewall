// Package codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package codegen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	Access_tokenScopes = "access_token.Scopes"
)

// BaseResponse defines model for BaseResponse.
type BaseResponse struct {
	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// Rule defines model for Rule.
type Rule = string

// GetRulesResponseOK defines model for GetRulesResponseOK.
type GetRulesResponseOK struct {
	Data *[]Rule `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// GetVersionResponseOK defines model for GetVersionResponseOK.
type GetVersionResponseOK struct {
	Data *string `json:"data,omitempty"`

	// Message message returned by server side if there is any
	Message *string `json:"message,omitempty"`
}

// ResponseServiceUnavailable defines model for ResponseServiceUnavailable.
type ResponseServiceUnavailable = BaseResponse

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get Rules
	// (GET /nftables)
	GetRules(ctx echo.Context) error
	// Get Version
	// (GET /nftables/version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetRules converts echo context to params.
func (w *ServerInterfaceWrapper) GetRules(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRules(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	ctx.Set(Access_tokenScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/nftables", wrapper.GetRules)
	router.GET(baseURL+"/nftables/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xV32/bNhD+V4jbHlpAsYQEAzYBfeg2LAgyzEPTbQ9xYJyps8SWIokj5cQz9L8PpBTF",
	"sTt3a9FhTzZ/6Lvv++54twNpW2cNmeCh3AGTd9Z4SotLCm86Tf7NuDm/jrvSmkAmxL/onFYSg7Imf+et",
	"iXteNtRiOtV6vobydgdfM62hhK/yp2D5cM/n36OnxwDQZztwbB1xUAOHCkMCU4HatHEKK7KFPoOwdQQl",
	"IDNuoX/asKt3JAP0dxkwYTU3egtl4I76eKkiL1m5qAZKmF9HpEsKvxN7Zc3/wQR6wNbpqKSYfTcrYBLm",
	"AytTf57Ux/g3xBsl6TeDG1QaVzHeRwVPzHbQkvdYRw4jktiH6rPPNOefCPpg4ClysvQZaHlo+KRhd4A8",
	"Hgim0LGhSqy2whNviIVXFQm1FqEhJqG8QLOFbC9ncMj8OH0xDd1g43FmPcmOVdjeRBUDT5SSvF8G+55S",
	"GlTk2BBWxJCBwTZCvO5CY1n9mTL2VDPo1DVth6DKrO2x2EVXFBfSKRk6prSghRFCiOHA244liZYqha8W",
	"8MIxrYn9mbTa8lnymkpRIb9/uQDhWXoKrxbQhOB8meeM97NahaZbdZ54rK+ZtG1+JemPBjW9Jdnk2tY2",
	"b1GZXKJH68ef5QqNIV5G+KVRdROW3xaFe5g5Uy/gU8nqCPQF2YZ7lUIsV7qj04RVWwvUkcIP6HF+M5D6",
	"7xkNbPJnVRCLSIVU1AM58ZNiuketxetfryCDzdAzoYTNeXzy1pFBp6CEi1kxu4AMHIYmlXBu1iE+0LSo",
	"KRyX4Yu38x/nLyGhcCriqwrKaTqlZ7U3tc6L4u/mxHQv/8Bo6zP4prj4+Kcn+mR6pV3bIm8HguKRYcDa",
	"Q3kLvwxiRUuhsZWHu/jJZEE++favrRin1KeacTzkvogdTyxPGLLX6tJseN7kbu/iCBh6rk/nHWsoId+c",
	"5+uxCtOUGCMcWvizvT/TtKHHUh2b5MjDT0QixMNZwPqSbecGpIO7J2Xc9X8FAAD//0uaSeBcCQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
