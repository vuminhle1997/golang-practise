package hello

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloRouterRegistration(router *gin.RouterGroup) {
	router.GET("/", sayHello)
	router.POST("/:name", greetings)
	router.DELETE("/:hate", sayHate)
	router.PUT("/:change", sayWish)
}

func sayHello(c *gin.Context) {
	p := Person{
		Age:      20,
		Gender:   "Male",
		Lastname: "Nguyen",
		Name:     "Kevin",
	}
	fmt.Println(p)
	c.PureJSON(200, gin.H{
		"person": p,
	})
}

func greetings(c *gin.Context) {
	name := c.Param("name")
	c.String(200, "Greetings %s", name)
}

func sayHate(c *gin.Context) {
	hate := c.Param("hate")
	c.String(200, "I hate %s", hate)
}

func sayWish(c *gin.Context) {
	change := c.Param("change")
	c.String(200, "I want . . .", change)
}
