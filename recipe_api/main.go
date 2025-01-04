//	Recipes API
//
//	This is a sample recipes API. You can find out more about the API at https://github.com/DeleMike/go-learn/recipe_api/
//
//		Schemes: http
// 		Host: localhost:8080
// 		BasePath: /
// 		Version: 1.0.0
//		Contact: Akindele Michael
//		<akindelemichael65@gmail.com> https://akindelemichael-1.web.app/
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
// swagger:meta

package main

import (
	"context"
	"github.com/delemike/recipe_api/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var recipesHandler *handlers.RecipesHandler

func main() {
	Init()
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.GET("/recipes/:id", recipesHandler.GetRecipeByIDHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.PATCH("/recipes/:id", recipesHandler.UpdateRecipeByPatchHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("recipes/search", recipesHandler.SearchRecipeHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func Init() {
	// Initialize context
	ctx := context.Background()
	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create the MongoDB client
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)

}
