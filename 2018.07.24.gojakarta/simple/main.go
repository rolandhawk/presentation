package main

import "fmt"

//go:generate echo "hello GoJakarta"
//go:generate echo "simple go generate command"

func main() {
	fmt.Println("simple go generate")
}
