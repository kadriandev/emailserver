package main

import (
	"fmt"
	"log"

	"emailserver/database"
	"emailserver/mail"
	"emailserver/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var (
	Config ConfigVars
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	mail.AddRouter("/mail", api)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "emailservice.db")
	if err != nil {
		panic("Error connecting to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&mail.MailTemplate{}, &user.User{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	// var err error
	// Config, err := LoadConfig(".")

	initDatabase()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
