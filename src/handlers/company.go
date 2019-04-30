package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skdeep12/jarvis/db"
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
		collection := ctx.Value("client").(*mongo.Client).Database(database).Collection(coll)
		info, err := db.GetOne(ctx, collection, bson.D{{Key: "Security Code", Value: securityCode}})
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		if info == nil {
			return c.JSON(http.StatusNotFound, "")
		}
		basicInfo := models.BasicRatiosCompany{
			PE:  info["P/E"].(string),
			EPS: info["EPS"].(string),
			RoE: info["RoE"].(string),
		}
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
		return c.JSON(http.StatusOK, basicInfo)
	}
}
