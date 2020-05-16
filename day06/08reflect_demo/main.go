package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" JKFAJSD:"290-0F"`
	Score int    `json:"score" JKFAJSD:"88"`
}

func main() {
	stu := student{
		Name:  "ahlfa",
		Score: 90,
	}
	v := reflect.TypeOf(stu)
	fmt.Println(v.Name(), v.Kind()) // student struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("JKFAJSD"))
	}
	if scoreField, ok := v.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("JKFAJSD"))
	}
}
