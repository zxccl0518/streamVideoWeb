package dbops

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		fmt.Printf(" 打开数据库失败了. ===========接下来就是 panic 异常.============ ")
		panic(err.Error())
	}
}
