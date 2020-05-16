package main

import "fmt"

//分金币

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	//1.遍历拿到每个人的名字
	for _, name := range users {
		//2.拿到一个人名根据分金币的规则分金币
		//2.1每个人分的金币数应该保存到map中
		//2.2记录剩余金币数
		for _, c := range name {
			switch c {
			case 'e', 'E':
				//满足这个条件分1枚金币
				distribution[name]++
				coins--
			case 'i', 'I':
				//满足这个条件分2枚金币
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				//满足这个条件分3枚金币
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				//满足这个条件分4枚金币
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	//3.整个第二步执行完就可以得到每个人最终分配数和剩余数
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下金币：", left)
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}
