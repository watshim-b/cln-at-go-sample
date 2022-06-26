package gateway

import (
	"github.com/watshim-b/cln-at-go/ent"

	_ "github.com/go-sql-driver/mysql"
)

var cli *ent.Client

// コネクションを開く
func OpenConnection() (*ent.Client, error) {
	c, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	if err != nil {
		return nil, err
	}
	cli = c
	return cli, nil
}

// コネクションを閉じる
func CloseConnection() {
	if cli != nil {
		cli.Close()
	}
}
