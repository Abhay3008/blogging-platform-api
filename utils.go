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
	// fmt.Printf("%T\n", db)

	// // doc := map[string]string{"name": "Abhay", "role": "DevOps"}
	// doc := blogpost{
	// 	Title:    "Test",
	// 	Content:  "test content",
	// 	Category: "Test",
	// 	Tags:     []string{"test1"},
	// }
	// insertResult, err := collection.InsertOne(ctx, doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted document ID:", insertResult.InsertedID)
}

func InsertDocument(db *mongo.Database, collname string, doc blogpost) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	collection := db.Collection(collname)
	insertResult, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document type:%T", insertResult.InsertedID)
	objectID, ok := insertResult.InsertedID.(bson.ObjectId)
	if !ok {
		log.Fatal("InsertedID is not of type primitive.ObjectID")
	}
	id := objectID.Hex()
	// id := insertResult.InsertedID.(primitive.ObjectID).Hex()

	var result bson.D
	err = collection.FindOne(ctx, bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	if err != nil {
		log.Fatal("Could not find inserted document:", err)
	}
	result.Id = id

	fmt.Printf("✅ Retrieved document: %+v\n", result)

}
