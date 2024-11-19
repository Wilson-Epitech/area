package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func sumNumbers(context *gin.Context) {
	firstNbr, err1 := strconv.Atoi(context.Param("numberFrst"))
	sndNbr, err2 := strconv.Atoi(context.Param("numberSnd"))

	if err1 != nil || err2 != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid numbers",
		})
		return
	}

	sum := firstNbr + sndNbr

	context.JSON(http.StatusOK, gin.H{
		"Result of sum": sum,
	})
}

func helloServer(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "HI YOU!",
	})
}

func randomNumber(context *gin.Context) {
	maximunNbr, err := strconv.Atoi(context.Param("maxRand"))

	if err != nil || maximunNbr <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid number",
		})
		return
	}
	randNbr := rand.Intn(maximunNbr) + 1

	context.JSON(http.StatusOK, gin.H{
		"Random Number": randNbr,
	})
}

func main() {
	var router *gin.Engine = gin.Default()
	router.GET("/sum/:numberFrst/:numberSnd", sumNumbers)
	router.GET("/hi", helloServer)
	router.GET("/random/:maxRand", randomNumber)
	router.Run(":8080")
}
