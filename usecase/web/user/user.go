package usecase

import (
	form "github.com/watshim-b/cln-at-go-sample/form/web"
)

type UserUsecase interface {

	// User情報を取得します
	GetUser(f form.UserForm) error
}
