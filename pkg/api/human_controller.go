package api

import (
	"fmt"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/auth"
	"net/http"
	"strconv"

	"github.com/kimxuanhong/go-campaign-no-02/pkg/dto"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/service"
	"github.com/labstack/echo/v4"
)

type HumanController interface {
	GetPersons(c echo.Context) error
	GetPerson(c echo.Context) error
	CreatePerson(c echo.Context) error
	UpdatePerson(c echo.Context) error
	DeletePerson(c echo.Context) error
}

type HumanControllerImpl struct {
	personService service.PersonService
}

func NewHumanController() *HumanControllerImpl {
	return &HumanControllerImpl{
		personService: service.NewPersonService(),
	}
}

func HumanControllerRouter(e *echo.Group) {
	controller := NewHumanController()

	e.GET("/GetPersons", controller.GetPersons, auth.HasRole("customer"))
	e.GET("/GetPerson/:id", controller.GetPerson, auth.HasRole("customer"))
	e.POST("/CreatePerson", controller.CreatePerson, auth.HasRole("admin"))
	e.PUT("/UpdatePerson/:id", controller.UpdatePerson, auth.HasRole("admin"))
	e.DELETE("/DeletePerson/:id", controller.DeletePerson, auth.HasRole("customer", "admin"))
}

func (ctr *HumanControllerImpl) GetPersons(c echo.Context) error {
	persons := ctr.personService.GetPersons()
	fmt.Printf("FullName = %v\n", persons)

	return c.JSON(http.StatusOK, persons)
}

func (ctr *HumanControllerImpl) GetPerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	person := ctr.personService.GetPerson(id)
	fmt.Printf("FullName = %v\n", person)

	return c.JSON(http.StatusOK, person)
}

func (ctr *HumanControllerImpl) CreatePerson(c echo.Context) error {
	personReq := dto.PersonRequest{}
	if err := c.Bind(&personReq); err != nil {
		return err
	}

	person := ctr.personService.CreatePerson(personReq)
	fmt.Printf("FullName = %v\n", person)

	return c.JSON(http.StatusOK, person)
}

func (ctr *HumanControllerImpl) UpdatePerson(c echo.Context) error {
	person := ctr.personService.GenPerson(1)
	fmt.Printf("FullName = %v\n", person)

	return c.JSON(http.StatusOK, person)
}

func (ctr *HumanControllerImpl) DeletePerson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctr.personService.DeletePerson(id)

	return c.JSON(http.StatusOK, "ok")
}
