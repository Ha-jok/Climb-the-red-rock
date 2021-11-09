package main
import (
	"fmt"
	"time"
)
var(
	account,password string
	ap=make(map[string]string,200)
)
func main(){
	allap()
	fmt.Println("请输入账号:")
	fmt.Scanln(&account)
	_,ok:=ap[account]
	if !ok {
		fmt.Println("查无此账号")
	} else  {
		fmt.Println("请输入密码：")
		fmt.Scanln(&password)
		if password != ap[account] {
			fmt.Println("密码错误")
		} else {
			fmt.Println("密码正确")
		}
	}
	time.Sleep(5*time.Second)
}

func allap(){
	ap["username"]="password"
	ap["asfd"]="werf"
	ap["理想"]="juan"
	ap["ddz"]="xuejie"    //狗头😜😜😜
}