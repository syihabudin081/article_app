package controllers

import (
	"OrdentTest/models"
	"OrdentTest/services"
	"OrdentTest/utils"
	"OrdentTest/utils/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Service   services.UserService
	Validator *validator.Validate
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		Service:   service,
		Validator: validator.New(),
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object that needs to be registered"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /users [post]
func (uc *UserController) Register(c *fiber.Ctx) error {
	userModel := new(models.User)
	if err := c.BodyParser(userModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	if err := uc.Validator.Struct(userModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to hash password",
			Error:   err.Error(),
		})
	}
	userModel.Password = string(hashedPassword)

	// If role is not provided, set it to user
	if userModel.RoleID == 0 {
		userModel.RoleID = 1 // Default role ID
	}

	if err := uc.Service.CreateUser(userModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to register user",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&utils.Response{
		Status:  "success",
		Message: "User registered successfully",
		Data:    userModel,
	})
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object that needs to be logged in"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /users/login [post]
func (uc *UserController) Login(c *fiber.Ctx) error {
	userModel := new(models.User)
	if err := c.BodyParser(userModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	// Get stored user by email
	storedUser, err := uc.Service.GetUserByEmail(userModel.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid credentials",
			Error:   err.Error(),
		})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(userModel.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid credentials",
			Error:   err.Error(),
		})
	}

	// Generate token
	token, err := jwt.GenerateToken(storedUser.Email, storedUser.Role.Name, storedUser.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to generate token",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "Login successful",
		Data:    token,
	})
}
