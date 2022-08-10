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

func NewUserService() *UserServiceImpl {
	var arr []models.PersonImpl
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345"), 10)
	arr = append(arr, *models.NewUser(1, "Hung Pham", 22, "dark nong", "hungpham", string(hashedPassword), "admin", "customer"))

	return &UserServiceImpl{
		arr: arr,
	}
}

func (r *UserServiceImpl) FindUserByEmail(username string) *models.PersonImpl {
	user := slice.FirstOrDefault(r.arr, func(p *models.PersonImpl) bool {
		return p.Username == username
	})

	return user
}
