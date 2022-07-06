package vm

import "go-practise/chapt04/model"

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := model.User{Username: "wanweibin"}
	u2 := model.User{Username: "wwb"}

	posts := []model.Post{
		{User: u1, Body: "Beautiful day in China"},
		{User: u2, Body: "Nice day in SZ"},
	}

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	return v
}
