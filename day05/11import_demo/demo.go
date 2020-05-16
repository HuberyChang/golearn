package main

import (
	"fmt"
	calc "golearn/day05/10calc" // 文件夹名称不合规，起一个别名“calc”，后面是路径。最好是英文名文件夹。
)

func init() {
	fmt.Println("ahoapa904i9奥防静电")
}

func main() {
	ret := calc.Add(2, 3)
	fmt.Println(ret)
}
