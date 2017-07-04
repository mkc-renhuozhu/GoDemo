package main

import (
	"fmt"
	"encoding/json"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}


func main() {



	var inp int = 5

	var inpt *int=&inp

	fmt.Println(inpt,*inpt)

	str:=Marshal()
	v,err:= UnMarshal(str)
	if err==nil{
		for _, val := range v.Addresses {
			fmt.Printf("Type:%s,City:%s,Country:%s\n",val.Type,val.City,val.Country)
		}
	}
}

func Marshal() string {
	ad1:=&Address{Type:"111",City:"222",Country:"333"}
	ad2:=&Address{Type:"444",City:"555",Country:"666"}
	vcard:=&VCard{FirstName:"zhu",LastName:"renhuo",Addresses:[]*Address{ad1,ad2},Remark:"xxx"}
	js, _ := json.Marshal(vcard)
	fmt.Println(string(js))
	return string(js)
}

func UnMarshal(str string)(*VCard,error) {
	var vcard VCard
	body := []byte(str)
	err:=json.Unmarshal(body,&vcard)
	if err!=nil{
		return nil,err
	}
	fmt.Println(&vcard)
	return &vcard,nil
}
