package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func GetTask(user_id int64) ([]orm.Params, error) {
	var maps []orm.Params
	sql := "select * from core_task where author_id = ? order by id desc limit 1"
	o := orm.NewOrm()
	num, err := o.Raw(sql, user_id).Values(&maps)
	if err == nil && num > 0 {
		return maps, nil
	} else {
		return nil, err
	}
}

func AddTask(user_id int64, content string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into core_task(content, timestamp, author_id) values(?, ?, ?)"
	timestamp := time.Now()
	res, err := o.Raw(sql, content, timestamp, user_id).Exec()
	if err != nil {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}
