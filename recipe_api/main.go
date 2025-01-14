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
	"fmt"
	"github.com/delemike/recipe_api/handlers"
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var authHandler *handlers.AuthHandler
var recipesHandler *handlers.RecipesHandler

func main() {
	router := gin.Default()

	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("session", store))
	// public routes
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.GET("recipes/search", recipesHandler.SearchRecipeHandler)

	// auth routes
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/signout", authHandler.SignOutHandler)
	//router.POST("/refresh", authHandler.RefreshHandler)
	router.POST("refresh", authHandler.RefreshCookie)

	// using middleware
	authorized := router.Group("/")
	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetRecipeByIDHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.PATCH("/recipes/:id", recipesHandler.UpdateRecipeByPatchHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	}

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func init() {
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

	// initialise cache via Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	status := redisClient.Ping()
	fmt.Println(status)

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)
	collectionUsers := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)

}
