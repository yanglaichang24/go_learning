package bean

import "fmt"

func (u *User) Working() {
	fmt.Println(u.Name + "is woring...")
}

func Default() *User {
	return &User{
		Name: "default",
		age:  0,
		addr: "dddd",
	}
}
