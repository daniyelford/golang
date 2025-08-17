package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// تابع اتصال به MongoDB
func connect(uri string, dbName string, collectionName string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

// تابع ایجاد سند (Create)
func createDocument(collection *mongo.Collection, doc map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
}

// تابع خواندن سند (Read)
func getDocument(collection *mongo.Collection, filter map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result map[string]interface{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found document:", result)
}

// تابع به‌روزرسانی سند (Update)
func updateDocument(collection *mongo.Collection, filter map[string]interface{}, update map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateData := bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Matched count:", result.MatchedCount, "Modified count:", result.ModifiedCount)
}

// تابع حذف سند (Delete)
func deleteDocument(collection *mongo.Collection, filter map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", result.DeletedCount)
}

func main() {
	// تنظیمات خارجی
	mongoURI := "mongodb://localhost:27017"
	dbName := "testdb"
	collectionName := "mycollection"

	collection := connect(mongoURI, dbName, collectionName)

	// مثال سند دلخواه
	doc := map[string]interface{}{
		"title":   "My Document",
		"content": "This is a test",
		"views":   100,
	}

	// Create
	createDocument(collection, doc)

	// Read
	getDocument(collection, map[string]interface{}{"title": "My Document"})

	// Update
	updateDocument(collection, map[string]interface{}{"title": "My Document"}, map[string]interface{}{"views": 150})

	// Read بعد از Update
	getDocument(collection, map[string]interface{}{"title": "My Document"})

	// Delete
	deleteDocument(collection, map[string]interface{}{"title": "My Document"})
}
