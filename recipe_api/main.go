package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

type Recipe struct {
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	PublishedAt string   `json:"published_at"`
}
