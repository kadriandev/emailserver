package mail

import (
	"emailserver/database"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllMailTemplates(c *fiber.Ctx) error {
	db := database.DBConn
	var templates []MailTemplate
	db.Find(&templates)
	return c.JSON(templates)
}

func GetMailTemplate(c *fiber.Ctx) error {
	id := c.Params("templateId")
	db := database.DBConn
	var template MailTemplate
	db.Find(&template, id)
	return c.JSON(template)
}

func CreateMailTemplate(c *fiber.Ctx) error {
	db := database.DBConn
	template := new(MailTemplate)
	if err := c.BodyParser(template); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Create(&template)
	return c.JSON(template)
}

func UpdateMailTemplate(c *fiber.Ctx) error {
	db := database.DBConn
	template := new(MailTemplate)
	if err := c.BodyParser(template); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Save(&template)
	return c.JSON(template)
}

func DeleteAllMailTemplates(c *fiber.Ctx) error {
	db := database.DBConn
	db.Delete(&MailTemplate{})
	return c.Status(200).SendString("Deleted")
}

func DeleteMailTemplate(c *fiber.Ctx) error {
	db := database.DBConn
	id, err := strconv.Atoi(c.Params("templateId"))
	if err != nil {
		return c.Status(500).SendString("ID provided was not a number")
	}
	db.Delete(&MailTemplate{}, id)
	return c.Status(200).SendString(fmt.Sprintf("Deleted template at id: %d", id))
}

func SendTemplate(c *fiber.Ctx) error {
	db := database.DBConn
	req := new(struct {
		TemplateId uint   `json:"templateId"`
		Recipients string `json:"recipients"`
	})
	if err := c.BodyParser(req); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var mail Mail
	mail.RecipientList = req.Recipients
	db.Find(&mail.Template, req.TemplateId)
	mail.Send()

	return c.Status(200).SendString(fmt.Sprintf("Successfully sent message"))
}
