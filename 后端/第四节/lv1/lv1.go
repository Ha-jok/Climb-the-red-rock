package main
//对题干进行分析：要求用户指定函数releaseSkillFunc
//func main() {
//	ReleaseSkill("龙卷风摧毁停车场", func(skillName string) {
//		fmt.Println("尝尝我的厉害吧！", skillName)
//	})
//}
//func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
//	releaseSkillFunc(skillNames)
//}

//引入包
import (
	"fmt"
	"time"
)

//编写输出模板函数
func output1(s string){
	fmt.Println("欧拉欧拉欧拉欧拉欧拉",s)
}
func output2(s string){
	fmt.Println("木大木大木大木大木大",s)
}
func output3(s string){
	fmt.Println("砸瓦鲁多",s)
}

//releaseskill函数
func ReleaseSkill(skillNames string,releaseSkillFunc func(string)){
	releaseSkillFunc(skillNames)
}

//主函数
func main(){
	for {
		//	用户指定文字模板
		fmt.Println("请选择输出模板:\n1,欧拉欧拉欧拉欧拉欧拉\n2,木大木大木大木大木大\n3,砸瓦鲁多")
		var out string
		fmt.Scanf("%s", &out)
		out = "output" + out
		//匹配技能
		fmt.Println("请输入技能")
		var skillName string
		fmt.Scanf("%s", &skillName)
		switch out {
		case "output1":
			ReleaseSkill(skillName, output1)
		case "output2":
			ReleaseSkill(skillName, output2)
		case "output3":
			ReleaseSkill(skillName, output3)
		}
		time.Sleep(1*time.Second)
	}
}