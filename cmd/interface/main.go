package main

import (
	"fmt"
	"reflect"
	"flag"
)

var (
	firstName,lastName string
)

type operation interface {
	Add()int
}

type Customer struct {
	height int
	weight int
}

type Person struct {
	sex int
}

func (c *Customer) Add()int {
	return c.height+c.weight
}

func (p *Person) Add()int  {
	return p.sex
}


var ip *int= flag.Int("flagname", 1234, "help message for flagname")

var name *string=flag.String("name","xxx","need to set name...")

func main() {

	flag.Parse()

	fmt.Println(*ip)

	fmt.Println(*name)

	//who:="johnny zhu"
	//if len(os.Args)>1{
	//	who+=strings.Join(os.Args[1:], " ")
	//}
	//fmt.Println("hello",who)

	c:=&Customer{height:1,weight:1}
	p:=&Person{sex:3}

	a:=[]operation{c,p}
	for _, v := range a {
		fmt.Println("current operation:",v)
		fmt.Println("add result:",v.Add())
	}

	var f float32
	f = 12.0

	fmt.Println(reflect.ValueOf(f))
	fmt.Println(reflect.TypeOf(f))

	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.ValueOf(c))

	//fmt.Println("please enter your firstName and lastName...")
	//fmt.Scanln(&firstName,&lastName)
	//
	//fmt.Println(firstName,lastName)

}