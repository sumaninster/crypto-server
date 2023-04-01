package main

import "github.com/gin-gonic/gin"

// NewError function
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPBadRequestError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPBadRequestError message
type HTTPBadRequestError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request: The server cannot or will not process the request due to something that is perceived to be a client error (for example, malformed request syntax, invalid request message framing, or deceptive request routing)."`
}

// HTTPFileNotFoundError message
type HTTPFileNotFoundError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found: Cannot retrieve the page that was requested. The following are some common causes of this error message: The requested file has been renamed. The requested file has been moved to another location and/or deleted."`
}

// HTTPInternalServerError message
type HTTPInternalServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error: The server encountered an unexpected condition that prevented it from fulfilling the request. This error response is a generic "catch-all" response."`
}