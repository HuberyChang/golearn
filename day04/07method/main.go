package main

import "fmt"

// 方法

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

type person struct {
	age  int
	name string
}

func newPerson(age int, name string) person {
	return person{
		age:  age,
		name: name,
	}
}

// 方法是作用于 “特定类型” 的 “函数”
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:wangwangwang\n", d.name)
}

//值接收者：传递拷贝进去
func (p person) guonian() {
	p.age++
}

// 指针接收者：传递内存地址进去
// 使用场景：
// 1.需要修改接收者中的值的时候
// 2.接收者是拷贝代价比较大的大对象，内存地址永远都是一个uint64的值
// 3.保证一致性，如果某个方法使用了指针接收者，那么其他方法也应该使用指针接收者
func (p *person) zhenguonian() {
	p.age++
}

func main() {
	d1 := newDog("hfak")
	d1.wang()
	p1 := newPerson(14, "hfakj")
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zhenguonian() //语法糖
	// (&p1).zhenguonian()
	fmt.Println(p1.age)
}
