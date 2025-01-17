package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"recipes": recipes,
	})
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	Picture     string       `json:"imageURL"`
}
type Ingredient struct {
	Quantity string `json:"quantity"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("/Users/mac/SWE/go-learn/recipe_api/recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)

	recipess := make([]Recipe, 0)
	for idx, recipe := range recipes {
		if idx%2 == 0 {
			recipe.Picture = "assets/images/burger.jpg"
		} else {

			recipe.Picture = "assets/images/oatmeal-cookies.jpg"
		}
		recipess = append(recipess, recipe)
	}

	// reset recipes to have the contents of recipesss
	recipes = recipess
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	err := router.Run()
	if err != nil {
		return
	}
}
