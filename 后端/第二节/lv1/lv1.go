package main
/*
视频详情页：
作者部分
相关推荐部分
弹幕列表
视频部分

作者部分：
头像
昵称
签名
关注
是否会员

相关推荐部分
结构体切片

视频部分
标题
播放
时间
点赞数
投币数
收藏数
转发数
 */

import "fmt"

//定义视频详情页主体
type videoDetail struct {
	up up
	recommends []recommend
	barrage string
	video video
}
//定义作者模块
type  up struct {
	name string
	introduce string
	attention int64
	vip bool
}
//定义推荐
type recommend struct {
	title string
	up1 string
}
//定义视频模块
type video struct {
	title string
	play int
	time string
	click int
	coin int
	collect int
	repost int
	introduce string
}

//初始化
func main(){
	var kaizige videoDetail
	kaizige.up.name="qwqVictor"
	kaizige.up.introduce="     "
	kaizige.up.attention=45
	kaizige.up.vip=false
	kaizige.video.title="红岩网校工作站运维2020第七节课《Linux网络管理基础》录播"
	kaizige.video.play=255
	kaizige.video.time="2020-12-05 20:01:09"
	kaizige.video.click=11
	kaizige.video.coin=10
	kaizige.video.collect=9
	kaizige.video.repost=1
	kaizige.barrage="       "
	kaizige.video.introduce="这节课人好少啊（\n卑微见习讲师前几天刚拔完牙声音有一点点不清楚，\n以及我不该用 Windows XP 做客户机 demo 的（破音\n啊总的来说就是这样吧，不懂的可以事后来问我的。\n课程资源地址：https://fileshare.qwq.ren:2/Victor/sre-course-7/"
	kaizige.recommends=[]recommend{{ "Linux运维14天完全入门（学不会我删视频）", "酷python"},
		{ "网络运维基础教程，运维快速上手", "达内网络安全学院"},
		{"网络管理与维护", "571my"},
		{ "【红岩网校 | 产品策划及运营】", "红岩网校工作站"},
		{"Linux 系统与网络管理(2021)", "黄药师漫步桃花岛"}}
	fmt.Println(kaizige.video.title)
	fmt.Print(kaizige.video.time)
	fmt.Println("\t\t",kaizige.up.name)
	fmt.Println("\t\t\t\t\t\tup签名:\n\t  ",kaizige.up.introduce)
	fmt.Println("\t\t\t\t\t\t",kaizige.up.attention)
	fmt.Println("\t\t\t",kaizige.barrage,"  ")
	fmt.Print(kaizige.video.play,"  ")
	fmt.Print(kaizige.video.click,"  ")
	fmt.Print(kaizige.video.coin,"  ")
	fmt.Print(kaizige.video.collect,"  ")
	fmt.Print(kaizige.video.repost,"  \n")
	for i:=0;i<5;i++ {
		fmt.Println(kaizige.recommends[i])
	}
}