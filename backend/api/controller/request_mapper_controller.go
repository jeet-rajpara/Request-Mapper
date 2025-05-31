package controller

import (
	"net/http"

	"request-mapper/api/service"
	er "request-mapper/error"

	"github.com/gin-gonic/gin"
)

type RequestMapperController struct {
	service service.RequestMapperService
}

func NewRequestMapperController(service service.RequestMapperService) *RequestMapperController {
	return &RequestMapperController{
		service: service,
	}
}

func (c *RequestMapperController) MapRequest(ctx *gin.Context) {
	var request struct {
		RequestJSON map[string]interface{} `json:"requestJson"`
		RequestMap  map[string]string      `json:"requestMapper"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		er.ErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid request format")
		return
	}

	result, err := c.service.MapRequest(request.RequestJSON, request.RequestMap)
	if err != nil {
		// if it's our custom error type
		if customErr, ok := err.(*er.Error); ok {
			er.ErrorResponse(ctx.Writer, customErr.HttpStatusCode, customErr.ErrMessage)
			return
		}
		er.ErrorResponse(ctx.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}
