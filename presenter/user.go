package presenter

import (
	"github.com/watshim-b/cln-at-go-sample/ent"
	"github.com/watshim-b/cln-at-go-sample/response"
)

type userPresenter struct{}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

func (p *userPresenter) GenerateResponse(users []*ent.User) response.UserResponse {
	ur := response.UserResponse{
		Users:     users,
		ListCount: len(users),
	}

	return ur
}
