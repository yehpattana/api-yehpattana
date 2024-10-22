package middlewarecontroller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/natersland/b2b-e-commerce-api/configs"
	middlewareservices "github.com/natersland/b2b-e-commerce-api/middleware/middleware_services"
	"github.com/natersland/b2b-e-commerce-api/modules/auth/pkg"
	commonresponse "github.com/natersland/b2b-e-commerce-api/modules/commons/common_response"
	commonstring "github.com/natersland/b2b-e-commerce-api/modules/commons/common_string"
)

type middlewareControllerErrorCode string

const (
	routerCheckError middlewareControllerErrorCode = "MIDWARE-001"
	jwtAuthError     middlewareControllerErrorCode = "MIDWARE-002"
	paramsCheckError middlewareControllerErrorCode = "MIDWARE-003"
	authorizeError   middlewareControllerErrorCode = "MIDWARE-004"
	apiKeyError      middlewareControllerErrorCode = "MIDWARE-005"
)

type MiddlewareControllerInterface interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
	JwtAuthenticator() fiber.Handler
}

func MiddlewareController(cfg configs.ConfigInterface, middlewareService middlewareservices.MiddlewareServiceInterface) MiddlewareControllerInterface {
	return &middlewareController{
		config:            cfg,
		middlewareservice: middlewareService,
	}
}

type middlewareController struct {
	config            configs.ConfigInterface
	middlewareservice middlewareservices.MiddlewareServiceInterface
}

func (c *middlewareController) Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	})
}

func (*middlewareController) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return commonresponse.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(routerCheckError),
			"route not found",
		).Res()
	}
}

func (*middlewareController) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Bangkok/Asia",
	})
}

func (m *middlewareController) JwtAuthenticator() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		result, err := pkg.ParseToken(m.config.Jwt(), token)
		if err != nil {
			return commonresponse.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(jwtAuthError),
				err.Error(),
			).Res()
		}
		claims := result.Claims
		if !m.middlewareservice.FindAccessToken(claims.Id, token) {
			return commonresponse.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(authorizeError),
				string(commonstring.Unauthorized),
			).Res()
		}

		// set the user id to the context
		c.Locals("user_id", claims.Id)
		c.Locals("user_role", claims.Role)
		return c.Next()
	}
}
