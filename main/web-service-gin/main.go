package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pcrTest/main/narcology"
)

func getNarcology(c *gin.Context) {
	s := narcology.Narcology()
	if s == "200 OK" {
		c.IndentedJSON(http.StatusOK, "200")
	} else {
		c.IndentedJSON(http.StatusInternalServerError, "500")
	}
}

func main() {
	router := gin.Default()
	router.GET("/narcology", getNarcology)

	router.Run("localhost:8080")
}
