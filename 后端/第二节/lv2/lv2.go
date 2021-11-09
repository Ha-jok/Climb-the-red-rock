package main

import (
	"fmt"
)

//定义视频模块
type video struct {
	name string
	title string
	click int
	coin int
	collect int
	repost int
}

//发布视频函数
func upvideo (n,t string) video {
	var vi video
	vi.name=n
	vi.title=t
	return vi
}


//三连方法
func (v *video) cl() {
	v.click++
}
func (v *video) co()  {
	v.coin++
}
func (v *video) re() {
	v.repost++
}
func (v *video) three() {
	v.click++
	v.coin++
	v.repost++
}

func main ()  {
	var (
		vid video
		up,title string
	)
	//发布视频
	fmt.Scanln(&up,&title)
	vid=upvideo(up,title)
	//三联系列
	fmt.Println(vid)
	var in int
	for  {
		fmt.Println("请选择：1，一键三联；2，点赞；3，投币；4，转发 ")
		fmt.Scanf("%d",&in)
		switch in {
		case 1:
			vid.three()
		case 2 :
			vid.cl()
		case 3:
			vid.co()
		case 4:
			vid.re()
		default:
			fmt.Println("error")
		}
		fmt.Println(vid)
	}
}













