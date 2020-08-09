package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.New()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello\n")
		// panic("panic-----")
	})

	r.GET("/world", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "world\n")
		// panic("panic-----")
	})
	// Listen and serve on 0.0.0.0:8080
	fmt.Println(r.Run(":8080"))
}
