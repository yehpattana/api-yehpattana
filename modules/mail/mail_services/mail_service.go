package mailservices

import (
	"github.com/yehpattana/api-yehpattana/configs"
	"github.com/yehpattana/api-yehpattana/modules/data/repositories"
	maildto "github.com/yehpattana/api-yehpattana/modules/mail/mail_dto"
)

type EmailService interface {
	SendEmail(email *maildto.Email) error
}

type EmailServiceImpl struct {
	config    configs.ConfigInterface
	emailRepo repositories.SMTPRepository
}

func NewEmailService(cfg configs.ConfigInterface, er repositories.SMTPRepository) EmailService {
	return &EmailServiceImpl{config: cfg, emailRepo: er}
}

// func (u *EmailServiceImpl) GetUserIdByEmail(c *fiber.Ctx) error {
// 	userEmail := c.Params("email")

// 	result, err := u.emailRepo.GetUserIdByEmail(userEmail)
// 	if err != nil {
// 		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "cannot find user", err.Error()).Res()
// 	}

// 	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
// }

func (uc *EmailServiceImpl) SendEmail(email *maildto.Email) error {
	return uc.emailRepo.SendEmail(email)
}
