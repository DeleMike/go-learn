package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

func main() {
	users := map[string]string{
		"admin":    "fCRmh4Q2J7Rseqkz",
		"akindele": "RE4zfHB35VPtTkbT",
		"michael":  "L3nSFRcZzNQ67bcc",
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	h := sha256.New()
	for username, password := range users {

		h.Write([]byte(password))
		passwordHash := hex.EncodeToString(h.Sum(nil)) // Convert hash to a hex string

		// Insert user into the database
		_, err := collection.InsertOne(ctx, bson.M{
			"username": username,
			"password": passwordHash,
		})
		if err != nil {
			log.Printf("Error adding user %s: %v", username, err)
			continue
		}

		log.Printf("User %s added successfully", username)
	}

	log.Println("All users added.")

}
