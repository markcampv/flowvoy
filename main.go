package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	r := gin.Default()

	// Serve HTML, JS, and CSS files
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/data", func(c *gin.Context) {
		resp, err := http.Get("http://localhost:19000/stats?filter=&format=json")
		if err != nil {
			//handle error
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			//handle error
		}
		c.JSON(200, gin.H{
			"data": string(body),
		})
	})

	r.Run()
}
