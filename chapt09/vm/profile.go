package vm

import "go-practise/chapt09/model"

type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post
	Editable       bool
	ProfileUser    model.User
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	BasePageViewModel
}

type ProfileViewModelOp struct{}

func (ProfileViewModelOp) GetVM(sUser, pUser string, page, limit int) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}

	posts, total, _ := model.GetPostsByUserIdPageAndLimit(u1.ID, page, limit)
	v.ProfileUser = *u1
	v.Editable = (pUser == sUser)
	v.SetBasePageViewModel(total, page, limit)

	if !v.Editable {
		v.IsFollow = u1.IsFollowedByUser(sUser)
	}
	v.FollowersCount = u1.FollowersCount()
	v.FollowingCount = u1.FollowingCount()

	v.Posts = *posts
	v.SetCurrentUser(sUser)

	return v, nil
}

func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}

	return u.Follow(b)
}

func Unfollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Unfollow(b)
}
