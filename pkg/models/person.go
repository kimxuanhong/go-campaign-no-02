package models

import "fmt"

type Person interface {
	ShowInfo()
	UpdateInfo(name string, age int, address string)
	GetName() string
}

type PersonImpl struct {
	Id       int
	Name     string
	Username string
	Password string
	Age      int
	Address  string
	Roles    []string
}

func (person *PersonImpl) ShowInfo() {
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Address: %s\n", person.Address)
}

func (person *PersonImpl) UpdateInfo(id int, name string, age int, address string) {
	person.Id = id
	person.Name = name
	person.Age = age
	person.Address = address
}

func (person *PersonImpl) GetName() string {
	return person.Name
}

func NewPerson(id int, name string, age int, address string) *PersonImpl {
	return &PersonImpl{
		Id:      id,
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func NewUser(id int, name string, age int, address string, username string, password string, roles ...string) *PersonImpl {
	return &PersonImpl{
		Id:       id,
		Name:     name,
		Age:      age,
		Address:  address,
		Username: username,
		Password: password,
		Roles:    roles,
	}
}
