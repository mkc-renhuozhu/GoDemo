package main

import (
	"errors"
	"fmt"
	"os"
)

var errNotFound error = errors.New("user Not found error")

func main() {

	defer func(){
		if e:=recover();e!=nil{
			fmt.Printf("panic %s \r\n",e)
		}
	}()

	var user = os.Getenv("TEST")
	if user==""{
		fmt.Printf("error: %v", errNotFound)
		panic("Unknown user: no value for $USER")
	}else {
		fmt.Println(user)
	}

	panic("this is...")
}
