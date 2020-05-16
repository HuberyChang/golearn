package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{
		"name": "jaklds",
		"age":  90
	}`
	var stu student
	json.Unmarshal([]byte(str), &stu)
	fmt.Printf("%#v\n", stu)
}
