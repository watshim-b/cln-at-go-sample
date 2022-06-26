package usecase

import (
	"github.com/watshim-b/cln-at-go/interface/web/form"
)

// 一般職向けのLogin機能を実装するUsecaseです
// どうしても共通的な実装としたい場合は、共通LoginUsecaseを定義し、移譲を使って機能提供してください。
type LowlyEmployeUsecase struct{}

func NewLowlyEmployeUsecase() *LowlyEmployeUsecase {
	return &LowlyEmployeUsecase{}
}

// 一般職のログインに関するusecase
func (u *LowlyEmployeUsecase) Login(f form.LginForm) error {
	return nil
}
