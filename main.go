package flowvoy

import (
	"github.com/gin-gonic/gin"
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
		// TODO: Fetch data from Envoy and return it as JSON
	})

	r.Run()
}
