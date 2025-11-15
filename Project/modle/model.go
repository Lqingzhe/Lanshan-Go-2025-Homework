package manage

import "fmt"

type User struct {
	name     string
	password string
}

func (u *User) Adduser(uid int) {
	var name string
	var password string
	_, _ = fmt.Scan(&name)
	_, _ = fmt.Scan(&password)
	nu := User{name: name, password: password}

	Users[uid] = nu
}

func (u *User) DeleteUser(uid int) {
	delete(Users, uid)
}
func (u *User) Updatepassword(uid int) {
	var newpassword string
	_, _ = fmt.Scan(&newpassword)
	un := Users[uid]
	un.password = newpassword
	Users[uid] = un
}

type Manage interface {
	Adduser(int)
	DeleteUser(int)
	Updatepassword(int)
}

var Users = make(map[int]User)

func Test() {
	fmt.Println(Users)
}
