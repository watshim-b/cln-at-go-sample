package entity

import (
	"github.com/watshim-b/cln-at-go/ent"
	"github.com/watshim-b/cln-at-go/interface/web/form"
)

type UserEntity interface {
	// formで指定された情報で認証できるかを確認する
	Auth(f form.LoginForm) (bool, error)

	// formで指定された条件に合致するUserEntityを返却する
	Query(f form.LoginForm) ([]ent.User, error)

	// formで指定されたデータを作成する
	Create(f form.LoginForm) error

	// formで指定されたデータを作成する（バルクインサート）
	Creates(f []form.LoginForm) error

	// formで指定されたデータを更新する（idをキー値として更新をかける）
	Update(f form.LoginForm) error
}
