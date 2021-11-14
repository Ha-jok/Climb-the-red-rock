package main

import "fmt"

//用户可以自定义每个技能使用的文字模板，并且用户在释放技能前不再需要指定输出文字的模板
//用户可以添加自定义的技能名字，判断用户输入的技能名字中是否有敏感词（敏感词词库自行设置），如果包含敏感词则输出对应的提示文字并取消本次添加。添加后也可以自定义每个技能使用的文字模板
//用户可以添加自定义的模板，并且同样需要对敏感词鉴权，添加后用户可以使用该模板。
//希望模板的自定义程度不要仅仅局限于只能自定义文字，也要可以自定义技能名出现的位置和次数，如："你死定了！"+技能名+"我不知道你该怎么从"+技能名+"活下来~"

//分析要求：1.自定义输出模板；2.自定义技能名 3.自定义输出位置
//思路：
//不会添加函数，所以排除原始难度做法
//将输出模板作为一个切片进行储存，技能名同理
//运行初始界面为：选择技能和选择输出模板
//选择后添加自定义选项。


//releaseskill函数移植
func ReleaseSkill(skillNames string,releaseSkillFunc func(string)){
	releaseSkillFunc(skillNames)
}
//定义releaseSkillFunc函数，设置技能前字符和后字符
var be,la,cen string
func releaseSkillFunc (cen string){
	fmt.Println(be,cen,la)
}

//实现功能1
var skilltext []string
func setOutput (s string){
	skilltext=append(skilltext,s)
}

//实现功能2
var skill []string
func setSkill(s string){
	skill=append(skill,s)
}







func main (){
	//初始化技能库和输出模板库
	skill=append(skill,"流星指刺","波纹疾走","光之流法")
	skilltext=append(skilltext," ","欧拉欧拉欧拉欧拉欧拉","木大木大木大木大","JoJo")

	//主体循环开始
	var i int
	i=0
	for i==0 {
		fmt.Println("1.选择输出模板\n2.选择技能\n3.退出程序")
		var sele int
		fmt.Scanf("%d",&sele)
		switch sele {

		//选择输出模板，进行循环
		case 1:{
			fmt.Println("1:定义输出前模板\n2:定义输出后模板")
			var sele2 int
			fmt.Scanln(&sele2)
			switch sele2 {
			case 1:{
				for i,j:=range skilltext{
					fmt.Println(i+1,":",j)
				}
				i:=len(skilltext)+1
				fmt.Println(i,":自定义模板")
				var sele3 int
				fmt.Scanln(&sele3)
				if sele3==i{
					fmt.Print("输入模板：")
					var str string
					fmt.Scanln(&str)
					setOutput(str)
				}
				be=skilltext[sele3-1]
			}
			case 2:{
				for i,j:=range skilltext{
					fmt.Println(i+1,":",j)
				}
				i:=len(skilltext)+1
				fmt.Println(i,":自定义模板")
				var sele3 int
				fmt.Scanln(&sele3)
				if sele3==i{
					fmt.Print("输入模板：")
					var str string
					fmt.Scanln(&str)
					setOutput(str)
				}
				la=skilltext[sele3-1]
			}
			//选择后模板结束
			}
		//定义前后模板结束
		}
		//选择模板结束


		//选择技能，直接输出技能
		case 2:{
			for i,j:=range skill{
				fmt.Println(i+1,":",j)
			}
			i:=len(skill)+1
			fmt.Println(i,":自定义技能")
			var sele1 int
			fmt.Scanln(&sele1)
			if sele1==i{
				fmt.Print("输入技能名：")
				var str string
				fmt.Scanln(&str)
				setSkill(str)
			}
			cen=skill[sele1-1]
			ReleaseSkill(cen,releaseSkillFunc)

		}
		//选择技能结束

		case 3:
			i++
		}
	//	选择结束
	}
//	循环结束
}
//主函数结束



