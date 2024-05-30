package main

import (
	"github.com/gin-gonic/gin"
)

// counter struct
type counter struct {
	ID    string `json:"id"`
	Count int    `json:"count"`
}

// counter variable
var counters = []counter{
	{ID: "1", Count: 0},
}
var count = 0
var reset = 0

var increaseCounter = func(c *gin.Context) {
	count++
	if count == 10 {
		count = 0
		reset++
	}
	c.JSON(200, gin.H{
		"count": count,
		"reset": reset,
	})
}

func main() {
	router := gin.Default()

	router.GET("/", increaseCounter)

	router.Run(":9090")
}
