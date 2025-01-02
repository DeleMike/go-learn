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
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

var recipes []Recipe

func main() {
	Init()
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.GET("/recipes/:id", GetRecipeByIDHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.PATCH("/recipes/:id", UpdateRecipeByPatchHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("recipes/search", SearchRecipeHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection

func Init() {
	// Initialize context
	ctx = context.Background()
	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create the MongoDB client
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	var listOfRecipes []interface{}

	for _, recipe := range recipes {
		listOfRecipes = append(listOfRecipes, recipe)
	}

	collection = client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

}

// swagger:parameters recipes newRecipe
type Recipe struct {
	//swagger:ignore
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Tags         []string           `json:"tags" bson:"tags"`
	Ingredients  []string           `json:"ingredients" bson:"ingredients" `
	Instructions []string           `json:"instructions" bson:"instructions"`
	PublishedAt  string             `json:"published_at" son:"published_at,omitempty"`
}

func NewRecipeHandler(c *gin.Context) {

	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = primitive.NewObjectID()
	recipe.PublishedAt = time.Now().String()
	_, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + "\nError while inserting a new recipe"})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
// '200':
// description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer func(cur *mongo.Cursor, ctx context.Context) {
		_ = cur.Close(ctx)
	}(cur, ctx)

	recipes = make([]Recipe, 0)
	for cur.Next(ctx) {
		var recipe Recipe
		err := cur.Decode(&recipe)
		if err != nil {
			slog.Error(err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		recipes = append(recipes, recipe)
	}
	c.JSON(http.StatusOK, recipes)
}

func GetRecipeByIDHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}
	var recipe Recipe
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&recipe)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			// Handle case where no document is found
			slog.Warn("Recipe not found with ID:", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		} else {
			// Handle other potential errors
			slog.Error("Error fetching recipe:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching recipe"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe found", "recipe": recipe})

}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{
		{"$set", bson.D{
			{"name", recipe.Name},
			{"instructions", recipe.Instructions},
			{"ingredients", recipe.Ingredients},
			{"tags", recipe.Tags},
		}},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been updated", "recipe": recipe})
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
//
//	'200':
//	  description: Successful operation
//	'400':
//	  description: Invalid input
//	'404':
//	  description: Invalid recipe ID
func UpdateRecipeByPatchHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	var recipeFromPayload Recipe
	if err := c.ShouldBindJSON(&recipeFromPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe Recipe
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&recipe)

	updateRecipe(recipeFromPayload, &recipe)

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{
		{"$set", bson.D{

			{"name", recipe.Name},
			{"instructions", recipe.Instructions},
			{"ingredients", recipe.Ingredients},
			{"tags", recipe.Tags},
		}},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe updated successfully",
		"recipe":  recipe,
	})
}

func updateRecipe(recipeFromPayload Recipe, recipeToUpdate *Recipe) {
	if recipeFromPayload.Name != "" {
		recipeToUpdate.Name = recipeFromPayload.Name
	}
	if recipeFromPayload.Tags != nil {
		recipeToUpdate.Tags = recipeFromPayload.Tags
	}
	if recipeFromPayload.Ingredients != nil {
		recipeToUpdate.Ingredients = recipeFromPayload.Ingredients
	}
	if recipeFromPayload.Instructions != nil {
		recipeToUpdate.Instructions = recipeFromPayload.Instructions
	}
	//if recipe.PublishedAt != "" {
	//	foundRecipe.PublishedAt = time.Now().Format(time.RFC3339)
	//}

}

func DeleteRecipeHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Perform the deletion
	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe deleted successfully",
	})
}

func SearchRecipeHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]Recipe, 0)

	// very bad search...LOL. We can do better...
	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.ToLower(t) == strings.ToLower(tag) {
				found = true
			}
		}

		if found {
			listOfRecipes = append(listOfRecipes, recipes[i])

		}
	}

	c.JSON(http.StatusOK, listOfRecipes)
}
