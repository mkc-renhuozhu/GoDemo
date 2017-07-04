package common

import (
	"fmt"
	"os"
)

const Pi int = 100

var str string="test"

var (
	id int=1
	name string=""
	value int=2
)

var(
	APPNAME=string(os.Getenv("app_name"))
)

func Init() {
	fmt.Println(str)
	fmt.Println(APPNAME)
}