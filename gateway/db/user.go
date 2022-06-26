package gateway

import (
	"github.com/watshim-b/cln-at-go/ent"
	"github.com/watshim-b/cln-at-go/interface/web/form"
)

type userGateWay struct{}

func NewUserGateWay() *userGateWay {
	return &userGateWay{}
}

// formで指定されたUserが存在するかをチェックする
func (g *userGateWay) Auth(f form.LoginForm) (bool, error) {
	return true, nil
}

// formで指定された条件に合致するUserEntityを返却する
func (g *userGateWay) Query(f form.LoginForm) ([]ent.User, error) {
	return nil, nil
}

// formで指定されたデータを作成する
func (g *userGateWay) Create(f form.LoginForm) error {
	return nil
}

// formで指定されたデータを作成する（バルクインサート）
func (g *userGateWay) Creates(f []form.LoginForm) error {
	return nil
}

// formで指定されたデータを更新する（idをキー値として更新をかける）
func (g *userGateWay) Update(f form.LoginForm) error {
	return nil
}
