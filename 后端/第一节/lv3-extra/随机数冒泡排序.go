package main

import (
	"fmt"
	"math/rand"
	"time"
)

//声明切片s
var (
s=make([]int,0)
transit int
)

func main() {

//生成随机数100个,储存到切片
for i:=0;i<100;i++{
rand.Int()
s=append(s,rand.Int())
}

//冒泡排序法
for i:=0;i<=98;i++ {
	for j := i + 1; j <= 99; j++ {
		if s[i] < s[j] {
			transit = s[i]
			s[i] = s[j]
			s[j] = transit
		}
	}
}
fmt.Println(s)
time.Sleep(5*time.Second)
}
