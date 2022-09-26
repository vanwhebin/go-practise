package vm

import "go-practise/chapt11/model"

type ExploreViewModel struct {
	BaseViewModel
	Posts []model.Post
	BasePageViewModel
}

type ExploreViewModelOp struct{}

func (ExploreViewModelOp) GetVM(username string, page, limit int) ExploreViewModel {
	posts, total, _ := model.GetPostsByPageAndLimit(page, limit)
	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.Limit = limit
	v.CurrentPage = page
	v.SetBasePageViewModel(total, page, limit)
	v.SetPrevAndNextPage()
	v.SetCurrentUser(username)
	return v
}
