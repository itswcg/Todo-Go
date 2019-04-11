package models

import (
	"Todo-Go/utils"
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
	pass := utils.MakePassword(password)
	res, err := o.Raw(sql, username, pass, EMAIL, cur_time, IS_SUPERUSER, FIRST_NAME, LAST_NAME, IS_STAFF, IS_ACTIVE).Exec()
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

func CheckPassword(username, password string) bool {
	var maps []orm.Params
	o := orm.NewOrm()
	pass := utils.MakePassword(password)
	sql := "select * from core_user where username = ? and password = ?"
	num, err := o.Raw(sql, username, pass).Values(&maps)
	if err == nil && num == 1 {
		return true
	} else {
		return false
	}
}

func ChangePassword(username, password string) error {
	o := orm.NewOrm()
	pass := utils.MakePassword(password)
	sql := "update core_user set password = ? where username = ?"
	_, err := o.Raw(sql, pass, username).Exec()
	return err
}
