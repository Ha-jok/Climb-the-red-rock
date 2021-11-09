#分为抓取信息，解析信息，写入文件

#引入包
import urllib.request
import urllib.parse
import re
from bs4 import BeautifulSoup
import xlwt
from xlwt.Workbook import Workbook
from xlwt.Worksheet import Worksheet


#抓取信息
#设置网页
baseurl='http://www.4399.com'
#构建头部信息，请求
head={
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36 Edg/95.0.1020.40"
}
req=urllib.request.Request(url=baseurl,headers=head)
#抓取信息
response=urllib.request.urlopen(req)
html=response.read().decode('gbk')
#print(html),测试，抓取信息成功


#定义保存信息的列表
data=[]
#解析信息
bs=BeautifulSoup(html,'html')
#print(bs.title)，测试，可以解析，但是遇到问题,问题保留
aname=bs.find_all('a')
for gamename in aname:
    gamename=str(gamename)
    gamename1=re.findall('/>.*?</a>',gamename)
    gamename1=str(gamename1)
    if gamename1!='[]':
        if str(re.findall('游戏',gamename1))=='[]':
                 if str(gamename1[4:-6])!='':
                     data.append(gamename1[4:-6])
#print(data)测试，列表建立成功

#保存数据
#创建对象，表单
Workbook=xlwt.Workbook(encoding='gbk')
Worksheet=Workbook.add_sheet('4399游戏')
#写入信息
i=0
for v in data:
    Worksheet.write(i,0,v)
    i=i+1
#保存信息
Workbook.save('4399.xls')