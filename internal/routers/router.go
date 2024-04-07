package routers

import (
	"github.com/SimonVillalonIT/filehub/internal/handlers"
	"github.com/SimonVillalonIT/filehub/internal/templates"
	"github.com/SimonVillalonIT/filehub/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo, config *config.Config) {
	component := templates.Layout(templates.Login(), "login")
	loginHandler := handlers.LoginHandler{Config: config}
	filesHandler := handlers.FilesHandler{Config: config}

	// authMiddleware := middlewares.AuthMiddleware{Config: config}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error { return templates.Render(c, 200, component) })
	e.POST("/login", loginHandler.Login)

	g := e.Group("/files")
	// g.Use(authMiddleware.Init)
	g.GET("/", filesHandler.GetFiles)
	g.POST("/", filesHandler.PostFiles)
}
