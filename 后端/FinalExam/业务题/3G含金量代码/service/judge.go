package service

import (
	"fmt"
	"strings"
)

//判断密码,在后台打印用户是否成功的信息信息,并返回一个布尔值，反应密码是否正确，
func JudgePassword(inputPw,mysqlPw,username,service string)bool{
	if inputPw != mysqlPw{
		fmt.Println(username,"密码错误",service)
		b := false
		return b
	}else {
		fmt.Println(username,"密码正确",service)
		b := true
		return b
	}
}


func JudgeUsernameDuplicate(username,err string) (bool) {
	if err == strings.Replace("Error 1062: Duplicate entry '?' for key 'users.PRIMARY'","?",username,1){
		b := false
		return b
	}else {
		b := true
		return b
	}

}