package user

import (
	"emailserver/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	Firstname string
	Lastname  string
	// Settings uint foreignkey to &UserSettings{}
	Settings UserSettings
	// MailTemplates uint foreinkey to &MailTemplate{}
}

type UserSettings struct {
	AppPassword string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string `json:"token"`
}

func AddRouter(path string, api fiber.Router) {
	router := api.Group(path)
	router.Post("/register", Register)
	router.Post("/login", Login)
}

func getUserByEmail(email string) (User, error) {
	db := database.DBConn
	var user User
	err := db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
