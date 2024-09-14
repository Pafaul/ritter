package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pafaul/ritter/models"
	"pafaul/ritter/services"
)

// CreatePostHandler Creates a new Post.
// @Summary Creates a new Post
// @Tags Post
// @Accept json
// @Produce json
// @Param post body models.Post true "Post information"
// @Success 201 {object} models.Post
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /post/create [post]
func (h *Handler) CreatePostHandler(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := services.CreatePost(c.Request().Context(), h.q, post); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, post)
}
