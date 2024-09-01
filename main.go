package main

import (
	"fmt"
)

func main() {
	b := Board{}
	b.Initialize()
	fmt.Println(b)
	b.Play(4,2,White)
	fmt.Println(b)
	b.Play(5,2,Black)
	fmt.Println(b)
}