package service

import (
	"FinalExam/业务题/3G含金量代码/dao"
	"fmt"
)

//通过原密码修改密码，传入用户名原密码和新密码，返回错误和一个布尔值.供api层调用
func Recharge (username string,money int)(error,bool){
	//	读取数据库中的信息。
	user,err := dao.ReadUser(username)
	//判断原密码是否正确，返回一个布尔值。若布尔值为true，调用dao层的updatepassword函数
	if b == true{
		dao.Recharge(username,money)
	}
	return err,b
}
