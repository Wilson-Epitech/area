package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func sumNumbers(context *gin.Context) {
	firstNbr, err1 := strconv.Atoi(context.Param("numberFrst"))
	sndNbr, err2 := strconv.Atoi(context.Param("numberSnd"))

	if err1 != nil || err2 != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid numbers provided",
		})
		return
	}

	sum := firstNbr + sndNbr

	context.JSON(http.StatusOK, gin.H{
		"Result of sum": sum,
	})
}

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/sum/:numberFrst/:numberSnd", sumNumbers)
	router.Run(":8080")
}
