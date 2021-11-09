#定义字符串和一个空列表
string='vimlnello redrockhtml'
list=[]
#将字符串遍历储存在列表
for x in string:
    list.append(x)
#删除多余字符
del list[:5]
del list[-4:]
#插入字符'h'
list.insert(0,'h')
#将列表重新转换为字符串，输出
output=''
for v in list:
    output+=v
print(output)