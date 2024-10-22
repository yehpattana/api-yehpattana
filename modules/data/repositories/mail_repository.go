package repositories

import (
	"fmt"

	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
	maildto "github.com/natersland/b2b-e-commerce-api/modules/mail/mail_dto"
	"gorm.io/gorm"
)

type SMTPRepository interface {
	SendEmail(email *maildto.Email) error
}

// type SMTPRepositoryImpl struct {
// 	cfg configs.Config
// }

// func NewSMTPRepository(cfg configs.Config) SMTPRepository {
// 	return &SMTPRepositoryImpl{cfg: cfg}
// }

type SMTPRepositoryImpl struct {
	DB *gorm.DB
}

func NewSMTPRepository(db *gorm.DB) SMTPRepository {
	return &SMTPRepositoryImpl{
		DB: db,
	}
}
func (r *SMTPRepositoryImpl) SendEmail(email *maildto.Email) error {
	var user entities.User

	userData := r.DB.Table("Users").Where("email = ?", email.To).Find(&user)
	if userData.Error != nil {
		return userData.Error
	}
	if user.Id == "" {
		return fmt.Errorf("cannot find user")
	}

	resetLink := "https://b2b.ymtinnovation.com//resetPassword/" + user.Id
	emailBody := fmt.Sprintf(`Click on the following link to reset your password:<br><a href="%s">%s</a>`, resetLink, resetLink)

	query := `
	EXEC msdb.dbo.sp_send_dbmail 
		@profile_name = 'FPSMail',
		@recipients = ?,
		@subject = 'Password Reset',
		@body = ?,
		@body_format = 'HTML';
	`

	// Execute the query with parameters
	err := r.DB.Exec(query, email.To, emailBody).Error
	if err != nil {
		return fmt.Errorf("failed to execute stored procedure: %v", err)
	}

	return nil
}
