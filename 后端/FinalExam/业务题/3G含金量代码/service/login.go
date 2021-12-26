package service

import "FinalExam/业务题/3G含金量代码/dao"

//创建登录主函数，传入用户输入信息用户名和密码哦，接收dao层的结构体和错误,将错误返回。
//供api层中的接口调用

func Login(un,pw string)(error,bool){
	u,err:=dao.ReadUser(un)
	b := JudgePassword(pw,u.Password,u.Username,"登录")
	return err,b
}
