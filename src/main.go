package main

import (
	"context"
	"log"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/skdeep12/jarvis/config"
	"github.com/skdeep12/jarvis/db"
	"github.com/skdeep12/jarvis/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func configDB() (*mongo.Database, error){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, config.ProtoKey, config.MongoPort)
	ctx = context.WithValue(ctx, config.IPKey, config.MongoIP)
	ctx = context.WithValue(ctx, config.PortKey, config.MongoPort)
	return db.GetDB(ctx, "finance")
}

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

	// db, err := configDB()
	// if err != nil {
	// 	log.Fatal("Failure")
	// }
	//insertAll(ctx, db)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//mux := http.NewServeMux()
	//e.GET("/", hello)
	//e.GET("/_/companies", handlers.GetAllCompanies(ctx, db, "listOfScrips"))
	e.GET("/company/:securityCode",handlers.TestHandler())
	for _,r := range e.Routes() {
		log.Println(r)
	}
	log.Println("Listening on port 3000")
	e.Logger.Fatal(e.Start(":3000"))
}
