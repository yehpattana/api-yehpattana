package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	middlewarecontroller "github.com/natersland/b2b-e-commerce-api/middleware/middleware_controller"
	middlewarerepositories "github.com/natersland/b2b-e-commerce-api/middleware/middleware_repositories"
	middlewareservices "github.com/natersland/b2b-e-commerce-api/middleware/middleware_services"
	authcontrollers "github.com/natersland/b2b-e-commerce-api/modules/auth/auth_controllers"
	authservices "github.com/natersland/b2b-e-commerce-api/modules/auth/auth_services"
	colorcontrollers "github.com/natersland/b2b-e-commerce-api/modules/color/color_controllers"
	colorservices "github.com/natersland/b2b-e-commerce-api/modules/color/color_services"
	companiescontrollers "github.com/natersland/b2b-e-commerce-api/modules/companies/companies_controllers"
	companiesservices "github.com/natersland/b2b-e-commerce-api/modules/companies/companies_services"
	configmenucontrollers "github.com/natersland/b2b-e-commerce-api/modules/config_menu/config_menu_controllers"
	configmenuservices "github.com/natersland/b2b-e-commerce-api/modules/config_menu/config_menu_services"
	"github.com/natersland/b2b-e-commerce-api/modules/data/repositories"
	logcontrollers "github.com/natersland/b2b-e-commerce-api/modules/log/log_controllers"
	logservices "github.com/natersland/b2b-e-commerce-api/modules/log/log_services"
	mailcontrollers "github.com/natersland/b2b-e-commerce-api/modules/mail/mail_controllers"
	mailservices "github.com/natersland/b2b-e-commerce-api/modules/mail/mail_services"
	monitorcontrollers "github.com/natersland/b2b-e-commerce-api/modules/monitor/monitor_controllers"
	orderscontrollers "github.com/natersland/b2b-e-commerce-api/modules/orders/orders_controllers"
	orderservices "github.com/natersland/b2b-e-commerce-api/modules/orders/orders_services"
	productscontrollers "github.com/natersland/b2b-e-commerce-api/modules/products/products_controllers"
	productsservices "github.com/natersland/b2b-e-commerce-api/modules/products/products_services"
	sizecontrollers "github.com/natersland/b2b-e-commerce-api/modules/size/size_controllers"
	sizeservices "github.com/natersland/b2b-e-commerce-api/modules/size/size_services"

	// productscontrollers "github.com/natersland/b2b-e-commerce-api/modules/products/products_controllers"
	// productsservices "github.com/natersland/b2b-e-commerce-api/modules/products/products_services"
	userscontrollers "github.com/natersland/b2b-e-commerce-api/modules/users/users_controllers"
	usersservices "github.com/natersland/b2b-e-commerce-api/modules/users/users_services"
)

type ModuleInterface interface {
	SwaggerModule()
	MonitorModule()
	AuthModule()
	CompanyModule()
	UserModule()
	ProductModule()
	TestModule()
	ConfigMenuModule()
	OrderModule()
	ColorModule()
	SizeModule()
	LogModule()
	MailModule()
	// PaypalModule()
}

func InitModule(router fiber.Router, server *server, middleware middlewarecontroller.MiddlewareControllerInterface) ModuleInterface {
	return &module{
		router:     router,
		server:     server,
		middleware: middleware,
	}
}

type module struct {
	router     fiber.Router
	server     *server
	middleware middlewarecontroller.MiddlewareControllerInterface
}

func InitMiddlewares(s *server) middlewarecontroller.MiddlewareControllerInterface {
	repository := middlewarerepositories.MiddlewareRepositoryImpl(s.db)
	service := middlewareservices.MiddlewareServiceImpl(repository)
	return middlewarecontroller.MiddlewareController(s.config, service)
}

