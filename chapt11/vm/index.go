package vm

import "go-practise/chapt11/model"

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string

	BasePageViewModel
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM(username, flash string, page, limit int) IndexViewModel {
	user, _ := model.GetUserByUsername(username)
	posts, total, _ := user.FollowingPostsByPageAndLimit(page, limit)
	v := IndexViewModel{}
	v.SetTitle("HomePage")
	v.Posts = *posts
	v.Flash = flash
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}

func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
