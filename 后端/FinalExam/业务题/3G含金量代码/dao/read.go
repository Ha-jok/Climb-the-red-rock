package dao

import (
	"FinalExam/业务题/3G含金量代码/model"
)

//读取某个用户的信息,返回一个结构体
func ReadUser(u string) (model.User,error){
	db,err := LinkMysql()
	sqlstr:="select * from users where username=?"
	var user model.User
	err=db.QueryRow(sqlstr,u).Scan(&user.Username,&user.Password,&user.Money)
	return user,err
}
