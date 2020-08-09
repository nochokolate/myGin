package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// Add 获取两数之和
func Add(a, b int) int {
	return a + b
}

func main() {
	r := gin.New()
	r.GET("/abcd", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
		// panic("panic-----")
	})
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
