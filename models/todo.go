package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	IS_DO = 0
)

func AddTodo(user_id int64, content string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into core_todo(content, create_date, is_do, author_id) values(?, ?, ?, ?)"
	cur_time := time.Now()
	res, err := o.Raw(sql, content, cur_time, IS_DO, user_id).Exec()
	if res != nil {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func GetAllTodo(user_id int64) ([]orm.Params, error) {
	var maps []orm.Params
	sql := "select * from core_todo where author_id = ?"
	o := orm.NewOrm()
	num, err := o.Raw(sql, user_id).Values(&maps)
	if err == nil && num > 0 {
		return maps, nil
	} else {
		return nil, err
	}
}

func GetTodo(user_id int64, is_do, gte bool) ([]orm.Params, error) {
	var maps []orm.Params
	var sql string
	if gte == true {
		sql = "select * from core_todo where author_id = ? and is_do = ? and create_date >= ? order by id desc"
	} else {
		sql = "select * from core_todo where author_id = ? and is_do = ? and create_date < ? order by id desc"
	}
	today := time.Now().Format("2006-01-02")
	o := orm.NewOrm()
	num, err := o.Raw(sql, user_id, is_do, today+" 00:00:00").Values(&maps)
	if err == nil && num > 0 {
		return maps, nil
	} else {
		return nil, err
	}
}

func UpdateTodo(todo_id int64, content string) error {
	sql := "update core_todo set content = ? where id = ?"
	o := orm.NewOrm()
	_, err := o.Raw(sql, content, todo_id).Exec()
	return err
}

func DoTodo(todo_id int64, do int) error {
	sql := "update core_todo set is_do = ? where id = ?"
	o := orm.NewOrm()
	_, err := o.Raw(sql, do, todo_id).Exec()
	return err
}

func DeleteTodo(todo_id int64) error {
	sql := "delete from core_todo where id = ?"
	o := orm.NewOrm()
	_, err := o.Raw(sql, todo_id).Exec()
	return err
}

func SearchTodo(content string) ([]orm.Params, error) {
	var maps []orm.Params
	sql := "select * from core_todo where content like ?"
	o := orm.NewOrm()
	num, err := o.Raw(sql, "%"+content+"%").Values(&maps)
	if err == nil && num > 0 {
		return maps, nil
	} else {
		return nil, err
	}
}
