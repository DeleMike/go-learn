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
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
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
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.PATCH("/recipes/:id", UpdateRecipeByPatchHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("recipes/search", SearchRecipeHandler)
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

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
// '200':
// description: Successful operation
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

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
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

	recipes = append(recipes[:index], recipes[index+1:]...)
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
