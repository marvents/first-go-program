package main

import (
	"fmt"
	"log"
	"program/posts"
	"program/users"
	// "errors"
	// "os"
)

func LoadData() error {
	err := users.LoadData();
	if err != nil { return err }
	err2 := posts.LoadData();
	if err2 != nil { return err2 }
	return nil
}

func main() {

	errx := LoadData()
	if errx != nil { log.Fatal(errx)}

	user := users.User{
		ID:       0,
		Name:     "Marouane",
		Email:    "marouan@proton.me",
		Password: "Marouane123",
	}
	err := users.Register(user) // there are no check if user already exist so i will add to when I learn http section cz I'll build a simple interface with react
	if err != nil { log.Fatal(err)}
	users.Login("marouan@proton.me", "Marouane123")
	fmt.Println(users.LoggedUser)

	posts.CreatePost("Go Lang", "Go lang is easy Lang to learn, just start", users.LoggedUser) // also here the same problem, the fix soon!
	fmt.Println(posts.Posts)
}
