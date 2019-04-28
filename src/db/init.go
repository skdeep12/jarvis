package db

import (
	"context"
	"fmt"
	"github.com/skdeep12/jarvis/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetDB returns the mongo database instance from mongo server
func GetDB(ctx context.Context, database string) (*mongo.Database, error) {
	// find out what options is
	uri := fmt.Sprintf(`%s://%s:%s`,
		ctx.Value(config.ProtoKey).(string),
		ctx.Value(config.IPKey).(string),
		ctx.Value(config.PortKey).(string),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}
	_db := client.Database(database)
	return _db, nil
}

//ReadAll reads all the documents from a collection
func ReadAll(ctx context.Context, collection *mongo.Collection) ([]primitive.M, error) {
	//collection := db.Collection(coll)
	cursor, err := collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Println(err)
		log.Println("Error in fetching records from:: collection: " + collection.Name())
		return nil, err
	}
	var data []primitive.M
	for cursor.Next(ctx) {
		elem := &bson.D{}
		if e := cursor.Decode(elem); e != nil {
			log.Println("Error decoding element")
		}
		m := elem.Map()
		data = append(data, m)
	}
	return data, nil
}

func GetOne(ctx context.Context, collection *mongo.Collection,criteria bson.D) (primitive.M, error){
	cursor, err := collection.Find(ctx,criteria)
	if err!= nil {
		return nil,err
	}
	defer cursor.Close(ctx)
	var data primitive.M
	for cursor.Next(ctx) {
		elem := &bson.D{}
		if e := cursor.Decode(elem); e != nil {
			log.Println("Error decoding element")
		}
		data = elem.Map()
	}
	return data, nil
}
