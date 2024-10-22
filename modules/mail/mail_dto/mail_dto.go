package maildto

type Email struct {
	To string
}

// SMTPRepository is an interface for sending emails
type SMTPRepository interface {
	SendEmail(email *Email) error
}
