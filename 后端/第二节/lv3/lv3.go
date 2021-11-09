package main

import (
	"fmt"

)

//接受函数，返回类型
func reci(x interface{}) {
	v:=fmt.Sprintf("%T", x)
	switch v {
	case "string":
		fmt.Println("这是一个string")
	case "int":
		fmt.Println("这是一个Int")
	case "float":
		fmt.Println("这是一个float")
	case "bool":
		fmt.Println("这是一个bool")
	}
}


