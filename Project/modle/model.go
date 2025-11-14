package manage

import "fmt"

type User struct {
	name     string
	password string
	uid      int
}

func (u *User) Adduser(uid int) {
	var name string
	var password string
	_, _ = fmt.Scan(&name)
	_, _ = fmt.Scan(&password)
	u.name = name
	u.password = password
	u.uid = uid
	Users[uid] = *u
}

func (u *User) DeleteUser(uid int) {
	delete(Users, uid)
}

type Manage interface {
	Adduser(int)
	DeleteUser(int)
}

var Users = make(map[int]User)

func Test() {
	fmt.Println(Users)
}