// @title Go Fiber Swagger Example API
// @version 1.0
// @description This is a sample server for a Fiber application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func (m *module) SwaggerModule() {
	router := m.router.Group("/swagger")

	// http://localhost:8080/v1/swagger
	router.Get("/*", swagger.HandlerDefault) // default

	router.Get("/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}

func (m *module) MonitorModule() {
	controller := monitorcontrollers.MonitorController(m.server.config)

	m.router.Get("/", controller.HealthCheck)
}

func (m *module) AuthModule() {
	authRepository := repositories.AuthRepositoryImpl(m.server.db)
	companyRepository := repositories.CompaniesRepositoryImpl(m.server.db)
	service := authservices.AuthServiceImpl(m.server.config, authRepository, companyRepository)
	controller := authcontrollers.AuthController(m.server.config, service)

	router := m.router.Group("/auth")
	// auth
	router.Post("/signup", controller.SignUpCustomer)
	router.Post("/signup/admin", controller.SignUpAdmin)
	router.Post("/signin", controller.SignInCustomer)
	router.Post("/signin/admin", controller.SignInAdmin)
	router.Post("/refresh", controller.RefreshCustomerPassport)
	router.Post("/refresh/admin", controller.RefreshAdminPassport)
	router.Delete("/signout", controller.SignOut)

	router.Get("/secret", m.middleware.JwtAuthenticator(), controller.GenerateAdminToken)

}

func (m *module) CompanyModule() {
	repository := repositories.CompaniesRepositoryImpl(m.server.db)
	service := companiesservices.CompaniesServiceImpl(m.server.config, repository)
	controller := companiescontrollers.CompaniesControllerImpl(m.server.config, service)

	router := m.router.Group("/companies")
	router.Get("/", controller.GetCompanies)
	router.Get("/:companyId", controller.GetCompanyDetail)
	router.Post("/create-company", controller.CreateCompany)
	router.Patch("/edit-company", controller.EditCompany)
	router.Delete("/delete-company", controller.DeleteCompany)
}

func (m *module) UserModule() {
	repository := repositories.UsersRepositoryImpl(m.server.db)
	service := usersservices.UserServiceImpl(m.server.config, repository)
	controller := userscontrollers.UserControllerImpl(m.server.config, service)

	router := m.router.Group("/users")
	router.Get("/", controller.GetAllCustomers)
	router.Get("/:userId", controller.GetCustomerDetail)
	router.Patch("change-password/:userId", controller.ChangePassword) // TODO auth required
	// TODO implement reset password by email
	router.Patch("/verify-user/:userId", controller.VerifyUser) // TODO auth required
	router.Patch("/ban-user/:userId", controller.BanUser)       // TODO auth required
	router.Patch("/update-customer/:userId", controller.UpdateCustomer)
	router.Delete("/remove-customer/:userId/:customerId", controller.DeleteCustomer)
}

func (m *module) ProductModule() {
	repository := repositories.ProductsRepositoryV2Impl(m.server.db, &m.server.config)
	service := productsservices.ProductServiceV2Impl(m.server.config, repository)
	controller := productscontrollers.ProductControllerV2Impl(m.server.config, service)

	router := m.router.Group("/products")
	// // common
	// router.Patch("/change-status", controller.ChangeProductStatus) // TODO auth required
	// router.Patch("/:productId", controller.UpdateProduct)          // TODO auth required
	// // router.Get("/:productId", controller.GetProductDetail)

	// // webview
	router.Get("/", controller.GetAllProducts)              // TODO auth required
	router.Get("/:masterCode", controller.GetProductDetail) // TODO auth required

	// // admin
	router.Get("/admin/check-product-exist/:master_code", controller.CheckIsProductExistByMasterCode)
	router.Post("/create-product", controller.CreateProduct)  // TODO auth required
	router.Post("/create-stock", controller.CreateStock)      // TODO auth required
	router.Patch("/update-stock", controller.UpdateStock)     // TODO auth required
	router.Patch("/decrease-stock", controller.DecreaseStock) // TODO auth required
	router.Get("/admin/get-cover-image/:master_code", controller.GetCoverImageByMasterCode)
	router.Get("/admin/get-size-chart/:master_code", controller.GetSizeChartByMasterCode)
	router.Post("/admin/upload-cover-image", controller.UploadCoverImage)
	router.Post("/admin/upload-size-chart", controller.UploadSizeChart)
	router.Patch("/admin/update-cover-image", controller.UpdateCoverImage)
	router.Patch("/admin/update-size-chart", controller.UpdateSizeChart)
	router.Patch("/admin/update-main-product-detail", controller.UpdateMainProductDetailByMasterCode)
	router.Patch("/admin/update-new-master-code", controller.UpdateMasterCode)
	router.Patch("/admin/update-product-variant", controller.UpdateProductVariant)
	// router.Post("/create-product-variant", controller.CreateProductVaraint) // TODO auth required
	// router.Post("/create-product-item", controller.CreateProductItem)       // TODO auth required
	router.Get("/all/admin", controller.GetAllAdminProducts)
	router.Get("/admin/:masterCode", controller.GetProductDetailAdmin)                                             // TODO auth required
	router.Delete("/admin/delete-all-product-in-master-code/:masterCode", controller.DeleteAllProductInMasterCode) // TODO auth required
	router.Delete("/admin/delete-product-variant/:productId", controller.DeleteProductVariant)                     // TODO auth required
	router.Get("admin/get-all-products-by-conmpany-id/:companyId", controller.GetAllProductsByCompany)

}

func (m *module) ConfigMenuModule() {
	repository := repositories.ConfigMenuRepositoryImpl(m.server.db)
	service := configmenuservices.ConfigMenuServiceImpl(m.server.config, repository)
	controller := configmenucontrollers.ConfigMenuControllerImpl(m.server.config, service)

	router := m.router.Group("/config-menu")
	router.Get("/:companyId", controller.GetAllConfigMenu)
	router.Post("/", controller.CreateConfigMenu)
	router.Put("/", controller.UpdateConfigMenu)
	router.Delete("/", controller.DeleteConfigMenu)
	router.Get("/sub-menu/:companyId", controller.GetAllConfigSubMenu)
	router.Post("/sub-menu", controller.CreateConfigSubMenu)
	router.Put("/sub-menu", controller.UpdateConfigSubMenu)
	router.Delete("/sub-menu", controller.DeleteConfigSubMenu)
}

func (m *module) TestModule() {
	router := m.router.Group("/test")
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	},
	)
}

