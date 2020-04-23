package main

import (
	"context"
	"time"

	"./hello"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)

func main() {
	r := gin.Default()

	// v1 REST API
	v1 := r.Group("/api")

	hello.HelloRouterRegistration(v1.Group("/hello"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil{
	// 	fmt.Println(err)
	// }

	// collection := client.Database("golang").Collection("users")

	// collection.InsertOne(ctx, bson.M{"name": "Kevin", "lastname": "Nguyen"})
	// r.GET('/', HomePage)
	// r.GET("/ping", func(c *gin.Context) {

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
