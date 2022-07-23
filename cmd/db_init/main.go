package main

import (
	"log"

	"go-practise/chapt09/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("wwb", "abc123", "wwb@example.com")
	model.AddUser("wanweibin", "abc123", "wanweibin@example.com")

	u1, _ := model.GetUserByUsername("wwb")
	u1.CreatePost("Beautiful day in Portland!")
	model.UpdateAboutMe(u1.Username, `I'm the author of Go-Mega Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("wanweibin")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u1.Follow(u2.Username)
}
