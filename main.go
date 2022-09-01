package main

import (
	"go-fiber-crm/database"
	"go-fiber-crm/lead"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setUpRoutes(app *fiber.App){
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil{
		panic("failed to connect to the database.")
	}
	fmt.Println("connection opened to the database....")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated......")
}

func main(){
	app := fiber.New()
	initDatabase()

	setUpRoutes(app)
	app.Listen(2022)
	defer database.DBConn.Close()

}

// right now my build attempts are failing. something to do with the sql database.

