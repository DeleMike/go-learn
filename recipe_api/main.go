package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"os"
	"time"
)

var recipes []Recipe

func main() {
	Init()
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipe/:id", UpdateRecipeHandler)
	router.PATCH("/recipe/:id", UpdateRecipeByPatchHandler)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func Init() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	_ = json.Unmarshal(file, &recipes)
}

type Recipe struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt  string   `json:"published_at"`
}

func NewRecipeHandler(c *gin.Context) {

	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now().Format(time.RFC3339)
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func UpdateRecipeByPatchHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}
	foundRecipe := recipes[index]
	updateRecipe(recipe, &foundRecipe)
	recipes[index] = foundRecipe
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe updated successfully",
		"recipe":  foundRecipe,
	})
}

func updateRecipe(recipe Recipe, foundRecipe *Recipe) {
	if recipe.Name != "" {
		foundRecipe.Name = recipe.Name
	}
	if recipe.Tags != nil {
		foundRecipe.Tags = recipe.Tags
	}
	if recipe.Ingredients != nil {
		foundRecipe.Ingredients = recipe.Ingredients
	}
	if recipe.Instructions != nil {
		foundRecipe.Instructions = recipe.Instructions
	}
	if recipe.PublishedAt != "" {
		foundRecipe.PublishedAt = time.Now().Format(time.RFC3339)
	}
}
