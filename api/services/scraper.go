package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testSearch (c *gin.Context) {
	c.JSON(http.StatusOK, "data")
}