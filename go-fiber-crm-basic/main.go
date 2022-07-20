package main

import (
	"github.com/amrmetwaly/go-projects/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func initDatabase() {

}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)

}
