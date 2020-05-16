package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini文件解析器

// MySQLConfig ...
type MySQLConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Tese     bool   `ini:"test"`
}

// Config ...
type Config struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") // 新创建一个错误
		return
	}
	// 0.2 传进来的data参数必须是结构体类型的指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data parame should be a struct pointer")
		return
	}
	// 1. 读文件，得到字节类型的数据
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	// 2. 逐行读取数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首位空格
		line = strings.TrimSpace(line)
		// 空行跳过
		if len(line) == 0 {
			continue
		}
		// 2.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头的就表示节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[开头就是=分隔的键值对
			// 1. 以等号分隔这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax serror", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2. 根据structName去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			// 3. 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息是存储在类型信息中的
				fieldType = field
				if field.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = field.Name
					break
				}
			}
			// 4. 如果key=tag，给这个字段赋值
			// 4.1 根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 对其赋值
			fmt.Println(fieldName, fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("%d行,type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("%d行 type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("%d行 type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}

		}
	}
	return
}

func main() {
	var config Config
	err := loadIni("./config.ini", &config)
	if err != nil {
		fmt.Printf("load file failed!err:%v\n", err)
		return
	}
	// fmt.Println(mc.Address, mc.Port, mc.UserName, mc.Password)
	fmt.Printf("%#v\n", config)
}
