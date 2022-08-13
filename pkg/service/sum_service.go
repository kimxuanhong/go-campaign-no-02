package service

//go:generate mockgen -source=sum_service.go -destination=mocks/sum_service_mock.go -package=mocks

type SumService interface {
	Valid(age int) bool
}

type SumServiceImpl struct {
}

var instanceSumService *SumServiceImpl

func NewSumService() *SumServiceImpl {
	if instanceSumService == nil {
		instanceSumService = &SumServiceImpl{}
	}
	return instanceSumService
}

func (r *SumServiceImpl) Valid(age int) bool {
	return age > 18
}
