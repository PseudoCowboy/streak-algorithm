package main

import (
	"fmt"
	"reflect"
)

type A struct {
}

type B interface {
}

func main() {
	a := A{}
	b := new(B)
	fmt.Println(reflect.TypeOf(a).String())
	fmt.Println(reflect.TypeOf(b).String())

	t := 1
	modify(&t)
	fmt.Println(t)
}

func modify(a *int) {
	*a += 1
}
