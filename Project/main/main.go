package main

import (
	"Project/modle"
	"fmt"
)

func main() {
	fmt.Println("输入账号，密码")
	var ma manage.Manage = &manage.User{}
	ma.Adduser()
	ma.Login(247915639)
	ma.Adduser()
	ma.Login(114514)
	manage.Test()
	fmt.Println("...............................") //这只是一个可有可无的，毫无存在感的分割线
	ma.DeleteUser(114514)
	manage.Test()
}
