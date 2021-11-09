package main

import"fmt"

//定义三个整数，一个字符串，一个整型切片
var(
inputNum1,inputNum2,result int
cmd string
history []int
)

func main() {

  //无限循环使程序多次执行
  for {

    //输入
    fmt.Scanln(&inputNum1,&cmd,&inputNum2)

    //switch语句进行判断
    switch cmd {
    case "+":
        result=inputNum1+inputNum2
    case "-":
        result=inputNum1-inputNum2
    case "*":
        result=inputNum1*inputNum2
    default:
        fmt.Println("error")
    }

    //记录本次结果
    history=append(history,result)
    //输出
    fmt.Println(inputNum1,cmd,inputNum2,"=",result)
    fmt.Println(history)
  }
}

