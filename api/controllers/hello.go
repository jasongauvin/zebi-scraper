package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SayHello simply returns a json with hello world text to check if the API is available
func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world!")
}