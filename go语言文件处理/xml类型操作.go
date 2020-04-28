package main

import (
	"encoding/xml"
	"fmt"
)

type xmlType struct {
	Name string `xml:"name"`
	Age  int
	Url  string
}

func main() {
	xml1 := []xmlType{{"曾浩", 20, "http://www.michael.com"}}
	xml_data, err := xml.MarshalIndent(xml1, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(xml_data))
	}
	var xml2 = new([]xmlType)
	err1 := xml.Unmarshal(xml_data, &xml2)
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println(xml2)
	}

}
