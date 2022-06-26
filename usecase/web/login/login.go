package usecase

import (
	"github.com/watshim-b/cln-at-go/interface/web/form"
)

type LoginUsecase interface {

	// ログイン機能を提供します
	Login(f form.LginForm) error

}
