package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/slice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type HasRoleConfig struct {
	JWTContextKey string
}

var DefaultHasRoleConfig = HasRoleConfig{
	JWTContextKey: middleware.DefaultJWTConfig.ContextKey,
}

func HasRole(roles ...string) echo.MiddlewareFunc {
	return HasRoleWithConfig(DefaultHasRoleConfig, roles...)
}

func HasRoleWithConfig(config HasRoleConfig, roles ...string) echo.MiddlewareFunc {

	var errForbidden = func(role string) error {
		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("role(s) [%s] required to access the resource", role))
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user := c.Get(config.JWTContextKey)
			if user == nil {
				return errForbidden(strings.Join(roles, ", "))
			}

			userRoles := user.(*jwt.Token).Claims.(*CustomClaims).User.Roles
			if userRoles == nil {
				return errForbidden(strings.Join(roles, ", "))
			}

			for _, r := range roles {
				if slice.Contains(userRoles, func(t *string) bool { return *t == r }) {
					return next(c)
				}
			}

			return errForbidden(strings.Join(roles, ", "))
		}
	}
}
