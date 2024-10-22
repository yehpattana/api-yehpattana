package logger

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	utilshelpers "github.com/natersland/b2b-e-commerce-api/utils/utils_helpers"
)

type LoggerInterface interface {
	Print() LoggerInterface
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

func InitLogger(c *fiber.Ctx, res any, code int) LoggerInterface {
	log := &logger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		HttpMethod: c.Method(),
		StatusCode: code,
		Path:       c.Path(),
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(c)
	return log
}

type logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	HttpMethod string `json:"http_method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func (l *logger) Print() LoggerInterface {
	utilshelpers.Debug(l)
	return l
}

func (l *logger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser error: %v", err)
	}
	l.Query = body
}

func (l *logger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser error: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "customer register request"
	case "v1/users/admin/signup":
		l.Body = "admin register request"
	case "v1/users/signin":
		l.Body = "user sign in request"
	case "v1/users/admin/signin":
		l.Body = "admin sign in request"
	default:
		l.Body = body
	}
}

func (l *logger) SetResponse(res any) {
	l.Response = res
}
