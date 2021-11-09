package main

//复习，并使用结构体
//思路：创建学生属性结构体，利用rand函数，给学生生成学号、班级、姓名、性别。假设有10个学生。

//引入包
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
//定义结构体
type  student struct {
	name string
	sex string
	number string
	class int
}

//创建列表，
var  students [10]student
var sur =[5]string{"赵","钱","孙","李","周"}
var sex =[2]string{"男","女"}
var foreg  =[5]string{"雪","莹","佳","莉","柯"}
var foreg1 =[5]string{"雪","莹","佳","莉","柯"}
var foreb  =[5]string{"昊","达","俊","嘉","伟"}
var foreb1  =[5]string{"凯","升","皓","威",""}

//随机生成姓名,性别
func name () (string,string) {

	var n string
	var i int
	i=rand.Intn(2)
	s:=sex[i]
	i=rand.Intn(5)
	n+=sur[i]
	if  s=="男" {
		i=rand.Intn(5)
		n+=foreb[i]
		i=rand.Intn(5)
		n+=foreb1[i]
	}
	if s=="女" {
		i=rand.Intn(5)
		n+=foreg[i]
		i=rand.Intn(5)
		n+=foreg1[i]
	}
	return s,n
}


//随机生成学号
func number() string {
	var n string
	for x:=0;x<4;x++ {
	     i:=rand.Intn(10)
		 n+=strconv.Itoa(i)
	}
	return n
}


//随机生成班级
func class () int {
	i:=rand.Intn(9)
	return i+1
}


//主函数
func main ()  {
	//设置十次循环
	for i:=0;i<10;i++ {
		rand.Seed(time.Now().Unix()+int64(i))
		s,n:=name()
		students[i].sex=s
		students[i].name=n
		nu:="202121"
		nu+=number()
		students[i].number=nu
		c:=class()
		students[i].class=c
		fmt.Println(students[i])
	}
}

