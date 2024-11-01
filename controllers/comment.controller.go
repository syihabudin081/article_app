package controllers

import (
	"OrdentTest/models"
	"OrdentTest/services"
	"OrdentTest/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CommentController struct {
	Service   services.CommentService
	Validator *validator.Validate
}

func NewCommentController(service services.CommentService) *CommentController {
	return &CommentController{
		Service:   service,
		Validator: validator.New(),
	}
}

// Create godoc
// @Summary Create a new comment
// @Description Create a new comment for an article
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment object that needs to be created"
// @Success 201 {object} models.Comment
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /comments [post]
func (cc *CommentController) Create(c *fiber.Ctx) error {
	var comment models.Comment

	// Bind the request body to the Comment struct, excluding UserID
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
	}

	// Automatically set UserID from the context/session
	userIDStr, ok := c.Locals("userId").(string)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "User not authenticated",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user ID",
		})
	}
	comment.UserID = userID

	// Ensure the ArticleID is valid
	if comment.ArticleID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Article ID must be a valid UUID",
		})
	}

	// Create the comment using the service
	if err := cc.Service.CreateComment(&comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}

// Get comments by article ID godoc
// @Summary Get comments for an article
// @Description Get all comments for a specific article
// @Tags comments
// @Accept json
// @Produce json
// @Param article_id path string true "Article ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /article/{articleId}/comments [get]
func (cc *CommentController) GetByArticleID(c *fiber.Ctx) error {

	// uuid to string
	articleIDString := c.Params("articleId")
	comments, err := cc.Service.GetCommentsByArticleId(articleIDString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&utils.Response{
			Status:  "error",
			Message: "No comments found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "Comments retrieved successfully",
		Data:    comments,
	})
}

// Get comment by ID godoc
// @Summary Get a comment by ID
// @Description Get a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /comments/{id} [get]
func (cc *CommentController) GetByID(c *fiber.Ctx) error {

	commentIDString := c.Params("id")
	comment, err := cc.Service.GetCommentById(commentIDString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&utils.Response{
			Status:  "error",
			Message: "Comment not found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status: "success",
		Data:   comment,
	})
}

// Update godoc
// @Summary Update a comment by ID
// @Description Update a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param comment body models.Comment true "Comment object that needs to be updated"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /comments/{id} [put]
func (cc *CommentController) Update(c *fiber.Ctx) error {

	commentIDString := c.Params("id")

	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	// Validate the comment
	if err := cc.Validator.Struct(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	if err := cc.Service.UpdateComment(commentIDString, comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to update comment",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "Comment updated successfully",
		Data:    comment,
	})
}

// Delete godoc
// @Summary Delete a comment by ID
// @Description Delete a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /comments/{id} [delete]
func (cc *CommentController) Delete(c *fiber.Ctx) error {
	commentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid comment ID",
			Error:   err.Error(),
		})
	}
	// uuid to string
	commentIDString := commentID.String()

	if err := cc.Service.DeleteComment(commentIDString); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to delete comment",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "Comment deleted successfully",
	})
}

// Get comments by user ID godoc
// @Summary Get comments by user ID
// @Description Get all comments for a specific user
// @Tags comments
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Router /user/{userId}/comments [get]
func (cc *CommentController) GetByUserID(c *fiber.Ctx) error {

	userIDString := c.Params("userId")
	comments, err := cc.Service.GetCommentsByUserId(userIDString)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&utils.Response{
			Status:  "error",
			Message: "No comments found",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "Comments retrieved successfully",
		Data:    comments,
	})
}
