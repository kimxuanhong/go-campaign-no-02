package service

//go:generate mockgen -source=sum_service.go -destination=mocks/sum_service_mock.go -package=mocks

type SumService interface {
	Valid(age int) bool
}

type SumServiceImpl struct {
}

func (r *SumServiceImpl) Valid(age int) bool {
	return age > 18
}
