package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//作业要求
//登录+注册
//可以通过密保（比如：你所就读的是哪一所高中，父亲叫什么名字什么的都行）来找回账号或者更改密码
//分析
//由于数据库的存在，不需要将所有数据提取，利用函数临时读取数据库相应内容返回。
//登录功能：
//正常登录，使用数据库匹配账号密码
//忘记密码，使用密保匹配答案。
//修改密码功能，利用update
//注册功能：
//创建写入函数，密保可以为空，若密保为空提示用户注意。
//使用参数传递信息

//创建和数据表结合的结构体
type user struct {
	username string
	password string
	mibao    string
	answer   string
}

//数据库相关函数
//创建sql.DB,并检查数据库连接
func createDB() *sql.DB {
	db, err := sql.Open("mysql", "root:2002@/redrockwork")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//读取某个用户的信息,返回一个结构体
func readmysql(db *sql.DB, us string) user {
	//封装语句，准备从数据库中读取数据，
	sqlstr := "select username,password,mibao,answer from lv2 where username = ?"
	var u user
	err := db.QueryRow(sqlstr, us).Scan(&u.username, &u.password, &u.mibao, &u.answer)
	if err != nil {
		var error string
		error = "sql: Scan error on column index 2, name \"mibao\": converting NULL to string is unsupported"
		er := err.Error()
		if er == error {
			return u
		} else {
			fmt.Println("请联系管理员")
		}
	}
	return u
}

//封装exec中错误处理函数，
func er(ret sql.Result, err error) {
	if err != nil {
		fmt.Println("请联系管理员")
		fmt.Println(err)
		return
	}
	_, err = ret.RowsAffected()
	if err != nil {
		fmt.Println("请联系管理员")
		return
	}
}

//更新用户信息，
func updatemysql(db *sql.DB, us, pw string) {
	sqlstr := "update lv2 set password=? where username=?"
	ret, err := db.Exec(sqlstr, pw, us)
	er(ret, err)
	fmt.Println("修改成功")
}

//添加用户信息
func insertmysql(db *sql.DB, us, pw, mb, aw string) {
	sqlstr := "insert into lv2(username,password,mibao,answer) values(?,?,?,?)"
	ret, err := db.Exec(sqlstr, us, pw, mb, aw)
	er(ret, err)
}

//登录函数，
func enter(db *sql.DB) {
	//初始界面
	fmt.Print("请输入账号:")
	var us string
	fmt.Scanf("%s", &us)
	//从数据库中提取信息，和用户输入进行比较
	u := readmysql(db, us)
	if u.password!=""{
		fmt.Println("该用户未注册")
		return
	}
	fmt.Print("请输入密码：")
	var pw string
	fmt.Scanf("%s", &pw)
	//判断密码是否正确
	if pw != u.password {
		fmt.Println("密码错误")
	} else {
		fmt.Println("成功登录")
		fmt.Println("1.退出\n2.修改密码")
		var i int
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			{
				return
			}
		case 2:
			{
				fmt.Print("请输入新密码：")
				fmt.Scanf("%s", &pw)
				updatemysql(db, us, pw)
			}

		}
	}
}

//注册函数
func regist(db *sql.DB) {
	var us, pw, mb, aw string
	fmt.Print("请输入注册的用户名:")
	fmt.Scanf("%s", &us)
	s := registCheck(db, us)
	if s != "" {
		fmt.Println("用户名重复，请使用其他用户")
		return
	}
	fmt.Print("请输入密码:")
	fmt.Scanf("%s", &pw)
	fmt.Print("请输入密保:")
	fmt.Scanf("%s", &mb)
	fmt.Print("请输入密保答案:")
	fmt.Scanf("%s", &aw)
	insertmysql(db, us, pw, mb, aw)
}

//判断是否注册用户名是否重复
func registCheck(db *sql.DB, us string) string {
	sqlstr := "select password from lv2 where username=?"
	var s string
	_ = db.QueryRow(sqlstr, us).Scan(&s)
	return s
}

//找回密码函数
func retrievePassword(db *sql.DB) {
	fmt.Print("请输入用户名：")
	var us string
	fmt.Scanf("%s", &us)
	u := readmysql(db, us)
	if u.mibao != "" {
		fmt.Println(u.mibao)
		var s string
		fmt.Scanf("%s", &s)
		if s != u.answer {
			fmt.Println("密保答案错误")
		} else {
			fmt.Print("密保正确。\n请输入新密码：")
			var pw string
			fmt.Scanf("%s", &pw)
			updatemysql(db, us, pw)
		}
	} else {
		fmt.Println("您未设置密保")
	}

}

//主函数
func main() {
	//创建sql.DB
	db := createDB()
	//起始界面
	fmt.Println("1.登录\n2.注册\n3.找回密码")
	var i int
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		{
			//登录函数
			enter(db)
		}
	case 2:
		{
			//注册函数
			regist(db)

		}
	case 3:
		{
			//找回密码函数
			retrievePassword(db)
		}
	}

}

//日志：
//第一次测试,仅含创建*DB,出现错误，
//sql: unknown driver "mysql" (forgotten import?)
//起因为未引入Mysql框架
//引入mysql框架后，成功.
//第二次测试，从数据库中读取某用户信息，
//返回u结构体为空，
//起因，queryrow函数参数未输入。
//加入参数后，成功
//第三次测试，密码判定，成功。
//第四次测试，测试修改密码功能，发现错误，scan报错。
//起因是，密保为空，指定若发生该错误直接返回，其他错误则联系管理员，
//添加判定特定错误函数，成功
//第五次测试，补充测试修改密码功能，报错
//起因，输入新密码的时候未给出变量地址，导致密码不变
//Scanf加上&，成功
//第六次测试，测试用户名注册重复问题成功。
//第七次测试，测试用户名注册问题，成功
//第八次测试，测试通过密保找回密码功能，成功
//第九次测试，源于发现没有判定是否有账户存在