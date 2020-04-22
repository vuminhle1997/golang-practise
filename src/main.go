package main

import (
	"./hello"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// v1 REST API
	v1 := r.Group("/api")

	hello.HelloRouterRegistration(v1.Group("/hello"))
	// r.GET('/', HomePage)
	// r.GET("/ping", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
