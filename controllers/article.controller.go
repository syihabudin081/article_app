package controllers

import (
	"OrdentTest/models"
	"OrdentTest/services"
	"OrdentTest/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ArticleController struct {
	Service   services.ArticleService
	Validator *validator.Validate
}

func NewArticleController(service services.ArticleService) *ArticleController {
	return &ArticleController{
		Service:   service,
		Validator: validator.New(),
	}
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body models.Article true "Article object that needs to be created"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /articles [post]
func (controller *ArticleController) CreateArticle(c *fiber.Ctx) error {
	// Create a pointer to the Article model directly
	articleModel := new(models.Article)

	// Parse the request body into the Article model
	if err := c.BodyParser(articleModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	// Validate the Article model (optional)
	if err := controller.Validator.Struct(articleModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   utils.ParseValidationErrors(err),
		})
	}

	// Get authorID from context and convert to uuid.UUID
	authorIDString, ok := c.Locals("authorID").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Author ID not found in context",
		})
	}

	authorID, err := uuid.Parse(authorIDString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid author ID",
			Error:   err.Error(),
		})
	}

	// Set the AuthorID in the Article model
	articleModel.AuthorID = authorID

	// Call the service to create the article
	if err = controller.Service.CreateArticle(articleModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to create article",
			Error:   err.Error(),
		})
	}

	// Return a success response with the article response DTO
	return c.Status(fiber.StatusCreated).JSON(&utils.Response{
		Status:  "success",
		Message: "Article created successfully",
		Data:    articleModel, // Return the created article model
	})
}

// GetArticle godoc
// @Summary Get an article by ID
// @Description Get an article by ID
// @Tags articles
// @Accept  json
// @Produce  json
// @Param id path string true "Article ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /articles/{id} [get]
func (controller *ArticleController) GetArticleByID(c *fiber.Ctx) error {
	id := c.Params("id")
	article, err := controller.Service.GetArticleById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to get article",
			Error:   err.Error(),
		})
	}

	return c.JSON(&utils.Response{
		Status: "success",
		Data:   article,
	})
}

// GetArticles godoc
// @Summary Get all articles
// @Description Get all articles
// @Tags articles
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /articles [get]
func (controller *ArticleController) GetArticles(c *fiber.Ctx) error {
	articles, err := controller.Service.GetAllArticles()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to get articles",
			Error:   err.Error(),
		})
	}

	return c.JSON(&utils.Response{
		Status: "success",
		Data:   articles,
	})
}

// UpdateArticle godoc
// @Summary Update an article by ID
// @Description Update an article by ID
// @Tags articles
// @Accept  json
// @Produce  json
// @Param id path string true "Article ID"
// @Param article body models.Article true "Article object that needs to be updated"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /articles/{id} [put]
func (controller *ArticleController) UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	articleModel := new(models.Article)

	if err := c.BodyParser(articleModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	// Validate the Article model (optional)
	if err := controller.Validator.Struct(articleModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   utils.ParseValidationErrors(err),
		})
	}

	// Map authorID from context and convert to uuid.UUID
	authorIDString := c.Locals("authorID").(string)
	authorID, err := uuid.Parse(authorIDString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid author ID",
			Error:   err.Error(),
		})
	}

	// Set the AuthorID in the Article model
	articleModel.AuthorID = authorID

	err = controller.Service.UpdateArticle(id, articleModel)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to update article",
			Error:   err.Error(),
		})
	}

	return c.JSON(&utils.Response{
		Status:  "success",
		Message: "Article updated successfully",
		Data:    articleModel, // Return the updated article model
	})
}

// DeleteArticle godoc
// @Summary Delete an article by ID
// @Description Delete an article by ID
// @Tags articles
// @Accept  json
// @Produce  json
// @Param id path string true "Article ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /articles/{id} [delete]
func (controller *ArticleController) DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	err := controller.Service.DeleteArticle(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to delete article",
			Error:   err.Error(),
		})
	}

	return c.JSON(&utils.Response{
		Status:  "success",
		Message: "Article deleted successfully",
	})
}
