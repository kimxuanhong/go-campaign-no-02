package handler

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/api"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router interface {
	Start()
}

type RouterImpl struct {
	jwtConfig auth.JwtConfig
	router    *echo.Echo
	address   string
}

func NewRouter(address string) *RouterImpl {
	return &RouterImpl{
		jwtConfig: &auth.JwtConfigImpl{},
		router:    echo.New(),
		address:   address,
	}
}

func (r *RouterImpl) Start() {

	// Middleware
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())

	uncheckGroup := r.router.Group("")

	adminGroup := r.router.Group("/admin")
	adminGroup.Use(middleware.JWTWithConfig(r.jwtConfig.GetConfig()))

	// Routes
	api.AuthControllerRouter(uncheckGroup)
	api.HumanControllerRouter(adminGroup)

	// Start server
	r.router.Logger.Fatal(r.router.Start(r.address))
}
