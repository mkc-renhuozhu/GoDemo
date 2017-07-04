package main

import "fmt"

//1. 定义struct

type Customer struct {
	Id int
	Name string
}

//2.类型等价定义
type str string

//3.定义接口类型

type operation interface {
	Add()int
	Remove()int
}

//4.定义函数类型

type handler func(param string)error

func Opr(param string) error  {
	fmt.Println(param)
	return nil
}

func handleError(f func(param string)error) error {
	return handle(handler(f))
}

func handle(h handler) error {
	return h("xxx")
}

func main() {
	handleError(Opr)
}


