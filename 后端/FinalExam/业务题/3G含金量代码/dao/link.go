package dao

import "database/sql"

func LinkMysql()(*sql.DB,error){
	//定义一个字符串，供各个函数调用，来连接数据库
	var Str = "root:2002@/finallyexam"
	db,error := sql.Open("mysql",Str)
	return db,error
}