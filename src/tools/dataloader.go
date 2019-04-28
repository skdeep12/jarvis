package main

import (
	"context"
	"log"
	"os"
	"strings"
	"fmt"
	"github.com/extrame/xls"
	"github.com/skdeep12/jarvis/config"
	"github.com/skdeep12/jarvis/db"
	"github.com/skdeep12/jarvis/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func listFiles(dir string) ([]os.FileInfo, error) {
	f, err := os.Open(dir)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func insertOne(ctx context.Context, collection *mongo.Collection, dirName, fileName string) {
	xlFile, err := xls.Open(dirName+"/"+fileName, "UTF-8")
	if err != nil {
		log.Println(err)
		log.Println("Error opening in file " + fileName)
	}
	numberOfSheets := xlFile.NumSheets()
	logger.Debug.Println(fmt.Sprintf("%d sheet(s) is/are in %s.",numberOfSheets, fileName))
	for i := 0; i < numberOfSheets; i++ {
		sheet := xlFile.GetSheet(i)
		var v bson.M
		if sheet != nil {
			v = make(bson.M)
			var j uint16

			logger.Debug.Println(fmt.Sprintf("%d row(s) is/are in %s.",sheet.MaxRow, sheet.Name))
			for j = 0; j < sheet.MaxRow; j++ {
				logger.Debug.Println(sheet.Row(int(j)).Col(0) + " : " + sheet.Row(int(j)).Col(1))
				v[sheet.Row(int(j)).Col(0)] = sheet.Row(int(j)).Col(1)
			}
			v["Security Code"] = strings.Split(fileName, "_")[0]
			collection.InsertOne(ctx, v)
		}
	}
}

func insertAll(ctx context.Context, database *mongo.Database, collection, dirName string) {
	coll := database.Collection(collection)
	files, err := listFiles(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		insertOne(ctx, coll, dirName, file.Name())
	}
}

func main() {
	logger.Init(os.Stdout, os.Stdout)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, config.ProtoKey, config.MongoPort)
	ctx = context.WithValue(ctx, config.IPKey, config.MongoIP)
	ctx = context.WithValue(ctx, config.PortKey, config.MongoPort)

	db, err := db.GetDB(ctx, "finance")
	if err != nil {
		log.Fatal("Failure")
	}
	insertAll(ctx, db, "basicRatios", os.Args[1])
}
