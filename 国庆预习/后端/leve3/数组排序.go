package main

import (
	"fmt"
	"log"
	//"sort"
	//感觉用sort有点儿流氓，不用又感觉傻傻的，😜😜😜，就注释掉了
	"time"
)
var (
	x int
	num=make([]int,0)
)
func main(){
	fmt.Println("input:")
	for  {
		_,err := fmt.Scanln(&x)
		if err!= nil {
			break
			log.Fatal(err)
		}
		num = append(num, x)
	}
	//sort.Ints(num)
	order()
	fmt.Println("output:")
	fmt.Println(num)
	time.Sleep(5*time.Second)
}
func order(){
	var q int
	for i:=0;i<len(num)-1;i++ {
		for j := i + 1;j<len(num) ; j++ {
			if num[i] > num[j] {
				q = num[i]
				num[i] = num[j]
				num[j] = q
			}
		}
	}
}

