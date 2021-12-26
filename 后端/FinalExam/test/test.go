package main

import "fmt"


//分析，创建一个二维切片储存用户输入的关于房间的信息
//house[行数]信息
func inputHouse(housein [][2]int,nm int)(housein [][2]int)  {
	var housetran [2]int
	for i:=0;i<nm;i++{
		fmt.Scanf("%d %d",housetran)
		housein = append(housein,housetran)
	}
}

//确定(i,m)房间的情况。
func housei(i,m,M int,housein [][2]int)(housei [2]int){
	var housei [2]int
	//	确定，（i,m)房间所在行数
		houseinrow:=(i-1)*M+m+1
		housei=housein[houseinrow]
		return housei
}
//确定第i层有楼梯的房间
func housefloor(floor [][2]int,M int) (floo []int){
	var floo []int
	for m:=0;m<M;m++{
		if floor[m][1]==1{
			floo = make(floo,m)
		}
	}
	return floo
}

func main()  {
	var N,M,start,password int
	var house [][][2]int
	var housefloors [][]int
	var floorstart  [2]int
	//确定，N,M
	fmt.Scanf("%d %d",&N,&M)
	housein:=inputHouse(N*M)
	fmt.Scanf("%d",&start)
	//分别执行不同层数的
	for i:=1;i<=N;i++{
		//确定第i层的房间情况
		for m:=0;m<M;m++{
			housei:= housei(i,m,M,housein)
			house=append(house[i][m],housei)
		}
		//确定第i层有楼梯的房间
		housefloors[i]=housefloor(house[i],M)
	  //确定第i层初始房间的情况。
	  floorstart=house[i][start]
	  //将指示数加入密码
	  password+=floorstart[1]
	  //确定下一层初始楼层
	  if floorstart[0]!=0{
		  for x,j:=range housefloors[i]{
			  if j==start{
				  start+=j
			  }
		  }
	  }
	}

}