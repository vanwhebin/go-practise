package main

import (
	"fmt"
	"log"

	"go-practise/chapt07/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "wwb",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "wwb@example.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("wwb@example.com")),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
			},
		},
		{
			Username:     "wanweibin",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "wwb@test.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("wwb@test.com")),
			Posts: []model.Post{
				{Body: "The Avengers movie was so cool!"},
				{Body: "Sun shine is beautiful"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
