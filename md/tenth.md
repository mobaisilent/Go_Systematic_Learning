## Interface

> go语言的多态性

```go
package main

import "fmt"

type AnimalIF interface {
    Sleep()
    GetColor() string // 获取动物颜色
    GetType() string  // 获取动物类型
}

// Cat结构体实现AnimalIF接口
type Cat struct {
    color string // 猫的颜色
}

func (this *Cat) Sleep() {
    fmt.Println("Cat is Sleeping")
}

func (this *Cat) GetColor() string {
    return this.color
}

func (this *Cat) GetType() string {
    return "Cat"
}

// Dog结构体实现AnimalIF接口
type Dog struct {
    color string // 狗的颜色
}

func (this *Dog) Sleep() {
    fmt.Println("Dog is Sleeping")
}

func (this *Dog) GetColor() string {
    return this.color
}

func (this *Dog) GetType() string {
    return "Dog"
}

// main函数，示例化Cat和Dog，展示多态特性
func main() {
    var animal AnimalIF

    animal = &Cat{"Green"}
    animal.Sleep() // Cat的Sleep()方法
    fmt.Println("This animal is a", animal.GetType(), "with", animal.GetColor(), "color")

    animal = &Dog{"Yellow"}
    animal.Sleep() // Dog的Sleep()方法
    fmt.Println("This animal is a", animal.GetType(), "with", animal.GetColor(), "color")
}
```

> 关于animal = &Cat{"Green"}的讲解
>
> 这一行代码 `animal = &Cat{"Green"}` 在 Go 语言中执行了几个操作： 
>
> 1. **实例化Cat结构体** 
> 2. **取地址操作** 
> 3.  **赋值给接口变：** 