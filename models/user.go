package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	EMAIL        = ""
	IS_SUPERUSER = 0
	FIRST_NAME   = ""
	LAST_NAME    = ""
	IS_STAFF     = 0
	IS_ACTIVE    = 1
)

func AddUser(username, password string) (int64, error) {
	o := orm.NewOrm()
	// passsword
	sql := "insert into auth_user(username, password, email, date_joined, is_superuser, first_name, last_name, is_staff, is_active) values(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	cur_time := time.Now()
	res, err := o.Raw(sql, username, password, EMAIL, cur_time, IS_SUPERUSER, FIRST_NAME, LAST_NAME, IS_STAFF, IS_ACTIVE).Exec()
	if err != nil {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func CheckUserName(username string) bool {
	var maps []orm.Params
	o := orm.NewOrm()
	sql := "select username from auth_user where username = ?"
	num, err := o.Raw(sql, username).Values(&maps)
	if err == nil && num == 0 {
		return true
	} else {
		return false
	}
}

func UpdateUser(username string)  {

}

