package main

import "fmt"

func printArray(myArray [4]int) {
	// 值拷贝 ：数组长度是限定的
	for index, value := range myArray {
		fmt.Println("index =", index, ", value =", value)
	}

	myArray[0] = 11
	// 实际可见myArray[0]的值并不会有任何改变

}

func main() {
	// 固定长度的数组
	var myArray1 [10]int

	// 使用简短声明方式声明另一个数组
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	// for遍历打印
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	// 遍历range遍历打印
	for index, value := range myArray2 {
		fmt.Println("index =", index, ", value =", value)
	}

	fmt.Print("the type of myArray3 is = %T\n", myArray3)
}
