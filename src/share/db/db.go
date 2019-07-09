package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

//数据库实例
var (
	db *sqlx.DB
)

// init实例化
func Init(mysqlDSN string)  {
	db = sqlx.MustConnect("mysql", mysqlDSN)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(3)
}
