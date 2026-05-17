package main

import (
	"fmt"
	"program/users"
	"program/posts"
)

func main() {
	user := users.User{
		ID:0,
		Name: "Marouane",
		Email: "marouan@proton.me",
		Password: "Marouane123",
	}
	users.Register(user)
	users.Login("marouan@proton.me", "Marouane123")
	fmt.Println(users.LoggedUser)

	posts.CreatePost("Go Lang", "Go lang is easy Lang to learn, just start", users.LoggedUser)
	fmt.Println(posts.Posts)
}
