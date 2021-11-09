package main

import (
	"fmt"
	"time"
)



//定义两个切片,两个字符串
var(
input,output string
transit,transits []string
)

func main(){

//输入
fmt.Scanln(&input)

//字符串遍历，赋值给切片
for _,v:=range input{
transit=append(transit,string(v))
}
//切片进行倒序处理
transits=make([]string,len(transit))
	var (
		i=len(transit)-1
		j=0
	)
for ;i>=0;{
transits[j]=transit[i]
	i--
	j++
}
//再将切片转化为字符串
for _,v:=range transits{
output+=v
}

//输出
fmt.Println(output)
time.Sleep(5*time.Second)
}