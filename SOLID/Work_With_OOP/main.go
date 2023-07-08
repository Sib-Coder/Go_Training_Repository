package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street string
	City   string
}
type User struct {
	FirstName string
	LastName  string
	Email     string
	DataBirth time.Time
	Address   Address
}

func (u *User) ShowCoolNameAndCity() {
	fmt.Printf("%s-%s %s\n", u.FirstName, u.LastName, u.Address.City)
}

func (u *User) SetFirstName(newname string) {
	u.FirstName = newname

}

func NewUser(firstname, lastname, email string, databirth time.Time, address Address) *User { // конструктор по умолчанию
	return &User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		DataBirth: databirth,
		Address:   address,
	}
}

func UpdateUser(u *User) {
	u.LastName = "some name"
}

type Admin struct {
	IsAdmin bool
	User
}

func (u *Admin) ShowCoolNameAndCity() {
	fmt.Printf("%s-%s %s Admin: %t\n", u.FirstName, u.LastName, u.Address.City, u.IsAdmin)
}

func main() {
	address := Address{Street: "Bacunina", City: "Tomsk"}

	user := NewUser(
		"Daniil",
		"Sinitsyn",
		"dansin@mail.ru",
		time.Now(),
		address,
	)
	//UpdateUser(user)
	//fmt.Println(user)
	user.ShowCoolNameAndCity()
	user.SetFirstName("Vlad")
	user.ShowCoolNameAndCity()

	admin := Admin{
		IsAdmin: true,
		User:    *user, //передали все параметры из юзера в админ
	}
	admin.ShowCoolNameAndCity()
}
