package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/delemike/recipe_api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"log/slog"
	"net/http"
)

type ProfileHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewProfileHandler(ctx context.Context, collection *mongo.Collection) *ProfileHandler {
	return &ProfileHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *ProfileHandler) CreateUserProfile(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h := sha256.New()
	h.Write([]byte(user.Password))
	passwordHash := hex.EncodeToString(h.Sum(nil))
	user.Password = passwordHash

	log.Printf("Attempting to insert user: %+v", user)
	_, err := handler.collection.InsertOne(handler.ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + "\nError while creating a new user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "username": user.Username})
}

func (handler *ProfileHandler) GetUserProfile(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error("Invalid ID format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user models.User

	err = handler.collection.FindOne(c, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {

		if errors.Is(mongo.ErrNoDocuments, err) {
			// Handle case where no document is found
			slog.Warn("User not found with ID:", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			// Handle other potential errors
			slog.Error("Error fetching user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User found", "user": user})

}
