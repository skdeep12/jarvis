package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skdeep12/jarvis/db"
	"github.com/skdeep12/jarvis/logger"
	"github.com/skdeep12/jarvis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
//GetAllCompanies returns a Handler function for echo package
//which returns all companes from a collection. 
func GetAllCompanies(ctx context.Context, database, coll string) echo.HandlerFunc {
	return func(c echo.Context) error {
		collection := ctx.Value("client").(*mongo.Client).Database(database).Collection(coll)
		docs, err := db.ReadAll(ctx, collection)
		if err != nil {
			log.Println(err)
			return nil
		}
		var companies []models.Company
		var company models.Company
		for _, cmp := range docs {
			company = models.Company{
				Name:        cmp["Security Name"].(string),
				Description: cmp["Industry"].(string),
			}
			companies = append(companies, company)
		}
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
		return c.JSON(http.StatusOK, companies)
	}
}
/*GetBasicInfo returns handler function for echo, which returns basic info about a company.
Handler expects a parameter in the url, which is a security code for that company.
*/
func GetBasicInfo(ctx context.Context, database, coll string) echo.HandlerFunc {
	return func(c echo.Context) error {
		securityCode := c.Param("securityCode")
		logger.Debug.Println("request for " + securityCode)
		collection := ctx.Value("client").(*mongo.Client).Database(database).Collection(coll)
		logger.Debug.Println("Reachd to basic ratios router.")
		info, err := db.GetOne(ctx, collection, bson.D{{Key: "Security Code", Value: securityCode}})
		if err != nil {
			//logger.Info.Println("Error in getting company basic info for " + securityCode)
			return c.JSON(http.StatusNotFound, err)
		}
		if info == nil {
			logger.Debug.Println("No record found.")
			return c.JSON(http.StatusNotFound, "")
		}
		// basicInfo := models.BasicRatiosCompany{
		// 	PE:  info["P/E"].(string),
		// 	EPS: info["EPS"].(string),
		// 	RoE: info["RoE"].(string),
		// }
		basicInfo := models.BasicRatiosCompany{
			PE:  "1",
			EPS: "1",
			RoE: "1",
		}
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
		return c.JSON(http.StatusOK, basicInfo)
	}
}

func TestHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// var info map[string]string
		// info["PE"] = "1"
		// info["EPS"] = "1"
		// info["RoE"] = "1"
		basicInfo := models.BasicRatiosCompany{
			PE:  "1",
			EPS: "1",
			RoE: "1",
		}

		//c.Response().Header().Add("Access-Control-Allow-Origin", "*")
		log.Println("param: " + c.Param("securityCode"))
		type ColorGroup struct {
			ID     int      `json:"ID"`
			Name   string   `json:"Name"`
			Colors []string `json:"Colors"`
		}
		_ = ColorGroup{
			ID:     1,
			Name:   "Reds",
			Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		}
		return c.JSON(200, basicInfo)
	}
}
