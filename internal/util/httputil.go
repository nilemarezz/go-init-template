package httputil

import "github.com/gin-gonic/gin"

// NewError creates a new HTTPError response.
// @Summary Create a new error response
// @Description Create a new error response with the given status code and message
// @Param status query int true "Status code"
// @Param message query string true "Error message"
// @Produce json
// @Success 200 {object} HTTPError
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
