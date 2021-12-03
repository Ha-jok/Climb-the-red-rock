package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//要求
//在2的基础上实现留言
//分析，参考贴吧模式：吧包含贴，贴包含楼，楼包含回复
//因为只有人与人，所以将每两个人建立一张数据库，作为两人的留言交互
//考虑实际，当两个人第一次交互时创建一张表。
//考虑作业，只设置两个人，创建一个存在的留言表，共数据库中两个人交互。
//将lv2代码移植到下面。

//将lv2的连接数据库和exec错误封装放在上方，以便引用
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

//创建符合留言板数据表的结构体
type liuyan struct {
	number int
	name   string
	str    string
}

//读取留言板所有内容
func liuyanban(db *sql.DB) {
	sqlstr := "select number,name,string from lv3liuyan where number>?"
	row, err := db.Query(sqlstr, 0)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer row.Close()
	for row.Next() {
		var ly liuyan
		err = row.Scan(&ly.number, &ly.name, &ly.str)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(ly.name, ":", ly.str)
	}
	return
}

//继续留言
func liuyanNEXT(db *sql.DB, us string) {
	fmt.Print("请输入您要留言的内容：")
	var lycontext string
	fmt.Scan("%s", &lycontext)
	sqlstr := "insert into lv3liuyan(name,string) values(?,?)"
	ret, err := db.Exec(sqlstr, us, lycontext)
	er(ret, err)
}

//创建和数据表结合的结构体
type user struct {
	username string
	password string
	mibao    string
	answer   string
}

//数据库相关函数

//读取某个用户的信息,返回一个结构体
func readmysql(db *sql.DB, us string) user {
	//封装语句，准备从数据库中读取数据，
	sqlstr := "select username,password,mibao,answer from lv3user where username = ?"
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

//更新用户信息，
func updatemysql(db *sql.DB, us, pw string) {
	sqlstr := "update lv3user set password=? where username=?"
	ret, err := db.Exec(sqlstr, pw, us)
	er(ret, err)
	fmt.Println("修改成功")
}

//添加用户信息
func insertmysql(db *sql.DB, us, pw, mb, aw string) {
	sqlstr := "insert into lv3user(username,password,mibao,answer) values(?,?,?,?)"
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
		fmt.Println("1.退出\n2.修改密码\n3.查看留言版\n4.留言")
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
		case 3:
			{
				liuyanban(db)
			}
		case 4:
			{
				liuyanNEXT(db, us)
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
	sqlstr := "select password from lv3user where username=?"
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
	fmt.Println(u.mibao)
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

//多次测试，都无法输入带有空格的留言。问题未解决
