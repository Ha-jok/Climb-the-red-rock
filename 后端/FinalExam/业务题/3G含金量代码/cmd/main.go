package main

import "FinalExam/业务题/3G含金量代码/api"

//创建一个数据表，储存
func main() {
	engine := api.CreateRout()
	//实现登录
	api.Login(engine)
	//注册接口
	api.Regist(engine)
   //充值接口

}