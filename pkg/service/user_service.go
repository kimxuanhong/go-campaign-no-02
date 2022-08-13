package service

import (
	"github.com/kimxuanhong/go-campaign-no-02/pkg/models"
	"github.com/kimxuanhong/go-campaign-no-02/pkg/slice"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindUserByEmail(email string) *models.PersonImpl
}

type UserServiceImpl struct {
	arr []models.PersonImpl
}

var instanceUserService *UserServiceImpl

func NewUserService() *UserServiceImpl {
	if instanceUserService == nil {
		var arr []models.PersonImpl
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), 10)
		arr = append(arr, *models.NewUser(1, "Hung Pham", 22, "dark nong", "hungpham", string(hashedPassword), "admin", "customer"))

		instanceUserService = &UserServiceImpl{
			arr: arr,
		}
	}
	return instanceUserService
}

func (r *UserServiceImpl) FindUserByEmail(username string) *models.PersonImpl {
	user := slice.FirstOrDefault(r.arr, func(p *models.PersonImpl) bool {
		return p.Username == username
	})

	return user
}
