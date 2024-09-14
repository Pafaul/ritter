package client_server

import (
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"pafaul/ritter/controllers"
	_ "pafaul/ritter/docs"
	"pafaul/ritter/ritter_db"
)

func NewServer(addr string, db *sql.DB) *http.Server {
	s := &http.Server{
		Addr:    addr,
		Handler: newServeMux(db),
	}

	return s
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func newServeMux(db *sql.DB) *echo.Echo {
	h := controllers.NewHandler(ritter_db.New(db))

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	pages := e.Group("")
	pages.File("/", "public/index.html")
	pages.File("/index.html", "public/index.html")

	apiV1Path := e.Group("/api/v1")

	userGroup := apiV1Path.Group("/user")
	userGroup.POST("/register", h.CreateUserHandler)

	postGroup := apiV1Path.Group("/post")
	postGroup.POST("/create", h.CreatePostHandler)
	postGroup.POST("/posts", h.GetPosts)
	return e
}

var _ http.Handler = (*echo.Echo)(nil)
