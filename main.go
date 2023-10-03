package main

import (
	"fmt"
	"log"

	conf "github.com/erjiridholubis/go-superindo-product/internal/config"
	"github.com/erjiridholubis/go-superindo-product/internal/middleware"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
	"github.com/erjiridholubis/go-superindo-product/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	httpHandler "github.com/erjiridholubis/go-superindo-product/internal/deliveries"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/erjiridholubis/go-superindo-product/docs"
)

// @title Go Superindo API Product
// @description This is a sample server for Lion Superindo API Product Service.
// @BasePath /api/v1

// @schemes  http
// @host 127.0.0.1:3000
// @Version 1.0.0

// @SecurityDefinitions.apikey  Authorization
// @in header
// @name Authorization
// @description This is a bearer token. Add 'Bearer ' before placing the token.
func main() {
	config, err := conf.InitConfig()
    if err != nil {
        log.Fatalf("Failed to initialize configuration: %v", err)
    }

	db, err := conf.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	app := fiber.New()
    app.Use(cors.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	pathApi := app.Group("/api/v1")
	
	// get health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "ok",
		})
	})

	// Initialize repository
	postgreRepo := repository.NewPostgreRepository(db)

	// Initialize service
	authService := service.NewAuthService(postgreRepo)
	productService := service.NewProductService(postgreRepo)
	categoryService := service.NewCategoryService(postgreRepo)
	cartItemService := service.NewCartItemService(postgreRepo)
	userService := service.NewUserService(postgreRepo)

	// Initialize handler
	apiAuth := pathApi.Group("/auth")
	httpHandler.NewAuthHandler(apiAuth, authService)

	// Using middleware
	pathApi.Use(middleware.JWTMiddleware())

	apiProduct := pathApi.Group("/products")
	httpHandler.NewProductHandler(apiProduct, productService)

	apiCategory := pathApi.Group("/categories")
	httpHandler.NewCategoryHandler(apiCategory, categoryService)

	apiCartItem := pathApi.Group("/cart-items")
	httpHandler.NewCartItemHandler(apiCartItem, cartItemService)

	apiUser := pathApi.Group("/users")
	httpHandler.NewUserHandler(apiUser, userService)
	
	httpAddr := config.Server.UserAddress
    fmt.Printf("Server started on %s\n", httpAddr)
    
    // Run server Fiber
    log.Fatal(app.Listen(httpAddr))
}