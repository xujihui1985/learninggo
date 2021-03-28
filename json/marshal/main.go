package main

import (
	"encoding/json"
	"fmt"
)

type Plant struct {
	MyColor string
}

func (p Plant) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["type"] = "plant"
	m["color"] = p.MyColor
	return json.Marshal(m)
}

type Animal struct {
	MyColor string
}

func main() {
	p := Plant{MyColor: "green"}
	a := Animal{MyColor: "red"}

	b, _ := json.Marshal(p)
	ba, _ := json.Marshal(a)

	fmt.Println(string(b))
	fmt.Println(string(ba))

}
