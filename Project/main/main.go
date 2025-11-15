package main

import (
	"Project/modle"
	"fmt"
)

func main() {
	fmt.Println("输入账号，密码")
	var ma manage.Manage = &manage.User{}
	ma.Adduser(114514)    //输入 一号 1
	ma.Adduser(247915639) //输入 二号 2
	manage.Test()
	fmt.Println("...............................") //这只是一个可有可无的，毫无存在感的分割线
	//ma.DeleteUser(114514)
	//manage.Test()
	//fmt.Println("...............................") //这只是一个可有可无的，毫无存在感的分割线
	ma.Updatepassword(247915639) //输入22
	manage.Test()
	fmt.Println("...............................") //这只是一个可有可无的，毫无存在感的分割线
	ma.Updatepassword(114514)                      //输入11
	manage.Test()
}
