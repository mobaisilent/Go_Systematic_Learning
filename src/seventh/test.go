package main

import "fmt"

func main() {
	//====map的声明方式====

	//第一种声明方式
	var myMap1 map[string]string
	//声明map
	myMap1 = make(map[string]string) // 先不给map分配空间
	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "c++"

	fmt.Println("myMap1 =", myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)

	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"

	fmt.Println("myMap2 =", myMap2)

	//第三种声明方式
	myMap3 := map[string]string{
		"one":   "pho",
		"two":   "c++",
		"three": "python",
	}
	// 也就是直接写出键值对的形式
	fmt.Println("myMap3 =", myMap3)
}
