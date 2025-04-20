package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testIncludeOtherRouter() {
	router := gin.Default()

	router.GET("/test", getGreeting)
}

func getGreeting(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}
