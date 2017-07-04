package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

var apiRateLimit  map[string]ratelimit.Limiter

func main() {

	apiRateLimit = make(map[string] ratelimit.Limiter)

	for i := 0; i < 10; i++ {
		userId:="00001"
		api(userId)
	}
}

func GetRateLimit(userId string) ratelimit.Limiter  {
	l:=apiRateLimit[userId]
	if l!=nil{
		return apiRateLimit[userId]
	}else {
		apiRateLimit[userId] = ratelimit.New(1)
		return apiRateLimit[userId]
	}
}

func api(userId string)  {
	prev := time.Now()
	l:=GetRateLimit(userId)

	now := l.Take()
	fmt.Println(userId, now.Sub(prev))
	prev = now

	//fmt.Printf("userId:%s test\n",userId)
}

func test() {
	rl := ratelimit.New(1) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}


