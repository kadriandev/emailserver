package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type MailTemplate struct {
	gorm.Model
	Sender  string `json:"sender"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Mail struct {
	RecipientList string
	Template      MailTemplate
}

func AddRouter(path string, api fiber.Router) {
	router := api.Group(path)
	router.Get("/", GetAllMailTemplates)
	router.Get("/:templateId", GetMailTemplate)
	router.Post("", CreateMailTemplate)
	router.Put("/:templateId", UpdateMailTemplate)
	router.Delete("/:templateId", DeleteMailTemplate)
	router.Post("/send", SendTemplate)

}

func (m Mail) Send() {

	password := "ygfkevwteeqxtgsm"
	fmt.Printf("Username: %s Password: %s\n", m.Template.Sender, password)

	addr := "smtp.gmail.com:587"
	host := "smtp.gmail.com"

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", m.Template.Sender)

	msg += fmt.Sprintf("To: %s\r\n", m.RecipientList)
	msg += fmt.Sprintf("Subject: %s\r\n", m.Template.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", m.Template.Body)

	auth := smtp.PlainAuth("", m.Template.Sender, password, host)
	err := smtp.SendMail(addr, auth, m.Template.Sender, strings.Split(m.RecipientList, ","), []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

}
