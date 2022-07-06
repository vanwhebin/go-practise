package vm

import "go-practise/chapt07/model"

type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	ProfileUser model.User
}

type ProfileViewModelOp struct{}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)

}
