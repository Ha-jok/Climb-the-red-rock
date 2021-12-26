package dao


//修改密码，传入用户名和新密码。返回错误。供service层调用
func Recharge(username string,money int)error{
	db,err := LinkMysql()
	user := ReadUser(username)
	sqlstr := "update users set money=? where username=?"
	_,err = db.Exec(sqlstr,user.Money+money,username)
	return err
}