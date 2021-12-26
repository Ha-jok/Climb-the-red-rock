package service

import "FinalExam/业务题/3G含金量代码/dao"

func Regist(username,password string,money int)(error,bool){
	err := dao.InsertUser(username,password,money)
	if err != nil {
		b := JudgeUsernameDuplicate(username,err.Error())
		return err,b
	}
	b := true
	return err,b
}
