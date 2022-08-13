package service

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/dto"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/models"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/slice"
)

type PersonService interface {
	GenPerson(age int) models.PersonImpl
	GetPerson(id int) models.PersonImpl
	GetPersons() []models.PersonImpl
	CreatePerson(personReq dto.PersonRequest) models.PersonImpl
	DeletePerson(id int)
}

type PersonServiceImpl struct {
	sumService SumService
	arr        []models.PersonImpl
	seq        int
}

var instancePersonService *PersonServiceImpl

func NewPersonService() *PersonServiceImpl {
	if instancePersonService == nil {
		var arr []models.PersonImpl
		arr = append(arr, *models.NewPerson(1, "hung", 22, "dark nong"))

		instancePersonService = &PersonServiceImpl{
			sumService: NewSumService(),
			arr:        arr,
			seq:        3,
		}
	}
	return instancePersonService
}

func (r *PersonServiceImpl) GenPerson(age int) models.PersonImpl {

	person := models.PersonImpl{}

	if r.sumService.Valid(age) {
		person.UpdateInfo(2, "Hong da du 18 tuoi", 10, "Tra Vinh")
	} else {
		person.UpdateInfo(2, "Hong nho tuoi", 10, "Tra Vinh")
	}

	r.arr = append(r.arr, person)

	return person
}

func (r *PersonServiceImpl) GetPerson(id int) models.PersonImpl {
	person := slice.FirstOrDefault(r.arr, func(p *models.PersonImpl) bool {
		return p.Id == id
	})

	return *person
}

func (r *PersonServiceImpl) GetPersons() []models.PersonImpl {

	return r.arr
}

func (r *PersonServiceImpl) CreatePerson(personReq dto.PersonRequest) models.PersonImpl {
	person := models.NewPerson(r.seq, personReq.Name, personReq.Age, personReq.Address)
	r.seq++
	r.arr = append(r.arr, *person)

	return *person
}

func (r *PersonServiceImpl) DeletePerson(id int) {
	r.arr = slice.Remove(r.arr, func(t *models.PersonImpl) bool {
		return t.Id == id
	})
}
