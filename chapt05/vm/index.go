package vm

import "go-practise/chapt05/model"

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	user, _ := model.GetUserByUsername("wanweibin")
	posts, _ := model.GetPostsByUserID(user.ID)

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *user, *posts}
	return v
}
