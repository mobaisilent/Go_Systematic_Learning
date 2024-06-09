package main

import "fmt"

// 定义函数及其两个返回值的类型
func swap(x, y string) (string, string) {
	return y, x
}

func fun1(x string, y int) (r1 int, r2 int) {
	r1 = len(x) // 字符串x的长度
	r2 = y      // 提供的整数值
	// 这里可以不初始化：当然不初始化就是0咯
	return
}

func main() {
	a, b := swap("Mahesh", "Kumar") // 用两个变量去接收两个参数
	fmt.Println(a, b)
	// 打印结果是：Kumar Mahesh

	r1, r2 := fun1("Mahesh", 100)
	fmt.Println(r1, r2)
	// 假设的打印结果，取决于fun1函数的实际逻辑。在这个例子中，如果使用"Mahesh"和100调用fun1，将打印：6 100
}
