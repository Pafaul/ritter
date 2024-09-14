package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pafaul/ritter/models"
	"pafaul/ritter/services"
)

// CreateUserHandler creates a new user.
// @Summary Create a new user
// @Description Create a new user with the given information.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "User info"
// @Success 201 {object} models.CreateUser
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/register [post]
func (h *Handler) CreateUserHandler(c echo.Context) error {
	user := new(models.CreateUser)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := services.CreateUser(c.Request().Context(), h.q, user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
