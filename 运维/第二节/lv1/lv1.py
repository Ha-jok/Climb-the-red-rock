#分为主体和三个函数

# 函数：机器出拳
def robot():
    robotguess=random.randint(1,4)
    return robotguess



#函数：积分记录
def cont(a,b):
    c,d=0,0
    if a==b:
        c+=1
        d+=1       
    elif b!=1&a==b-1:
        c+=1
    elif b==1&a==3:
        c+=1
    else :
        d+=1
    return c, d

#函数：判断胜负
def judge(a,b):
    if a==b:
        print('平局')
    elif a>b:
        print('胜利')
    else :
        print('失败')


#主体，输入，判定
import random
possible={'石头':1,'剪刀':2,'布':3}
player_mark=0
yyz_mark=0
for x in range(3):
     print("请选择：石头，剪刀，布")
     player=input()
     yyz=robot()
     tran1,tran2 = cont(possible[player],yyz)
     player_mark+=tran1
     yyz_mark+=tran2
judge(player_mark,yyz_mark)











