package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func main() {
	r := gin.Default()

	r.POST("/parse", func(c *gin.Context) {
		body, _ := c.GetRawData()
		// Use gjson to extract a value
		value := gjson.GetBytes(body, "name")
		c.JSON(http.StatusOK, gin.H{
			"name": value.String(),
		})
	})

	r.Run(":8080")
}