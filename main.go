package main

import (
	"log"

	conf "github.com/erjiridholubis/go-superindo-product/internal/config"
	"github.com/erjiridholubis/go-superindo-product/internal/repository"
	"github.com/erjiridholubis/go-superindo-product/internal/service"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
	httpHandler "github.com/erjiridholubis/go-superindo-product/internal/deliveries"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Go Superindo API Product
// @description This is a sample server for Lion Superindo API Product Service.
// @BasePath /api/v1

// @schemes  http
// @host 127.0.0.1:3000
// @Version 1.0.0
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
	productRepo := repository.NewPostgreProductRepository(db)

	// Initialize service
	productService := service.NewProductService(productRepo)

	// Initialize handler
	apiProduct := pathApi.Group("/products")
	httpHandler.NewProductHandler(apiProduct, productService)


	log.Fatal(app.Listen(":3000"))
}