package users

import (
	"strings"
)

type User struct {
	ID int
	Name string
	Email string
	Password string
}

type LoggedUserStruct struct {
	ID int
	Name string
}

var Users []User
var LoggedUser LoggedUserStruct

func Register(u User) (string ,bool) {
	if len(u.Name) < 4 || len(u.Password) < 6 {
		return "name or password is shorter",  false
	}

	isEmail := strings.Contains(u.Email, "@")
	if !isEmail { return "pls enter true email", false }

	// u.password := crypt(u.password, salt) in real app or smh like this
	u.Email = strings.TrimSpace(u.Email)

	u.ID = len(Users) +1
	Users = append(Users, u)

	return "User Created successfully", true
}

func Login(email string, password string) (string, bool) {
	isEmail := strings.Contains(email, "@")
	if !isEmail { return "pls enter true email", false }

	for _, u := range Users {
		if u.Email == strings.TrimSpace(email) && u.Password == password {
			LoggedUser = LoggedUserStruct{
				ID: u.ID,
				Name: u.Name,
			}

			return "User Logged in successfully", true
		}
	}
	return "Email or password doesn't exist", false
}