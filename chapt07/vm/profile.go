package vm

import "go-practise/chapt07/model"

type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	Editable    bool
	ProfileUser model.User
}

type ProfileViewModelOp struct{}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}

	posts, _ := model.GetPostsByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Editable = (pUser == sUser)
	v.Posts = *posts
	v.SetCurrentUser(sUser)

	return v, nil
}
