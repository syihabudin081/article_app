package routes

import (
	"OrdentTest/controllers"
	"OrdentTest/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App, userController *controllers.UserController, articleController *controllers.ArticleController, comentController *controllers.CommentController) {
	app.Post("/register", userController.Register)
	app.Post("/login", userController.Login)
	app.Post("/articles", middleware.AuthMiddleware, middleware.AuthorMiddleware, articleController.CreateArticle)
	app.Get("/articles", articleController.GetArticles)
	app.Get("/articles/:id", articleController.GetArticleByID)
	app.Put("/articles/:id", middleware.AuthMiddleware, articleController.UpdateArticle)
	app.Delete("/articles/:id", middleware.AuthMiddleware, middleware.AdminMiddleware, articleController.DeleteArticle)
	app.Post("/comments", middleware.AuthMiddleware, comentController.Create)
	app.Get("/comments", middleware.AuthMiddleware, comentController.Create)
	app.Get("/comments/:id", middleware.AuthMiddleware, comentController.GetByID)
	app.Put("/comments/:id", middleware.AuthMiddleware, comentController.Update)
	app.Delete("/comments/:id", middleware.AuthMiddleware, middleware.AdminMiddleware, comentController.Delete)
	app.Get("/comments/article/:articleId", middleware.AuthMiddleware, comentController.GetByArticleID)
	app.Get("/comments/user/:userId", middleware.AuthMiddleware, comentController.GetByUserID)

}
