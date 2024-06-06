package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type counter struct {
	ID       string `json:"id"`
	Count    int    `json:"count"`
	Time     string `json:"time"`
	TimeDiff string `json:"timeDiff"`
	SN       string `json:"sn"`
}

var counters = []counter{
	{ID: "1", Count: 0},
}
var count = 0
var reset = 0
var prevTime time.Time

var increaseCounter = func(c *gin.Context) {
	count++
	if count == 10 {
		count = 0
		reset++
	}

	now := time.Now()

	diff := now.Sub(prevTime)

	sn := c.Query("SN")

	c.JSON(200, gin.H{
		"count":    count,
		"reset":    reset,
		"time":     now.Format(time.RFC3339),
		"timeDiff": diff.String(),
		"sn":       sn,
	})

	// Update prevTime to the current time
	prevTime = now
}

func main() {
	router := gin.Default()

	router.GET("/", increaseCounter)

	err := router.Run(":9090")
	if err != nil {
		log.Fatal(err)
		return
	}
}
