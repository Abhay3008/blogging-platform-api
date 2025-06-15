package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func InitDbConnection() *mongo.Database {

	client, _ := mongo.Connect(options.Client().ApplyURI("mongodb://user:pass123@localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	fmt.Println("Connected to MongoDB!")
	db := client.Database("testdb")
	// collection := db.Collection("testcollection")
	return db
}

func InsertDocument(db *mongo.Database, collname string, doc blogpost) bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	collection := db.Collection(collname)
	insertResult, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	// objectID, ok := insertResult.InsertedID.(bson.ObjectID)
	// if !ok {
	// 	log.Fatal("Error while fetching doc ID")
	// }
	// id := objectID.Hex()

	var result bson.M
	err = collection.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		log.Fatal("Could not find inserted document:", err)
	}

	fmt.Printf("✅ Retrieved document: %+v\n", result)

	return result

}

func UpdateDocument(db *mongo.Database, collectionName string, filter bson.M, doc blogpost) bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	collection := db.Collection(collectionName)
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	update := bson.M{
		"$set": bson.M{
			"title":    doc.Title,
			"content":  doc.Content,
			"category": doc.Category,
			"tags":     doc.Tags,
		},
	}

	var updatedDoc bson.M
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		log.Fatalf("FindOneAndUpdate failed: %v", err)
	}

	fmt.Printf("✅ Updated document: %+v\n", updatedDoc)
	return updatedDoc

}
