package usecase

import (
	form "github.com/watshim-b/cln-at-go-sample/form/web"
	gateway "github.com/watshim-b/cln-at-go-sample/gateway/db"
	"github.com/watshim-b/cln-at-go-sample/presenter"
	"github.com/watshim-b/cln-at-go-sample/response"
)

// 一般職向けのLogin機能を実装するUsecaseです
// どうしても共通的な実装としたい場合は、共通LoginUsecaseを定義し、移譲を使って機能提供してください。
type userUsecaseForLowlyEmploye struct{}

func NewUserUsecaseForLowlyEmploye() *userUsecaseForLowlyEmploye {
	return &userUsecaseForLowlyEmploye{}
}

// 一般職のログインに関するusecase
func (u *userUsecaseForLowlyEmploye) GetUser(f form.UserForm) (*response.UserResponse, error) {

	// コネクションの取得
	c, err := gateway.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// user情報の取得
	ug := gateway.NewUserGateWay(c)
	users, err := ug.Query(f)
	if err != nil {
		return nil, err
	}

	// レスポンスのデータを整える
	up := presenter.NewUserPresenter()
	response := up.GenerateResponse(users)

	return &response, nil
}
