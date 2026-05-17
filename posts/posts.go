package posts

import (
	"encoding/json"
	"errors"
	"program/db"
	"program/users"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"userid"`
}

var Posts []Post

func LoadData() error {
	p, err2 := db.LoadData("db/json/posts.json")
	if err2 != nil {
		return err2
	}

	var data2 []Post
	err4 := json.Unmarshal(p, &data2)
	if err4 != nil {
		return err4
	}

	Posts = data2
	return nil
}

func CreatePost(t, c string, u users.LoggedUserStruct) error {
	if len(t) < 5 || len(c) < 10 {
		return errors.New("title or content is so shorter, pls make title.min > 4 & content.min > 10")
	}

	post := Post{
		ID:      len(Posts) + 1,
		Title:   t,
		Content: c,
		UserId:  u.ID,
	}

	Posts = append(Posts, post)

	data, err := json.MarshalIndent(Posts, "", "  ")
	if err != nil {
		return err
	}
	err2 := db.SaveData("db/json/posts.json", data, 0666)
	if err2 != nil {
		return err2
	}

	err3 := LoadData()
	if err3 != nil {
		return err3
	}

	return nil
}
