package handlers

import (
	"encoding/json"
	"errors"
	"github.com/delemike/recipe_api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type RecipesHandler struct {
	collection  *mongo.Collection
	ctx         context.Context
	redisClient *redis.Client
}

// NewRecipesHandler Initialise the key resources
func NewRecipesHandler(
	ctx context.Context,
	collection *mongo.Collection,
	redisClient *redis.Client,
) *RecipesHandler {
	return &RecipesHandler{
		collection:  collection,
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (handler *RecipesHandler) NewRecipeHandler(c *gin.Context) {

	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = primitive.NewObjectID()
	recipe.PublishedAt = time.Now().String()
	_, err := handler.collection.InsertOne(c, recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + "\nError while inserting a new recipe"})
		return
	}
	log.Println("Remove data from Redis")
	handler.redisClient.Del("recipes")
	c.JSON(http.StatusOK, recipe)
}

func (handler *RecipesHandler) ListRecipesHandler(c *gin.Context) {
	recipes, err := loadAllRecipes(c, handler)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, recipes)
}

func (handler *RecipesHandler) GetRecipeByIDHandler(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}
	var recipe models.Recipe

	val, err := handler.redisClient.Get("recipes").Result()
	if errors.Is(err, redis.Nil) {
		log.Printf("Request to Mongo")

		err = handler.collection.FindOne(c, bson.M{"_id": objectID}).Decode(&recipe)
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
	} else {
		log.Printf("Request to Redis")
		var recipes []models.Recipe
		err = json.Unmarshal([]byte(val), &recipes)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not load all recipes"})
		}

		for _, recipe := range recipes {
			if recipe.ID == objectID {
				c.JSON(http.StatusOK, gin.H{"message": "Recipe found", "recipe": recipe})
				return
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe found", "recipe": recipe})

}

func (handler *RecipesHandler) UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := handler.collection.UpdateOne(c, bson.M{"_id": objectID}, bson.D{
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

	log.Println("Remove data from Redis")
	handler.redisClient.Del("recipes")
	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been updated", "recipe": recipe})
}

func (handler *RecipesHandler) UpdateRecipeByPatchHandler(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	var recipeFromPayload models.Recipe
	if err := c.ShouldBindJSON(&recipeFromPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recipe models.Recipe
	err = handler.collection.FindOne(c, bson.M{"_id": objectID}).Decode(&recipe)

	updateRecipe(recipeFromPayload, &recipe)

	_, err = handler.collection.UpdateOne(c, bson.M{"_id": objectID}, bson.D{
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

	log.Println("Remove data from Redis")
	handler.redisClient.Del("recipes")
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe updated successfully",
		"recipe":  recipe,
	})
}

func (handler *RecipesHandler) DeleteRecipeHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Perform the deletion
	_, err = handler.collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Remove data from Redis")
	handler.redisClient.Del("recipes")
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe deleted successfully",
	})
}

func (handler *RecipesHandler) SearchRecipeHandler(c *gin.Context) {
	tag := c.Query("tag")
	listOfRecipes := make([]models.Recipe, 0)
	recipes, err := loadAllRecipes(c, handler)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

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

func loadAllRecipes(c *gin.Context, handler *RecipesHandler) ([]models.Recipe, error) {
	val, err := handler.redisClient.Get("recipes").Result()
	var recipes []models.Recipe
	if errors.Is(err, redis.Nil) {
		log.Printf("Request to MongoDB")
		cur, err := handler.collection.Find(c, bson.D{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return []models.Recipe{}, errors.New(err.Error())
		}

		defer func(cur *mongo.Cursor, ctx context.Context) {
			_ = cur.Close(ctx)
		}(cur, c)

		recipes = make([]models.Recipe, 0)
		for cur.Next(c) {
			var recipe models.Recipe
			err := cur.Decode(&recipe)
			if err != nil {
				slog.Error(err.Error())

				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return []models.Recipe{}, errors.New(err.Error())
			}

			recipes = append(recipes, recipe)

		}
		data, _ := json.Marshal(recipes)
		handler.redisClient.Set("recipes", string(data), 0)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return []models.Recipe{}, errors.New(err.Error())
	} else {
		log.Printf("Request to Redis")
		recipes = make([]models.Recipe, 0)
		err = json.Unmarshal([]byte(val), &recipes)
		if err != nil {
			return nil, err
		}
	}

	return recipes, nil
}

func updateRecipe(recipeFromPayload models.Recipe, recipeToUpdate *models.Recipe) {
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
