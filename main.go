package main

import (
	"OrdentTest/config"
	"OrdentTest/config/migrations"
	"OrdentTest/controllers"
	_ "OrdentTest/docs"
	"OrdentTest/repositories"
	"OrdentTest/routes"
	"OrdentTest/services"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Fiber instance
	// docs.SwaggerInfo.Title = "OrdentTest ( Article Online )"
	// docs.SwaggerInfo.Description = "This is a simple article online API"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:3000"
	// docs.SwaggerInfo.BasePath = "/"
	app := fiber.New()
	app.Get("/swagger/*", fiberSwagger.WrapHandler) // Use the fiber-swagger handler
	config.DatabaseInit()
	migrations.Migrate()

	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	articleRepo := repositories.NewArticleRepository(config.DB)
	articleService := services.NewArticleService(articleRepo)
	articleController := controllers.NewArticleController(articleService)
	commetRepo := repositories.NewCommentRepository(config.DB)
	commentService := services.NewCommentService(commetRepo)
	commentController := controllers.NewCommentController(commentService)
	routes.RouteInit(app, userController, articleController, commentController)
	app.Listen(":3000")
}
