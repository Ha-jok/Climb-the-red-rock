package main

import (
	"fmt"
	"math/rand"
	"sort"
                "time"
)

//声明切片s
var s=make([]int,0)

func main() {

//生成随机数100个,储存到切片
for i:=0;i<100;i++{
rand.Int()
s=append(s,rand.Int())
}

//sort包排序
	sort.Sort(sort.Reverse(sort.IntSlice(s)))

fmt.Println(s)
time.Sleep(5*time.Second)
}