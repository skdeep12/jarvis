package main

import (
	"fmt"
	"context"
	"log"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/skdeep12/jarvis/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/skdeep12/jarvis/handlers"
)


func some(c echo.Context) error {
	// var info map[string]string
	// info["P/E"] = "1"
	// info["EPS"] = "1"
	// info["RoE"] = "1"
	//c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	log.Println("param: " + c.Param("securityCode"))
	type ColorGroup struct {
		ID     int `json:"ID"`
		Name   string `json:"Name"`
		Colors []string `json:"Colors"`
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, _ := json.Marshal(group)
	log.Println(b)
	return c.JSON(200,group)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, config.ProtoKey, config.MongoProto)
	ctx = context.WithValue(ctx, config.IPKey, config.MongoIP)
	ctx = context.WithValue(ctx, config.PortKey, config.MongoPort)

	uri := fmt.Sprintf(`%s://%s:%s`,
		ctx.Value(config.ProtoKey).(string),
		ctx.Value(config.IPKey).(string),
		ctx.Value(config.PortKey).(string),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
	  fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	 	return
	}
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	if err != nil {
		fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
		return
	}
	if err != nil {
		log.Println(err)
		log.Fatal("Failure")
	}
	//insertAll(ctx, db)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//mux := http.NewServeMux()
	//e.GET("/", hello)
	ctx = context.WithValue(ctx, "client", client)
	e.GET("/_/companies", handlers.GetAllCompanies(ctx, "finance", "listOfScrips"))
	e.GET("/company/:securityCode",handlers.GetBasicInfo(ctx,"finance", "basicRatios"))
	for _,r := range e.Routes() {
		log.Println(r)
	}
	log.Println("Listening on port 3000")
	e.Logger.Fatal(e.Start(":3000"))
}
