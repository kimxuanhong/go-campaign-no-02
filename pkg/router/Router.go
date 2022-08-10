package router

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router interface {
	Start()
}

type RouterImpl struct {
}

func (r *RouterImpl) Start() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api.HumanControllerRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
