package main

// import "github.com/gin-gonic/gin"
// import "github.com/369329303/myGin/gin"
import "github.com/369329303/myGin/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
