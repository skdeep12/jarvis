package db

import (
	"fmt"
	"log"
	"testing"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/skdeep12/jarvis/config"
)

func TestReadAll(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	uri := fmt.Sprintf(`%s://%s:%s`,
		config.MongoProto,
		config.MongoIP,
		config.MongoPort,
	)
	fmt.Println("Connecting to " + uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
	  t.Errorf("todo: couldn't connect to mongo: %v", err)
	 	return
	}
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		t.Errorf("todo: mongo client couldn't connect with background context: %v", err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	collection := client.Database("finance").Collection("basicRatios")
	fmt.Println(collection)
	result, err := ReadAll(ctx,collection)
	if err!=nil {
		t.Errorf("error in fetching record from collection " + collection.Name())
	}
	fmt.Println(result)
}

func TestGetOne(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	uri := fmt.Sprintf(`%s://%s:%s`,
		config.MongoProto,
		config.MongoIP,
		config.MongoPort,
	)
	t.Log("Connecting to " + uri)
	log.Println("Connecting to " + uri)
	fmt.Println("Connecting to " + uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
	  t.Errorf("todo: couldn't connect to mongo: %v", err)
	 	return
	}
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		t.Errorf("todo: mongo client couldn't connect with background context: %v", err)
		return
	}
	collection := client.Database("finance").Collection("basicRatios")
	t.Log(collection)
	result, err := GetOne(ctx,collection,bson.D{{Key: "Security Code", Value: "500180"}})
	if err!=nil {
		t.Errorf("error in fetching record from collection " + collection.Name())
	}
	t.Log(result)
	log.Println(result)
	fmt.Println(result)
}