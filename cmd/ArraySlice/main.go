package main

import "fmt"

//定义数组：

//1.定义，需要指定长度

var IntArr [5]int

//2.定义并初始化

var IntArr2 = [5]int{1,2,3,4,5}

//3.定义并初始化

var IntArr3 = [...]int{1,2,3}

//4. 遍历数组
func RangeArr(arr [5]int)  {
	for key, value := range arr {
		fmt.Printf("index:%d,value:%d\r\n",key,value)
	}
}

//定义Slice:

//1.定义,不需要指定长度

var IntSlice []int

//2. 定义,使用make并指定长度

var IntSlice2=make([]int,5)

//3.定义并初始化

var IntSlice3 = []int{1,2,3}

//4.定义并初始化

var IntSlice4 = IntArr2[:]

//5.遍历Slice

func RangeSlice(arr []int)  {
	for key, value := range arr {
		fmt.Printf("index:%d,value:%d\r\n",key,value)
	}
}

func main() {
	fmt.Println("range arr...")
	RangeArr(IntArr2)
	fmt.Println("range slice ...")
	RangeSlice(IntSlice3)
}
