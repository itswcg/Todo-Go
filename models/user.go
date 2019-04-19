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

func GetUser(user_id int64) (interface{}, error) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql := "select * from auth_user where id = ?"
	num, err := o.Raw(sql, user_id).Values(&maps)
	if err == nil && num == 1 {
		return maps[0]["username"], nil
	} else {
		return "", err
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

func CheckPassword(username, password string) (interface{}, bool) {
	var maps []orm.Params
	o := orm.NewOrm()
	pass := utils.MakePassword(password)
	sql := "select * from auth_user where username = ? and password = ?"
	num, err := o.Raw(sql, username, pass).Values(&maps)
	if err == nil && num == 1 {
		return maps[0]["id"], true
	} else {
		return 0, false
	}
}

func ChangePassword(username, password string) error {
	o := orm.NewOrm()
	pass := utils.MakePassword(password)
	sql := "update auth_user set password = ? where username = ?"
	_, err := o.Raw(sql, pass, username).Exec()
	return err
}

func CheckSession(sessionid string) (interface{}, error) {
	var maps []orm.Params
	o := orm.NewOrm()
	sql := "select * from django_session where session_key = ? and expire_date >= ? order by expire_date desc"
	cur_time := time.Now()
	num, err := o.Raw(sql, sessionid, cur_time).Values(&maps)
	if err == nil && num == 1 {
		return maps[0]["session_data"], nil
	} else {
		return nil, err
	}
}

func AddSession(sessionid string, user_id int64) (int64, error) {
	o := orm.NewOrm()
	expire_date := time.Now().AddDate(0, 1, 0)
	sql := "insert into django_session(session_key, session_data, expire_date) values(?, ?, ?)"
	res, err := o.Raw(sql, sessionid, user_id, expire_date).Exec()
	if err != nil {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func UpdateSession(sessionid string) error {
	o := orm.NewOrm()
	expire_date := time.Now().AddDate(0, 1, 0)
	sql := "update django_session set expire_date = ? where session_key = ?"
	_, err := o.Raw(sql, expire_date, sessionid).Exec()
	if err != nil {
		return err
	} else {
		return nil
	}
}
