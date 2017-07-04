package main

import (
    "fmt"
    "math"
    "strings"
    "log"
)


var atTime map[string]int32
var atValue map[string]string


var where = log.Print

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    where()
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    //v.Scale(10)
    //fmt.Println(v.Abs())
    fmt.Println(strings.FieldsFunc("xxx", split))
    test(1)
    SetAtTime("1")

    fmt.Print(atTime["1"])

    fmt.Println("hello go world!")
    setValue(&v)

    fmt.Printf("v.x: %v v.y: %v",v.X,v.Y)

    /*
0x00000000
0x00000001     a
0x00010001     00000001
    */

    a:=1
    a <<= 10

    fmt.Println("dsd:%v",a)

    var ptr *int

    ptr = &a

    fmt.Printf("ptr: %p",ptr)
    fmt.Printf("ptr value: %v",*ptr)

    for i:=0;i<3;i++{
        fmt.Printf("value is %d\n",i)
    }

 arr :=[]int{1,2,3}
    for pos,val :=range arr{
        if val==2{
            break
        }
        fmt.Printf("position is :%d value is: %d",pos,val)
    }

    strs :=[]string{"11","22","33"}
    testMoreParam(strs...)

    //Callback(1,2,Add)

Slice()

    n1:=number{1}
    n2:=nr{1}

    var c = number(n2)
    fmt.Println(n1,n2,c)

    var tt=&ba{1,qa{2}}
    fmt.Println(tt)

fmt.Println("test new commit for merge")
}

type number struct {
    i int   "this is a tag"
}

type qa struct {
    j int
}

type ba struct {

    k int
    qa
}

type nr number



func Slice(){
    s:=make([]int,2,3)
    fmt.Printf("len:%d",len(s))
    fmt.Printf("cap:%d",cap(s))

    for k,v:=range s{
        fmt.Printf("k:%d,v:%d\r\n",k,v)
    }

    s=append(s,4)
    for k,v:=range s{
        fmt.Printf("k:%d,v:%d\r\n",k,v)
    }
    fmt.Printf("len:%d",len(s))
    fmt.Printf("cap:%d",cap(s))
}


func Callback(a,b int,f func(x,y int)){
        f(a,b)
}

func Add(x,y int){
    fmt.Printf("x is %d, y is %d",x,y)
}


func testMoreParam(s ... string){
    for pos,val :=range s{
        fmt.Printf("pos: %d val:%s\n",pos,val)
    }
}

func setValue(v *Vertex){

    v.X=100
    v.Y=100
}

func SetAtTime(key string){
    if atTime==nil{
        atTime=make(map[string]int32)
    }
    atTime[key]=1;
}

func SetAtValue(key string,value string){
    if(atValue==nil){
        atValue=make(map[string]string)
    }
    atValue[key]=value
}


func split(s rune) bool {
    if s == 'n' {
        return true
    }
    return false
}

func test(i int){
    if(i==1) {
        return
    }

    if(i==2){
        return
    }

    fmt.Println(i)
}


