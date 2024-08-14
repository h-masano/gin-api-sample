package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func (server *SampleServer) GetSample(c *gin.Context, params GetSampleParams) {
	responseArray := []SampleResponse{
		{Value: params.Param1},
		{Value: lo.ToPtr(strconv.Itoa(server.ServerId))},
	}
	var responseBody GetSample200JSONResponse = GetSample200JSONResponse(responseArray)
	c.JSON(http.StatusOK, responseBody)
}
