package main
//要求：复习本节代码。
//有时间包，json等知识。
//思路：实现一个宝可梦游戏训练师面板。（想搞出来一个有可玩性的游戏，不是一时兴起。>_<）
//训练师面板：
//姓名
//性别
//宝可梦
//角色日志：（不同剧情（jsosn omitempty编码））
//成就
//并设置一个每半个小时就会输出“请休息一下，眼睛”，但是因为是作业看不出效果，改为3秒钟
//只是模块，主要是太麻烦了，这是个练习作业，真不是我懒。（主要是我现在很多东西都还不会）


//引入包
import (
	"encoding/json"
	"fmt"
	"time"
)

//定义结构体
type trainer struct {
	Name string
	Sex string
	Demon []string `json:"Demon,omitempty"`
	Journal []string  `json:"Journal,omitempty"`
	Achievement []string `json:"Achievement,omitempty"`
}

//显示
func show (t *trainer) string{
	x,_:=json.Marshal(t)
	return string(x)
}


//提醒
func remind() {
	tiker:=time.NewTicker(3*time.Second)
	for {
		<-tiker.C
		fmt.Println("请休息一下眼睛")
	}
}

//初始化训练师ddz(๑•̀ㅂ•́)و✧
var ddz trainer
func main(){
	ddz.Name="ddz"
	ddz.Sex="ddz"
	ddz.Demon=append(ddz.Demon,"皮卡皮卡")
	ddz.Achievement=append(ddz.Achievement ,"后端男神")
	x:=show(&ddz)
	fmt.Println(x)
	remind()
}















