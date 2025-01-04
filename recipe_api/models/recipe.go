package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
