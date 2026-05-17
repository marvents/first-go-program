package users

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"program/db"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) TrimFieldSpace() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)
}

type LoggedUserStruct struct {
	ID   int
	Name string
}

var Users []User
var LoggedUser LoggedUserStruct

func LoadData() error {
	u, err := db.LoadData("db/json/users.json")
	if err != nil {
		return err
	}

	var data []User
	err3 := json.Unmarshal(u, &data)
	if err3 != nil {
		return err3
	}

	Users = data

	return nil
}

func Register(u User) error {
	u.TrimFieldSpace()
	if len(u.Name) < 4 || len(u.Password) < 6 {
		return errors.New("name or password is shorter")
	}

	isEmail := strings.Contains(u.Email, "@")
	if !isEmail {
		return errors.New("pls enter true email")
	}

	b, err0 := bcrypt.GenerateFromPassword(
		[]byte(u.Password),
		bcrypt.DefaultCost,
	)
	if err0 != nil {
		return err0
	}
	u.Password = string(b)

	u.ID = len(Users) + 1
	Users = append(Users, u)
	data, err := json.MarshalIndent(Users, "", "  ")
	if err != nil {
		return err
	}
	err2 := db.SaveData("db/json/users.json", data, 0666)

	return err2
}

func Login(email string, password string) (string, bool) {
	email = strings.TrimSpace(email)
	isEmail := strings.Contains(email, "@")
	if !isEmail {
		return "pls enter true email", false
	}

	for _, u := range Users {
		if u.Email == email && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil {
			LoggedUser = LoggedUserStruct{
				ID:   u.ID,
				Name: u.Name,
			}

			return "User Logged in successfully", true
		}
	}
	return "Email or password doesn't exist", false
}
