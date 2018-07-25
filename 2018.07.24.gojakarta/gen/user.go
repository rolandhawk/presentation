package user

import (
	"fmt"
	"io"
	"math/rand"
)

type Token string

type User string

// GOGENERATE OMIT
//go:generate go run decorator/main.go UserService
// INTERFACE OMIT

type UserService interface {
	Login(User) (Token, error)
	Logout(Token) error
}

// INTERFACE OMIT
// GOGENERATE OMIT

type userService struct {
	LoginUser map[Token]User
}

func (us *userService) Login(u User) (Token, error) {
	token := fmt.Sprintf("%d", rand.Int)
	us.LoginUser[token] = u
	return u, nil
}

func (us *userService) Logout(t Token) error {
	delete(us.LoginUser, t)
	return nil
}

func New() UserService {
	return &userService{
		LoginUser: make(map[Token]User),
	}
}
