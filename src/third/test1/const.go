package main

import "fmt"

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

func main() {
	fmt.Println(IgEggs)         // 1
	fmt.Println(IgChocolate)    // 2
	fmt.Println(IgNuts)         // 4
	fmt.Println(IgStrawberries) // 8
	fmt.Println(IgShellfish)    // 16
}

// 打印结果如右边所示：虽然其对应的二进制形式如定义那边的代码，但是显示的当然不是二进制啦 
