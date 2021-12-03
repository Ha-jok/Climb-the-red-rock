package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type data struct{
	username string `json:"Username,omitempty"`
	password string `json:"Password,omitempty"`
}
var s string

func writ (d data,name string){

	s=s+" "+string(d.username)+" "+string(d.password)
	fmt.Println(s)
	x,_:=json.Marshal(s)
	fmt.Println(string(x))
	file,err:=os.OpenFile(name,os.O_WRONLY,0000)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_,err=file.Write(x)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file.Close()
}

func main(){
	var userName,passWord string
	fmt.Scanln(&userName,&passWord)
	var datas data
	datas.username="用户名"
	datas.password="密码"
	writ(datas,"./src/prac/homework/第四节/testDataTxt/testTxt.txt")
	writ(datas,"./src/prac/homework/第四节/testDataTxt/testData.data")
}


