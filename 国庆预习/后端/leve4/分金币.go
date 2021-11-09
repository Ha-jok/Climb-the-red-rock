package main
/*
你要发金币了，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后发出去了多少金币？
*/
//创建数组，一个储存每个人的信息，一个储存每个人的金币，并声明一个整型变量，记录发出去的金币量
import (
	"fmt"
	"time"
)
var (
	users =[...]string{"Matthew","Sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Elizabeth"}
	gcs,gcu,gcus int

)
//计算每个人所分金币，记录获取金币量，并将金币数量累加.遍历，判断
func main(){
	for _,u:= range users{
		gcus=0
		gcu=0
		for  _,v:= range u{
			switch v {
			case 'e':
				{
					gcu += 1
				}
			case 'E':
				{
					gcu += 1
				}
			case 'i':
				{
					gcu += 2
				}
			case 'I':
				{
					gcu += 2
				}
			case 'o':
				{
					gcu += 3
				}
			case 'O':
				{
					gcu += 3
				}
			case 'u':
				{
					gcu += 4
				}
			case 'U':
				{
					gcu += 4
				}
		}
		gcus+=gcu
		}
		gcs+=gcu
		fmt.Printf("%s分到的金币为%d\n",u,gcu)
	}
		fmt.Printf("发出去的金币为%d\n",gcs)
	time.Sleep(5*time.Second)

}
