package handler

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Validator interface {
	Validate(i interface{}) error
}

type ValidatorImpl struct {
	Validator *validator.Validate
}

func (cv *ValidatorImpl) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
