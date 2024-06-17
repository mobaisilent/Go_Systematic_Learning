## 继承

```go
type Human struct {
    name string
    sex  string
}

func (this *Human) Eat() {
    fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
    fmt.Println("Human.Walk()...")
}

type SuperMan struct {
    Human //SuperMan继承了Human的方法
    level int
}
```

父级代码和继承简单如上

```go
// 定义父类的方法Eat()
func (this *SuperMan) Eat() {
    fmt.Println("SuperMan.Eat()...")
}

// 定义额外方法
func (this *SuperMan) Fly() {
    fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
    fmt.Println("name = ", this.name)
    fmt.Println("sex = ", this.sex)
    fmt.Println("level = ", this.level)
}

// 定义一个子类来继承
// s := SuperMan{Human{"li4", "female"}, 88}
var s SuperMan
s.name = "li4"
s.sex = "male"
s.level = 88
```

结合上面的代码完整代码如下：
```go
package main

import "fmt"

type Human struct {
    name string
    sex  string
}

func (this *Human) Eat() {
    fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
    fmt.Println("Human.Walk()...")
}

type SuperMan struct {
    Human  
    level int
}

func (this *SuperMan) Eat() {
    fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() {
    fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
    fmt.Println("name =", this.name)
    fmt.Println("sex =", this.sex)
    fmt.Println("level =", this.level)
}

func main() {
    s := SuperMan{Human{"li4", "male"}, 88}
    s.Eat()
    s.Walk()
    s.Fly()
    s.Print()
}
// gpt4
```