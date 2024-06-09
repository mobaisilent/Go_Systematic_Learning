package main

import "fmt"

func main() {
	//第一种 使用默认值
	var a int
	fmt.Printf("a = %d\n", a)

	//第二种
	var b int = 10
	fmt.Printf("b = %d\n", b)

	//第三种 省略后面的数据类型,自动匹配类型
	var c = 20
	fmt.Printf("c = %d\n", c)

	//第四种 省略var关键字
	d := 3.14
	fmt.Printf("d = %f\n", d)
}
