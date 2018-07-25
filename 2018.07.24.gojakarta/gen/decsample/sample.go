package user

import "fmt"

type UserServiceLog struct {
	Origin UserService
}

func (al *UserServiceLog) Login(u User) (Token, error) {
	defer fmt.Println("Login called")
	return al.Origin.Login(u)
}
