package lead

import (
	"go-fiber-crm/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct{
	gorm.Model 
	Name string				`json:"name"`
	Company string			`json:"company"`
	Email string 			`json:"email"`
	Phone int				`json:"phone"`
} // we also tell Golang what we want our struct to look like in JSON, hence the backticks.
// unlike Javascript, Golang does not have native JSON support.

func GetLeads(context *fiber.Ctx){
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	context.JSON(leads)
}

func GetLead(context *fiber.Ctx){
	id := context.Params("id")
	db := database.DBConn
	var lead Lead 
	db.Find(&lead, id)
	context.JSON(lead)
}

func NewLead(context *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err:= context.BodyParser(lead); err != nil{
		context.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	context.JSON(lead)
}

func DeleteLead(context *fiber.Ctx){
	id := context.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{
		context.Status(500).Send("no lead found with ID.")
		return
	}

	db.Delete(&lead)
	context.Send("successfully deleted.....")
}