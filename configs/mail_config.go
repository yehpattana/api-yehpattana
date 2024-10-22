package configs

import (
	"os"
	"strconv"
)

type Config struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

func LoadConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	return Config{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     port,
		SMTPUser:     os.Getenv("SMTP_USER"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}
}
