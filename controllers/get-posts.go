package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pafaul/ritter/models"
	"pafaul/ritter/services"
)

// GetPosts godoc
// @Summary Get posts
// @Description Get posts from the system
// @Tags Post
// @Accept json
// @Produce json
// @Param getPosts body models.GetPosts false "describe your GetPosts object"
// @Success 201 {array} models.Post
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /post/posts [post]
func (h *Handler) GetPosts(c echo.Context) error {
	getPosts := new(models.GetPosts)
	if err := c.Bind(getPosts); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(getPosts); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	posts, err := services.GetPosts(c.Request().Context(), h.q, getPosts)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}
