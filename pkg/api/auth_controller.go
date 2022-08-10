package api

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/auth"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/dto"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/service"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthController interface {
	Login(c echo.Context) error
	RefreshToken(c echo.Context) error
}

type AuthControllerImpl struct {
	jwtConfig   auth.JwtConfig
	userService service.UserService
}

func NewAuthController() *AuthControllerImpl {
	return &AuthControllerImpl{
		jwtConfig:   &auth.JwtConfigImpl{},
		userService: service.NewUserService(),
	}
}

func AuthControllerRouter(e *echo.Group) {
	controller := NewAuthController()
	e.POST("/Login", controller.Login)
	e.POST("/RefreshToken", controller.RefreshToken)
}

func (r *AuthControllerImpl) Login(c echo.Context) error {
	loginReq := dto.UserLogin{}
	if err := c.Bind(&loginReq); err != nil {
		return err
	}

	user := r.userService.FindUserByEmail(loginReq.Username)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Username is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		// If the two passwords don't match, return a 401 status.
		return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
	}

	userLogin := &dto.User{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Roles:    user.Roles,
	}

	token, _ := r.jwtConfig.GenerateAccessToken(userLogin)

	return c.JSON(http.StatusOK, token)
}

func (r *AuthControllerImpl) RefreshToken(c echo.Context) error {
	refreshTokenReq := dto.Token{}
	if err := c.Bind(&refreshTokenReq); err != nil {
		return err
	}
	user, err := r.jwtConfig.ValidateRefreshToken(refreshTokenReq)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "invalid token",
		})
	}

	token, err := r.jwtConfig.GenerateAccessToken(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "unable to create access token",
		})
	}

	return c.JSON(http.StatusOK, token)
}
