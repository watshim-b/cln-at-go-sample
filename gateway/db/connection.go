package gateway

import (
	"github.com/watshim-b/cln-at-go-sample/ent"

	_ "github.com/go-sql-driver/mysql"
)

// コネクションを開く
func OpenConnection() (*ent.Client, error) {
	return ent.Open("mysql", "root:rootpass@(my-cont:3306)/test?parseTime=true")
}

// docker コンテナ同士を接続させる際に発生する, NYSQL ERRコード 111の対策
// https://nainaistar.hatenablog.com/entry/2021/06/14/120000
// localhostではなく、コンテナ名をホスト名に指定する必要がある。

// parseTime=trueに関して
// https://budougumi0617.github.io/2019/03/31/go_db_unsupported_scan_storing_driver_value_type_uint8_into_type_time_time/

//ログ出力に関して
// https://pod.hatenablog.com/entry/2020/09/30/073034
