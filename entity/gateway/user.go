package entity

import (
	"github.com/watshim-b/cln-at-go-sample/ent"
	form "github.com/watshim-b/cln-at-go-sample/form/web"
)

type UserEntity interface {
	// formで指定された情報で認証できるかを確認する
	Auth(f form.UserForm) (bool, error)

	// formで指定された条件に合致するUserEntityを返却する
	Query(f form.UserForm) ([]*ent.User, error)

	// formで指定されたデータを作成する
	Create(f form.UserForm) error

	// formで指定されたデータを作成する（バルクインサート）
	Creates(f []form.UserForm) error

	// formで指定されたデータを更新する（idをキー値として更新をかける）
	Update(f form.UserForm) error
}
