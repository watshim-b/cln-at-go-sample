package gateway

import (
	"context"

	"github.com/watshim-b/cln-at-go-sample/ent"
	"github.com/watshim-b/cln-at-go-sample/ent/user"
	form "github.com/watshim-b/cln-at-go-sample/form/web"
)

type userGateWay struct {
	c *ent.Client
}

func NewUserGateWay(c *ent.Client) *userGateWay {
	return &userGateWay{c}
}

// formで指定されたUserが存在するかをチェックする
func (g *userGateWay) Auth(f form.UserForm) (bool, error) {
	return true, nil
}

// formで指定された条件に合致するUserEntityを返却する
func (g *userGateWay) Query(f form.UserForm) ([]*ent.User, error) {

	q := g.c.User.Query()

	// IDの検索条件を追加する
	if f.ID != 0 {
		q.Where(user.ID(f.ID))
	}

	// Nameの検索条件を追加する
	if f.Name != "" {
		q.Where(user.Name(f.Name))
	}

	// Activeの検索条件を追加する
	if f.Active != 2 {
		q.Where(user.Active(f.Active))
	}

	return q.All(context.TODO())
}

// formで指定されたデータを作成する
func (g *userGateWay) Create(f form.UserForm) error {
	return nil
}

// formで指定されたデータを作成する（バルクインサート）
func (g *userGateWay) Creates(f []form.UserForm) error {
	return nil
}

// formで指定されたデータを更新する（idをキー値として更新をかける）
func (g *userGateWay) Update(f form.UserForm) error {
	return nil
}
