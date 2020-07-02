package main

import "fmt"

/*
	结构体：由一系列具有相同或者不同类型的数据构成的数据集合
*/
//声明一个结构体
type People struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个people类型的变量ha
	//方法一
	var ha People
	fmt.Println(ha)
	//通过字段赋值
	ha.name = "chen"
	ha.age = 120
	ha.gender = "哈"
	ha.hobby = []string{"篮球", "离开家啊哈", "阿克纠纷"}
	// 访问变量ha的字段
	fmt.Println(ha)

	// 方法二
	ha1 := People{}
	ha1.name = "阿斯蒂芬"
	ha1.age = 20
	ha1.gender = "nan"
	ha1.hobby = []string{"asdh", "阿道夫"}
	fmt.Println(ha1)

	//方法三
	ha2 := People{"fhajs9q阿道夫", 20, "nan", []string{"l;kdf", "io13阿斯蒂芬"}}
	fmt.Println(ha2)
	fmt.Printf("%p,%T\n", &ha2, ha2) //0xc0001160c0,main.People

	ha3 := ha2
	fmt.Println(ha3)
	fmt.Printf("%p,%T\n", &ha3, ha3) //0xc000116180,main.People

	ha3.name = "12㐇"
	fmt.Println(ha3) //{12㐇         20 nan [l;kdf io13阿斯蒂芬]}
	fmt.Println(ha2) //{fhajs9q阿道夫 20 nan [l;kdf io13阿斯蒂芬]}

	//定义结构体指针
	var ha4 *People
	ha4 = &ha2
	fmt.Println(ha4)
	fmt.Printf("%p,%T\n", &ha4, ha4) //0xc0000ca020,*main.People
	fmt.Println(*ha4)                //{fhajs9q阿道夫 20 nan [l;kdf io13阿斯蒂芬]}
	ha4.name = "附件却颇"
	fmt.Println(*ha4) //{附件却颇 20 nan [l;kdf io13阿斯蒂芬]}
	fmt.Println(ha2)  //{附件却颇 20 nan [l;kdf io13阿斯蒂芬]}

	// 匿名结构体。多用于临时场景
	var s struct {
		name   string
		age    int
		gender string
		hobby  string
	}
	s.name = "hakj"
	s.age = 9
	s.gender = "难"
	fmt.Println(s)

	//或者
	s1 := struct {
		name string
		age  int
	}{
		"j安顺地区",
		20,
	}
	fmt.Println(s1)

}
