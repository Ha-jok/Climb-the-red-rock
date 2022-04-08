package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main(){
	//建立文件进行储存
	fileName := "./计算机网络/lv1/段子.txt"
	d,_ := os.Create(fileName)
	//获取网页内容
	url := "http://xiaodiaodaya.cn/"
	content := HttpGet(url)
	//通过正则表达式，获取段子对应的url
	r,err := regexp.Compile("(?U)a href=\"(.*)\"")
	if err != nil {

	}
	urlJoke := r.FindAllString(content,-1)
	for _,v := range urlJoke{
		//判断有效段子的url
		if len(v)>17{
			//提取url中的有效字段
			r2,err := regexp.Compile("/(.*)\"")
			if err != nil{
				fmt.Println("errUrlJoke:",err)
			}
			v1 := r2.FindString(v)
			v1 = strings.Replace(v1,"\"","",1)
			//调用函数通过完整的url获取网页内容
			urlWhole := url+v1
			contentJoke := HttpGet(urlWhole)
			//通过正则表达式获取有效片段
			r3,_ := regexp.Compile("<!--listS-->(.*)<!--listE-->")
			contentJokes := r3.FindAllString(contentJoke,-1)
			//通过调用正则表达式筛选有用信息
			for _,v := range contentJokes{
				match1 := "(?U)<img(.*)/>"
				v := ReplaceString(match1,v,"")
				match2 := "<!--listS-->"
				v = ReplaceString(match2,v,"")
				match3 := "<br/>"
				repl := "\n"
				v = ReplaceString(match3,v,repl)
				fmt.Println(v)
				d.WriteString(v)

			}
		}
	}
}

func HttpGet(url string)string{
	resp,err := http.Get(url)
	if err != nil{
		fmt.Println("errHttpGet:",err)
	}
	content1, _ := ioutil.ReadAll(resp.Body)
	content := string(content1)
	return content
}

func ReplaceString(match,src,repl string)string{
	r,_ := regexp.Compile(match)
	v := r.ReplaceAllString(src,repl)
	return v
}