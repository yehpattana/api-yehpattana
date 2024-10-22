package servers

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/natersland/b2b-e-commerce-api/configs"
	"gorm.io/gorm"
)

type ServerInterface interface {
	Start()
	GetServer() *server
}

type server struct {
	app        *fiber.App
	config     configs.ConfigInterface
	configMail configs.Config
	db         *gorm.DB
}

func NewServer(config configs.ConfigInterface, configMail configs.Config, db *gorm.DB) ServerInterface {
	return &server{
		app: fiber.New(fiber.Config{
			AppName:      config.Service().Name(),
			BodyLimit:    100 * 1024 * 1024,
			ReadTimeout:  config.Service().ReadTimeout(),
			WriteTimeout: config.Service().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
		config:     config,
		configMail: configMail,
		db:         db,
	}
}

func (s *server) Start() {
	// middleware
	middleware := InitMiddlewares(s)
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Cors())

	// Modules
	v1 := s.app.Group("/v1")
	modules := InitModule(v1, s, middleware)

	modules.MonitorModule()
	modules.SwaggerModule()
	modules.AuthModule()
	modules.CompanyModule()
	modules.UserModule()
	modules.ProductModule()
	modules.TestModule()
	modules.ConfigMenuModule()
	modules.OrderModule()
	modules.ColorModule()
	modules.SizeModule()
	modules.LogModule()
	modules.MailModule()
	// modules.PaypalModule()
	s.app.Use(middleware.RouterCheck())
	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down...")
		_ = s.app.Shutdown()
	}()

	// Listen to host:port
	log.Printf("⚡️server is starting on %v", s.config.Service().Url())

	// Run server.
	if err := s.app.Listen(s.config.Service().Url()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

func (s *server) GetServer() *server {
	return s
}
