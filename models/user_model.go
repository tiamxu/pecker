package models

import (
	"fmt"
	"log"

	"github.com/tiamxu/pecker/database"
)

type User struct {
	Id       int    `form:"id" json:"id" db:"id" ini:"id"`
	UserName string `form:"username" json:"username" db:"username" ini:"username"`
	Password string `form:"password" json:"password" db:"password" ini:"password"`
}

// 插入
func InsertUser(user User) (int64, error) {
	sqlStr := `INSERT INTO user (id,username, password) VALUES (?,?,?);`
	return database.ModifyDB(sqlStr, user.Id, user.UserName, user.Password)
}

// 按条件查询
//
//	func QueryUserWightCon(con string) int {
//		sql := fmt.Sprintf("select id from user %s", con)
//		fmt.Println(sql)
//		row := database.QueryRowDB(sql)
//		fmt.Println("------")
//		id := 0
//		row.Scan(&id)
//		return id
//	}
func QueryUserWightCon(sql string) ([]User, error) {
	sql = "select id,username,password from user " + sql
	fmt.Println(sql)
	rows, err := database.QueryDB(sql)
	fmt.Println(&rows)
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}

// 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) ([]User, error) {
	whereSql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(whereSql)
}
