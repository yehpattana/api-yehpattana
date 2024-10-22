package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	maildto "github.com/yehpattana/api-yehpattana/modules/mail/mail_dto"
	mailservices "github.com/yehpattana/api-yehpattana/modules/mail/mail_services"
)

type EmailController interface {
	SendEmail(c *fiber.Ctx) error
}

type EmailControllerImpl struct {
	cfg          configs.Config
	emailService mailservices.EmailService
}

func NewEmailController(cfg configs.Config, emailService mailservices.EmailService) EmailController {
	return &EmailControllerImpl{cfg: cfg, emailService: emailService}
}

func (h *EmailControllerImpl) SendEmail(c *fiber.Ctx) error {
	email := new(maildto.Email)

	if err := c.BodyParser(email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if err := h.emailService.SendEmail(email); err != nil {
		msg := fmt.Sprintf("Failed to send email %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msg,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
	})
}
