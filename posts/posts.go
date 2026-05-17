package posts

import "program/users"

type Post struct {
	ID int
	Title string
	Content string
	UserId int
}

var Posts []Post

func CreatePost(t, c string, u users.LoggedUserStruct) bool {
	if len(t) < 5 || len(c) < 10 {
		return false
	}

	post := Post{
		ID: len(Posts) +1,
		Title: t,
		Content: c,
		UserId: u.ID,
	}

	Posts = append(Posts, post)
	return true
}