package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 查询单行
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// 查询多行
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
