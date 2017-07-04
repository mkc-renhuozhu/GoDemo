package main

import (
	"fmt"
)

type TwoInts struct {
	A int
	B int
}

type IntVector []int

func (iv IntVector) For() {
	for index := 0; index < len(iv); index++ {
		fmt.Printf("index:%d,value:%d\r\n",index,iv[index])
	}
}

func (iv IntVector)Range()  {
	for key, value := range iv {
		fmt.Printf("index:%d,value:%d\r\n",key,value)
	}
}


//值类型，传递的是一个拷贝

func (tn TwoInts)Add()int  {
	return tn.A+tn.B
}

func (tn TwoInts)Reduce()int  {
	return tn.A-tn.B
}

func (tn TwoInts) Change()  {
	tn.A=333
	tn.B=444
}

//指针类型，传递的是一个引用

func (tn *TwoInts) Add2()int  {
	return tn.B+tn.A
}

func (tn *TwoInts) Reduce2()int {
	return tn.B-tn.A
}

func (tn *TwoInts) Change2()  {
	tn.A=555
	tn.B=666
}

func main() {

	in:=IntVector{1,2,3}
	in.For()
	in.Range()

	tn:=new(TwoInts)
	tn.A=100
	tn.B=10
	tn.Change()
	//fmt.Printf("%d",tn.Add())
	//fmt.Printf("%d",tn.Reduce())
	fmt.Printf("A:%d\r\n",tn.A)
	fmt.Printf("B:%d\r\n",tn.B)

	tn2:=&TwoInts{A:1,B:2}
	tn2.Change2()
	//fmt.Printf("%d",tn2.Add2())
	//fmt.Printf("%d",tn2.Reduce2())
	fmt.Printf("2A:%d\r\n",tn2.A)
	fmt.Printf("2B:%d\r\n",tn2.B)
}