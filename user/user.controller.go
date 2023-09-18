package user

import (
	"emailserver/auth"
	"emailserver/database"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *fiber.Ctx) error {
	db := database.DBConn
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user, existErr := getUserByEmail(req.Email)
	if existErr != nil {
		return c.Status(500).SendString("No account with that email exists.")
	}

	// Hashing the password with the default cost of 10
	var err error
	user.Password, err = auth.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	db.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	db := database.DBConn
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user := new(User)
	err := db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return c.Status(500).SendString("Could not find account with that email.")
	}

	isMatch := auth.CheckPasswordHash(req.Password, user.Password)
	if isMatch == false {
		return c.Status(500).SendString("Password does not match.")
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	hashSecret := os.Getenv("HASH_SECRET")
	tokenString, signErr := token.SignedString([]byte(hashSecret))
	if signErr != nil {
		c.Status(500).SendString("Error signing jwt token")
	}

	return c.JSON(LoginResponse{
		Token: tokenString,
	})
}
