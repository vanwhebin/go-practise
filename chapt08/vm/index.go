package vm

import "go-practise/chapt08/model"

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM(username string) IndexViewModel {
	user, _ := model.GetUserByUsername(username)
	posts, _ := model.GetPostsByUserID(user.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *user, *posts}
	v.SetCurrentUser(username)
	return v
}