func (m *module) OrderModule() {
	repository := repositories.OrdersRepositoryImpl(m.server.db)
	service := orderservices.OrderServiceImpl(m.server.config, repository)
	controller := orderscontrollers.OrdersControllerImpl(m.server.config, service)

	router := m.router.Group("/order")
	router.Get("/", controller.GetAllOrder)
	router.Get("/:customerId", controller.GetByCustomerId)
	router.Get("/:customerId/:orderId", controller.GetByOrderId)
	router.Delete("/:orderId", controller.DeleteByOrderId)
	router.Patch("/:orderId", controller.AttachPackingListByOrderId)
	router.Patch("/", controller.UpdateOrderTracking)
	router.Patch("/payment/paypal", controller.UpdateOrderPayment)
	router.Post("/", controller.CreateOrder)
}

func (m *module) ColorModule() {
	repository := repositories.ColorRepositoryImpl(m.server.db)
	service := colorservices.ColorServiceImpl(m.server.config, repository)
	controller := colorcontrollers.ColorControllerImpl(m.server.config, service)

	router := m.router.Group("/color")
	router.Get("/", controller.GetAllColor)
	router.Post("/", controller.CreateColor)
	router.Delete("/:colorId", controller.DeleteColor)
}
func (m *module) SizeModule() {
	repository := repositories.SizeRepositoryImpl(m.server.db)
	service := sizeservices.SizeServiceImpl(m.server.config, repository)
	controller := sizecontrollers.SizeControllerImpl(m.server.config, service)

	router := m.router.Group("/size")
	router.Get("/", controller.GetAllSize)
	router.Post("/", controller.CreateSize)
	router.Delete("/:sizeId", controller.DeleteSize)
}
func (m *module) LogModule() {
	repository := repositories.LogRepositoryImpl(m.server.db)
	service := logservices.LogServiceImpl(m.server.config, repository)
	controller := logcontrollers.LogControllerImpl(m.server.config, service)

	router := m.router.Group("/log")
	router.Get("/", controller.GetLog)
}

func (m *module) MailModule() {
	// Instantiate the repository, service, and controller

	emailRepository := repositories.NewSMTPRepository(m.server.db)
	emailService := mailservices.NewEmailService(m.server.config, emailRepository)
	emailController := mailcontrollers.NewEmailController(m.server.configMail, emailService)

	// Set up routes under "/email"
	router := m.router.Group("/email")
	router.Post("/send", emailController.SendEmail)
}

// func (m *module) PaypalModule() {
// 	repository := repositories.NewPayPalRepository(m.server.db)
// 	service := paypalservices.NewPayPalService(m.server.config, repository)

// 	// Initialize PayPal controller using the service
// 	controller := paypalcontrollers.NewPayPalController(m.server.config, service)

// 	// Set up the routes for the PayPal module
// 	router := m.router.Group("/paypal")

// 	// Define PayPal-specific routes, for example:
// 	router.Post("/create-order", controller.CreatePayPalOrderHandler) // Route for creating a PayPal order    // Route for getting PayPal order details
// }
