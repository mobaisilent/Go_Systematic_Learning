# GO指南笔记汇总

> 转自：
>
> https://tour.go-zh.org/welcome/1
>
> 显然该篇目文档更为详细，算是对8小时go的补充和复习
>
> 于6.15完成它吧~ 自己看的时候也就作为复习使用稍微仔细点看吧

## 包,变量与函数

### 包

**每个 Go 程序都由包构成**。

程序从 `main` 包开始运行。

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("我最喜欢的数字是 ", rand.Intn(10))
}

```

### 导入

```
import "fmt"
import "math"
```

示例代码：

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("现在你有了 %g 个问题。\n", math.Sqrt(7))
}

```

### 导出名

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，`Pizza` 就是个已导出名，`Pi` 也同样，它导出自 `math` 包。

`pizza` 和 `pi` 并未以大写字母开头，所以它们是未导出的。

在导入一个包时，你只能引用其中已导出的名字。 任何「未导出」的名字在该包外均无法访问。

示例代码：

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
}
// 这段代码是错误代码：将pi改为大写的P就行了 Pi
```

### 函数

函数可接受零个或多个参数。

在本例中，`add` 接受两个 `int` 类型的参数。

示例代码：

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```

### 函数（续）

> 将 多个同类型的函数参数简化书写

```go
package main

import "fmt"

// 简化那就见这里
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

```

### 多返回值

函数可以返回任意数量的返回值。

`swap` 函数返回了两个字符串。

> 可以用多个参数去接收多返回值：如果是不需要使用的参数 那么就用_ 接收

示例代码：
```go
package main

import "fmt"

// 返回的两个参数都为string  下面的那个示例才是将返回的参数命名化
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
	
```

### **带名字的返回值**

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的命名应当能反应其含义，它可以作为文档使用。

没有参数的 `return` 语句会直接返回已命名的返回值，也就是「裸」返回值。

裸返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

示例代码：

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

```

### 变量

`var` 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。

如例中所示，`var` 语句可以出现在包或函数的层级。

示例代码：

```go
package main

import "fmt"

// 简单变量的声明 其类型都是bool  bool变量自动初始化为false
var c, python, java bool

func main() {
	var i int  // int变量自动初始化为0
	fmt.Println(i, c, python, java)
}

```

### 变量的初始化

```go
package main

import "fmt"

var i, j int = 1, 2
// 需要辨析下初始化变量的位置 像c++ int a = 0  :: 注意语法之间的不同

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

```

### 短变量声明

在函数中，短赋值语句 `:=` 可在隐式确定类型的 `var` 声明中使用。

函数外的每个语句都 **必须** 以关键字开始（`var`、`func` 等），因此 `:=` 结构不能在函数外使用。

示例代码：
```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
    k := 3 // 声明短变量  隐式确定类型的 var 声明中使用  := 不能在函数外中使用：：注意全局变量不可使用
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

```

### 基本类型

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 表示一个 Unicode 码位

float32 float64

complex64 complex128
```

### 0值

没有明确初始化的变量声明会被赋予对应类型的 **零值**。

零值是：

- 数值类型为 `0`，
- 布尔类型为 `false`，
- 字符串为 `""`（空字符串）。

示例代码：
```go
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

```

运行结果是：

````go
0 0 false ""

````

### 类型转换

一些数值类型的转换：

```
var i int = 42  // 初始化
var f float64 = float64(i)  // 转为float64
var u uint = uint(f) // 转为uint
```

或者，更加简短的形式：（推荐使用）

```
i := 42
f := float64(i)
u := uint(f)
```

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
    
    // 注意这里都是显示转换：：如果不进行显示转换编译器会报错
    
	fmt.Println(x, y, z)
	fmt.Println(f,z)
}

```

> 注意：
> var z = uint(f)  // 这样也是正确的  ；； 这个也是显示转换：笔记是将 f转为 uint然后赋值给 z
>
> var z uint = f  // 这里是错误的 ：： go不允许隐式转换的出现

### 类型推断

在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法 `var =` 表达式语法），变量的类型会通过右值推断出来。

当声明的右值确定了类型时，新变量的类型与其相同：

```
var i int
j := i // j 也是一个 int
// 符合类型推断：：这句其实就是 类型的传递
```

不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 `int`、`float64` 或 `complex128` 了，这取决于常量的精度：

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
// 由右边的精度直接推断数据类型
```

试着修改示例代码中 `v` 的初始值，并观察它是如何影响类型的。

示例代码：
```go
package main

import "fmt"

func main() {
	v := 42 // 修改这里看看！
	fmt.Printf("v is of type %T\n", v)
}

```

补充go语言小数的精度：

- float32  -> 6，7位
- float64  -> 15，17位

### 常量

常量的声明与变量类似，只不过使用 `const` 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 `:=` 语法声明。

示例代码：
```go
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

```

### 数值常量

数值常量是高精度的 **值**。

一个未指定类型的常量由上下文来决定其类型。

再试着一下输出 `needInt(Big)` 吧。

（`int` 类型可以存储最大 64 位的整数，根据平台不同有时会更小。）

示例代码：
```go
package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
    // 将Big转为float64
}

```

运行结果：
```go
21
0.2
1.2676506002282295e+29
// 科学计数法“后面跟着 29个 0
```

## 流程控制语句

### for 循环

Go 只有一种循环结构：`for` 循环。

基本的 `for` 循环由三部分组成，它们用分号隔开：

- 初始化语句：在第一次迭代前执行
- 条件表达式：在每次迭代前求值
- 后置语句：在每次迭代的结尾执行

初始化语句通常为一句短变量声明，该变量声明仅在 `for` 语句的作用域中可见。

一旦条件表达式求值为 `false`，循环迭代就会终止。

**注意**：和 C、Java、JavaScript 之类的语言不同，Go 的 `for` 语句后面的三个构成部分外没有小括号， 大括号 `{ }` 则是必须的。

示例代码：

```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // 注意没有小括号 但是 必有大括号  但是依旧是 起点 终点 处理 基本三个件
		sum += i
	}
	fmt.Println(sum)
}
// 注意 不需要指明i的类型 也就是 类型推断功能
```

for循环的变式

```go

import "fmt"

func main() {
    sum := 1  // 将 变量初始化 之后移动出去：c++ 等 也支持这样 ; sum < 1000; 但是 : 记得写
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

```

### for 是 Go 中的「while」

c的while在go中叫做for

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
// 感觉就是上面之前那个示例代码中将 : 去除的样子
```

### 无限循环

```go
package main

func main() {
	for {
	}
}

```

### if判断

示例代码：

```go
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
// 与其他类编程语言也没有太多区别
```

### if 和简短语句

和其他类编程代码依旧没啥区别的那种

示例代码：

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
    // if的后面就是 一个 执行语句 然后判断 然后执行if内的语句
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

### if和else

和其他语言也没啥区别

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

```

### 练习

> for循环的简单使用和练习

```go
package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2 * z)
        fmt.Println(z)
    }
    return z
}

func main() {
    for i := 1; i <= 10; i++ {
        sqrt := Sqrt(float64(i))
        realSqrt := math.Sqrt(float64(i))
        fmt.Printf("Sqrt(%v): %v, math.Sqrt(%v): %v, difference: %v\n", i, sqrt, i, realSqrt, math.Abs(sqrt-realSqrt))
    }
}
```

### switch分支语句

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go 运行的系统环境：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
// os := runtime.GOOS; 先执行了一个简单的赋值语句 然后后面才是根据os的值进行判断 调用case中的判断
```

### switch中的求值顺序

顺序：：自上而下

示例代码：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	switch time.M {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}
}
//  根据差值进行case计算 从插值是0天 开始 到插值是 3天以及以上
```

### 无条件 switch

无条件的 `switch` 同 `switch true` 一样。

这种形式能将一长串 `if-then-else` 写得更加清晰。

示例代码：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:  // 每个case后面接的就是 相当于一个if语句
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}

```

> 作用：将一长串 的 if else if  等简化   虽然实际效果大差不差的

### defer 推迟

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

> 将defer后面的语句压入栈

示例代码：
```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
/*
输出结果是
hello
world
也就是main块执行务必的时候将栈中的world那条语句给放出来了
*/
```

### defer 栈

推迟调用的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的调用会按照后进先出的顺序调用。

示例代码：

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
// 结果显然是先输出done 然后逆序输出i序列
```

> 补充 下go语言的函数命名规则
>
> 函数名 ->  参数列表 -> 函数返回值 -> 函数体
>
> 例如：func add(x int, y int) int {}
>
> ​			 func swap(x, y string) (string, string) {}
>
> 有些带有应引用数据类型的那么注意是这样
>
> 指针引用数据类型 -> 函数名 ->  参数列表 -> 函数返回值 -> 函数体
>
> 例如：func (this *Server)add(x int, y int) int {}

## 更多数据类型（结构体，切片，映射...)

### 指针

Go 拥有指针。指针保存了值的内存地址。

类型 `*T` 是指向 `T` 类型值的指针，其零值为 `nil`。

```
var p *int
```

`&` 操作符会生成一个指向其操作数的指针。

```
i := 42
p = &i
```

`*` 操作符表示指针指向的底层值。

```
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

这也就是通常所说的「解引用」或「间接引用」。

与 C 不同，Go 没有指针运算。

> 指针的定义倒是大差不差 但是 go语言没有 指针的运算

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
    fmt.Println(i)  // 查看 i 的值 :: 可以发现i的值发生了变化

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值  ；； 还是可见j的值发生了变化

```

### 结构体

一个 结构体（`struct`）就是一组 字段（field）。

示例代码：

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
// 定义结构体 注意前面是有Vertex的

func main() {
	fmt.Println(Vertex{1, 2})
}

```

结构体数组的构建：

```go
package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    vertices := []Vertex{
        {1, 2},
        {3, 4},
        {5, 6},
        // 你可以添加更多的Vertex结构体
    }

    for _, v := range vertices {
        fmt.Println(v)
    }
    // 注意 range的使用 返回两个参数 第一个是下标 第二个是每一项
}
```

输出结果是： 直接printf结构体就可以输出结果来
```go
{1 2}
{3 4}
{5 6}
```



结构体字段的访问：

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
// 也就是通过点号访问
```

结构体指针

示例代码：
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
// 也就是普通指针的使用而已：不过多赘述
```

**结构体字面量**

特殊的前缀 `&` 返回一个指向结构体的指针。

示例代码：

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
    // 注意这种批量定义变量的形式
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予零值
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

```

### 数组

```
var a [10]int  // 注意数组的创建方式
```

示例代码：
```go
package main

import "fmt"

func main() {
	var a [2]string  // 字符串数组  // 注意数组的创建语法 也就是普通变量那里添加下[]表示数量即可
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

```

### 切片

> 相比于普通数组来说就是动态长度可变的数组了

每个数组的大小都是固定的。而切片则为数组元素提供了动态大小的、灵活的视角。 **在实践中，切片比数组更常用。**

类型 `[]T` 表示一个元素类型为 `T` 的切片。.

切片通过两个下标来界定，一个下界和一个上界，二者以冒号分隔：

```
a[low : high]
```

它会选出一个半闭半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```
a[1:4]
```

示例代码：

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}  // 创建数组  // 这里再次强调记忆一下 
    // 普通数据的创建 蔽日 test := int 5 也就是在数据类型前面用方括号表示有多少个元素即可

	var s []int = primes[1:4]  // 创建切片数组 从primes数组的第二个元素到第四个元素
	fmt.Println(s)
}

```

> 关于go语言中切片的解释：
>
> 尤其是这句：var s []int = primes[1:4]
>
> 创建切片数组  从上一个数组中取数据 从 下标1到下标4  尝试了可以从下标0开始取得数据

### 切片类似数组的引用

**切片就像数组的引用 切片并不存储任何数据，它只是描述了底层数组中的一段。**

**更改切片的元素会修改其底层数组中对应的元素。**

和它共享底层数组的切片都会观测到这些修改。

示例代码：
```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)  // 创建和打印数组

	a := names[0:2]
	b := names[1:3]  // 取得切片
	fmt.Println(a, b)  // 打印切片

	b[0] = "XXX"   // 修改切片
	fmt.Println(a, b)
	fmt.Println(names)  // 把原数组也修改了
}

```

### 切片字面量

> 关于go语言中的字面量的概念解释：
>
> 

切片字面量类似于没有长度的数组字面量。

这是一个数组字面量：

```
[3]bool{true, true, false}
```

下面这样则会创建一个和上面相同的数组，然后再构建一个引用了它的切片：

```
[]bool{true, true, false}
```

示例代码：
```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

```

### 切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 0，上界则是该切片的长度。

对于数组

```
var a [10]int  // 创建长度固定的数组
```

来说，以下切片表达式和它是等价的：

```
a[0:10]
a[:10]
a[0:]
a[:]  // 使用切片的默认行为来表示原来的数组 默认就是从0到最后一个
```

示例代码：

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)  // 1到4

	s = s[:2]
	fmt.Println(s)  // 0到2

	s = s[1:]
	fmt.Println(s)  // 1到最后一个
}

```

### 切片的长度与容量

切片拥有 **长度** 和 **容量**。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

你可以通过重新切片来扩展一个切片，给它提供足够的容量。 试着修改示例程序中的切片操作，向外扩展它的长度，看看会发生什么。

> 容量是可以扩充的  每次空间不够了 扩充到原来的2倍

示例代码：
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}  // 创建切片数组
	printSlice(s)  // 打印切片数组

	// 截取切片取其0长度开始  所以 len为0  但是cap为 6
	s = s[:0]  // 切割显示
	printSlice(s)  // len=0 但是cap=6

	// 扩展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)  // 注意这里才是舍去其前面的值 之前那里不是舍弃  这里的len=2 但是cap=4
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

```

### nil 切片

切片的零值是 `nil`。

nil 切片的长度和容量为 0 且没有底层数组。

示例代码：
```go
package main

import "fmt"

func main() {
	var s []int  // 空切片 [] 0 0
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")  // nil!
	}
}

```

### 用 make 创建切片

切片可以用内置函数 `make` 来创建，这也是你创建动态数组的方式。

`make` 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

```
a := make([]int, 5)  // len(a)=5  // int类型5个空间
```

要指定它的容量，需向 `make` 传入第三个参数：

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5   // int类型 len为0 cap为6

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

示例代码：

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]  // 0到2
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

```

### 切片的切片

切片可以包含任何类型，当然也包括其他切片。

示例代码：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字棋（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
    // 创建切片二维数组

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
    // 填充数据

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

```

输出结果是：
```go
X _ X
O _ X
_ _ O

```

### 向切片追加元素

使用append函数

示例代码：

```go
package main

import "fmt"

func main() {
	var s []int  // 创建切片数组
	printSlice(s)

	// 可在空切片上追加
	s = append(s, 0)
	printSlice(s)  // 追加数据0

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

```

### range 遍历

示例代码：
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}  // 创建切片数组

func main() {
	for i, v := range pow {  // 使用range进行遍历 前为下标 后为值
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

```

可以将下标或值赋予 `_` 来忽略它。（推荐用法）

```
for i, _ := range pow
for _, value := range pow
```

若你只需要索引，忽略第二个变量即可。

```
for i := range pow
```

示例代码：
```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

```

### map 映射

`map` 映射将键映射到值。

映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。

`make` 函数会返回给定类型的映射，并将其初始化备用。

示例代码：

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
// 创建结构体用来别是键值对每一项

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)  // 创建map映射 键是string  值是两个 float64	
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
    // 创建一个键是 “Bell Labs" 值是 上面两个float64
	fmt.Println(m["Bell Labs"])
}

```

代码解释如下：

- `type Vertex struct { Lat, Long float64 }`：这行代码定义了一个名为 `Vertex` 的结构体，它有两个字段：`Lat` 和 `Long`，它们的类型都是 `float64`。
- `var m map[string]Vertex`：这行代码声明了一个名为 `m` 的变量，它的类型是一个映射，该映射的键是字符串，值是 `Vertex` 结构体。
- `m = make(map[string]Vertex)`：这行代码使用 `make` 函数创建了一个新的映射，并将它赋值给 `m`。在 Go 语言中，映射必须使用 `make` 函数创建后才能使用。
- `m["Bell Labs"] = Vertex{ 40.68433, -74.39967, }`：这行代码创建了一个新的 `Vertex` 结构体，并将它添加到映射 `m` 中，键是 `"Bell Labs"`。
- `fmt.Println(m["Bell Labs"])`：这行代码打印出映射 `m` 中键为 `"Bell Labs"` 的值。

### 映射字面量

映射的字面量和结构体类似，只不过必须有键名。

示例代码：

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
// 创建结构体

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
// 创建m键值对

func main() {
	fmt.Println(m)
}
// 稍微区分一下
```

若顶层类型只是一个类型名，那么你可以在字面量的元素中省略它。

代码示例：

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  // 在字面量的元素中省略
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

```

> 补充：什么是go语言的字面量  --> 简言之 每种数据类型后买你那堆具体数据就是 字面量
>
> 在 Go 语言中，字面量（Literal）是一种表示固定值的表示法。字面量可以是整数、浮点数、布尔值、字符串、数组、切片、结构体、映射等类型的值。
>
> - 整数字面量：`10`、`-3`、`0x7f` 等。
> - 浮点数字面量：`3.14`、`-0.01`、`1.0e7` 等。
> - 布尔值字面量：`true` 和 `false`。
> - 字符串字面量：`"hello"`、`"world"` 等。
> - 数组字面量：`[3]int{1, 2, 3}`、`[...]string{"apple", "banana", "cherry"}` 等。
> - 切片字面量：`[]int{1, 2, 3}`、`[]string{"apple", "banana", "cherry"}` 等。
> - 结构体字面量：`Vertex{1.0, 2.0}`、`Person{"Alice", 30}` 等。
> - 映射字面量：`map[string]int{"apple": 1, "banana": 2}`、`map[int]Vertex{1: {1.0, 2.0}, 2: {3.0, 4.0}}` 等。

### 修改映射

> 直接根据键修映射即可

在映射 `m` 中插入或修改元素：

```
m[key] = elem
```

获取元素：

```
elem = m[key]
```

删除元素：

```
delete(m, key)
```

通过双赋值检测某个键是否存在：

```
elem, ok = m[key]
```

若 `key` 在 `m` 中，`ok` 为 `true` ；否则，`ok` 为 `false`。

若 `key` 不在映射中，则 `elem` 是该映射元素类型的零值。

**注**：若 `elem` 或 `ok` 还未声明，你可以使用短变量声明：

```
elem, ok := m[key]
```

示例代码：
```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["答案"] = 42
	fmt.Println("值：", m["答案"])

	m["答案"] = 48
	fmt.Println("值：", m["答案"])

	delete(m, "答案")
	fmt.Println("值：", m["答案"])

	v, ok := m["答案"]
	fmt.Println("值：", v, "是否存在？", ok)
    // 根据映射关系直接修改字面量
}

```

### 函数值 -> 高阶函数

函数也是值。它们可以像其他值一样传递。

函数值可以用作函数的参数或返回值。

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {  // 仔细分析结构知道 函数的蚕食是一个函数值
	return fn(3, 4)  // 也就是将3 和 4 向上传递给fn
}
// 看着像是套娃的感觉
// 稍微复杂点的函数 调用 但是实际看好像效果还不错

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
        // 13的平凡是5的平方和12的平方的和
	}
    // 创建函数值
    
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

```

运行结果：
```go
13
5
81
```

关于那三行函数的相关调用信息信息i:

```go
fmt.Println(hypot(5, 12))：这行代码首先调用了 hypot 函数，参数是 5 和 12，然后将 hypot 函数的返回值作为参数传递给 fmt.Println 函数，最后打印出 hypot 函数的返回值。

fmt.Println(compute(hypot))：这行代码首先调用了 compute 函数，参数是 hypot 函数，然后将 compute 函数的返回值作为参数传递给 fmt.Println 函数，最后打印出 compute 函数的返回值。在 compute 函数内部，hypot 函数被调用，参数是 3 和 4。
3的平方加4的平凡开放的结果是5

fmt.Println(compute(math.Pow))：这行代码首先调用了 compute 函数，参数是 math.Pow 函数，然后将 compute 函数的返回值作为参数传递给 fmt.Println 函数，最后打印出 compute 函数的返回值。在 compute 函数内部，math.Pow 函数被调用，参数是 3 和 4
81就是3的4次方
```



> 相关解释如下：
>
> 这个函数名为 `compute`，它接受一个函数作为参数，并返回一个 `float64` 类型的值。
>
> 参数 `fn` 是一个函数，这个函数接受两个 `float64` 类型的参数，并返回一个 `float64` 类型的值。在 `compute` 函数内部，`fn` 函数被调用，并传入了两个参数 `3` 和 `4`，然后 `compute` 函数返回了 `fn` 函数的返回值。
>
> 这是一个高阶函数的例子，高阶函数是一种可以接受其他函数作为参数或者将其他函数作为结果返回的函数。在这个例子中，`compute` 函数接受一个函数作为参数，然后调用这个函数，并返回其结果。这种模式在函数式编程中非常常见，可以用来创建各种复杂的行为。

### 函数闭包

> 闭包？？？

Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。 该函数可以访问并赋予其引用的变量值，换句话说，该函数被“绑定”到了这些变量。

例如，函数 `adder` 返回一个闭包。每个闭包都被绑定在其各自的 `sum` 变量上。

示例代码：
```go
package main

import "fmt"

func adder() func(int) int {  // 注意函数的构造类型：参数为空，返回值是一个函数2 函数2的参数为int 返回值也是int
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
// 创建闭包函数 adder

func main() {
	pos, neg := adder(), adder()  // 先定义函数
	for i := 0; i < 10; i++ {
		fmt.Println(
            pos(i),  //  展开来看就是 add((i))
			neg(-2*i),  // 往函数里面传递参数  
		)
	}
}

```

> 有一种闭包的美感：但是感觉实际上使用率应该不高

斐波那契闭包：

```go
package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b  // 斐波那契数列的处理方式  // 然后每次将结果传递给下一层  有一总递归函数的美感
        return a
    }
}

func main() {
    f := fibonacci() // 创建函数
    for i := 0; i < 10; i++ {
        fmt.Println(f())  // 闭包调用函数  // 因为重复执行了里面的匿名函数 导致每次都可以继承之前的结果
    }
    // 尝试改为2开始的for循环 但是依旧从1 1 开始显示 然后显示8个数字
}
// 显示结果就是从1 1 开始 然后显示多少个 与 for循环的范无关：  范围内为越大 显示的内容就越多
```

> 闭包概念是在之前的编程语言中没了解过的知识点 可以稍微多看看 了解了解

关于该斐波那契闭包的解释如下：

1. `func fibonacci() func() int {...}`：这是一个名为 `fibonacci` 的函数，它没有参数，返回值是一个函数，这个返回的函数没有参数，返回值是 `int` 类型。在 `fibonacci` 函数内部，定义了两个 `int` 类型的变量 `a` 和 `b`，并分别初始化为 `0` 和 `1`。然后返回了一个匿名函数，这个匿名函数在被调用时，会更新 `a` 和 `b` 的值，然后返回新的 `a` 的值。
2. `f := fibonacci()`：这行代码在 `main` 函数中调用了 `fibonacci` 函数，并将返回的函数赋值给了变量 `f`。
3. `for i := 0; i < 10; i++ {...}`：这是一个循环，它会执行10次。在每次循环中，都会执行其中的代码。
4. `fmt.Println(f())`：这行代码在循环中被执行，它调用了 `f` 函数，并将 `f` 函数的返回值打印出来。

## 方法

Go 没有类。不过你可以为类型定义方法。

方法就是一类带特殊的 **接收者** 参数的函数。  就是我上面提到的 func (this *server) ms (int) int {} 这种特殊的构造

方法接收者在它自己的参数列表内，位于 `func` 关键字和方法名之间。

在此例中，`Abs` 方法拥有一个名字为 `v`，类型为 `Vertex` 的接收者。

什么是方法示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// 创建go结构体

// 带一个Vertex参数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

```

### 方法即函数

记住：方法只是个带接收者参数的函数。

现在这个 `Abs` 的写法就是个正常的函数，功能并没有什么变化。

也就是函数调用用到了一个参数Vertex中的值而已：注意上面是浅拷贝

### 方法（续）

你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 `Abs` 方法的数值类型 `MyFloat`。

你只能为在同一个包中定义的接收者类型声明方法，而不能为其它别的包中定义的类型 （包括 `int` 之类的内置类型）声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内。）

示例代码如下：
```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64  //数据类型自定义 名称

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
// 用f.Abs() 实现给Abs带参数
```

### 指针类型的接收者

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// 创建结构体

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// 创建带参方法

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// 创建带引用值的方法 ：： 使用了引用（指针）那么就可以改变原来的数值了

func main() {
	v := Vertex{3, 4}  // 定义一个结构体变量 v
	v.Scale(10)  // 对v调用Scale改变v的值
	fmt.Println(v.Abs())
}

```

### 指针与函数

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// 有点像像是指针的调用  ：：可以实现改变v的值 ：：感觉这里和上卖弄方法大差不大的

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

```

### 方法与指针重定向

比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

```
var v Vertex
ScaleFunc(v, 5)  // 编译错误！
ScaleFunc(&v, 5) // OK
```

而接收者为指针的的方法被调用时，接收者既能是值又能是指针：

```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
// 可见调用方法更适配些
```

对于语句 `v.Scale(5)` 来说，即便 `v` 是一个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 `Scale` 方法有一个指针接收者，为方便起见，Go 会将语句 `v.Scale(5)` 解释为 `(&v).Scale(5)`。

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

```



### 方法与指针重定向（续）



示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// 创建结构体

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// 创建浅拷贝方法

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// 创建浅拷贝函数

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
    // 调用浅拷贝方法和浅拷贝函数

	p := &Vertex{4, 3}  // 指针结构体
	fmt.Println(p.Abs())  
	fmt.Println(AbsFunc(*p))
    // 这两种调用也是可以的  非指针的函数接收指针的变量，但是指针的函数不接受非指针的变量
}

```

### 选择值或指针作为接收者

使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样会更加高效。

在本例中，`Scale` 和 `Abs` 接收者的类型为 `*Vertex`，即便 `Abs` 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。 （我们会在接下来几页中明白为什么。）

示例代码：

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
// 深拷贝方法

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// 深拷贝方法

func main() {
	v := &Vertex{3, 4}  // 结构体指针
	fmt.Printf("缩放前：%+v，绝对值：%v\n", v, v.Abs())
	v.Scale(5)  // 调用函数
	fmt.Printf("缩放后：%+v，绝对值：%v\n", v, v.Abs())
}

```

> 并不应该混用

### 接口

**接口类型** 的定义为一组方法签名。

接口类型的变量可以持有任何实现了这些方法的值。

**注意:** 示例代码的第 22 行存在一个错误。由于 `Abs` 方法只为 `*Vertex` （指针类型）定义，因此 `Vertex`（值类型）并未实现 `Abser`。

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
// 创建接口 ：但是并未实现

func main() {
	var a Abser // 创建接口a
	f := MyFloat(-math.Sqrt2)  // 这里就是对2开方而已
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64  // 自定义数据类型的名称
 
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

关于go语言的接口以及如何使用

```
在 Go 语言中，接口是一种类型，它定义了一组方法，但是这些方法的具体实现是由其他类型（实现类型）来完成的。接口类型的变量可以保存任何实现了这些方法的值。

例如定义一个接口：
type Shape interface {
    Area() float64
}

这个示例和上面的基本一直

然后，我们可以定义 Rectangle 和 Circle 类型，它们都实现了 Shape 接口的 Area 方法：
type Rectangle struct {
    Width, Height float64
}
//定义结构体用于实现接口功能

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
// 实现机构体功能

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
// 实现机构体功能  几乎都是采用方法实现的:非指针：浅引用方法
```

### 接口与隐式实现

类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。

隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

示例代码：
```go
package main

import "fmt"

type I interface {
	M()
}
// 创建接口

type T struct {
	S string
}
// 创建T结构体

// 此方法表示类型 T 实现了接口 I，不过我们并不需要显式声明这一点。
func (t T) M() {
	fmt.Println(t.S)
}
// 没有显示的声明实现了接口

func main() {
	var i I = T{"hello"}
	i.M()
}

```

### 接口值

**接口也是值。它们可以像其它值一样传递。**

**接口值可以用作函数的参数或返回值。**

在内部，接口值可以看做包含值和具体类型的元组：

```
(value, type)
```

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

示例代码：
```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}
// 创建接口

type T struct {
	S string
}
// 用于实现接口的结构体

func (t *T) M() {
	fmt.Println(t.S)
}
// 实现接口：用指针方法实现的

type F float64 // 自定义数据类型名称

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I  // 创建I接口赋值给i

	i = &T{"Hello"}  // 然后直接调用接口
	describe(i)
	i.M()  // 调用I接口中的函数

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

> 关于输出那里的%v和%T的区别如下：
>
> 1. `%v`：这是最基本的格式化动词，它可以打印任何类型的值。对于复合类型的值，`%v` 会以 Go 语言的语法格式输出。
> 2. `%T`：这个动词用于打印一个值的类型。

> 关于上面代码的简要解释：
>
> 

### 底层值为 nil 的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 `M` 方法）。

**注意:** 保存了 nil 具体值的接口其自身并不为 nil。

示例代码：

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {  // 带指针的方法实现接口
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
// 创建接口 结构体 以及实现接口函数M

func main() {
	var i I  // 创建接口

	var t *T  // 创建指针结构体T赋值给t
	i = t  // 将t赋值给接口
	describe(i)
	i.M()  // 调用M

	i = &T{"hello"}  //为接口中的T结构体进行赋值
	describe(i)
	i.M()  // 然后调用M方法
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```



> 可以看看 describe的输出结果：上面代码的输出结果是：  可见describe可以将传递的接口的详细信息显示出来
>
> ```
> (<nil>, *main.T)  // 赋值先
> <nil>  //以为是nil自然输出nil
> (&{hello}, *main.T)  // 赋值后
> hello  // 赋值了hello那么输出hello
> ```

> 这段代码还是很有意思的~：：尤其是赋值那段：尝试下非指针赋值结果会是什么？ -> 将& 号去掉之后就报错啦

### nil接口值

nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 **具体** 方法的类型。

示例代码：

```go
package main

import "fmt"

type I interface {
	M()
}
// 创建接口

func main() {
	var i I  // 创建接口赋值给i
	describe(i)  // 接口信息打印调试
	i.M()  // 调用接口的函数实现
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)  // 值为空 接口也为空
}

```

> `fmt.Printf("(%v, %T)\n", i, i)` 这行代码会打印出两个值，第一个是变量 `i` 的值，第二个是变量 `i` 的类型。
>
> 打印结果是： 运行时错误
>
> ```
> (<nil>, <nil>)
> panic: runtime error: invalid memory address or nil pointer dereference
> [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x481b59]
> ```

### 空接口

指定了零个方法的接口值被称为 *空接口：*

```
interface{}
```

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。例如，`fmt.Print` 可接受类型为 `interface{}` 的任意数量的参数。

示例代码：

```go
package main

import "fmt"

func main() {
	var i interface{}  // 创建空接口
	describe(i)  // 描述空接口

	i = 42
	describe(i)  // 对上面 那个空接口进行赋值

	i = "hello"  // 对空接口赋值
	describe(i)
}

func describe(i interface{}) { 
	fmt.Printf("(%v, %T)\n", i, i)
}

```

运行结果如下： 空接口后面自动补全数据类型

```
(<nil>, <nil>)
(42, int)
(hello, string)
```

### 类型断言

**类型断言** 提供了访问接口值底层具体值的方式。

```
t := i.(T)
```

该语句断言接口值 `i` 保存了具体类型 `T`，并将其底层类型为 `T` 的值赋予变量 `t`。

若 `i` 并未保存 `T` 类型的值，该语句就会触发一个 panic。

为了 **判断** 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

```
t, ok := i.(T)
```

若 `i` 保存了一个 `T`，那么 `t` 将会是其底层值，而 `ok` 为 `true`。

否则，`ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，程序并不会产生 panic。

请注意这种语法和读取一个映射时的相同之处。

示例代码：
```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"  // 创建接口  // 其值是 “hello”

	s := i.(string)
	fmt.Println(s)  将i的string赋值给s
 
	s, ok := i.(string)  // 这里可以检查接口的值  hello true
	fmt.Println(s, ok)

	f, ok := i.(float64)  // 上面是匹配正确：这里是匹配失败  0 false
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

```

> 关于什么是go语言里的断言？：
> 检查接口值是否保存了特定类型的值的方式。类型断言的语法如下：

### 类型选择

示例代码：

```go
package main

import "fmt"

// 传入的i是一个空接口，其中参数值以及类型不明确：由后面的switch语句来判断
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("二倍的 %v 是 %v\n", v, v*2)
	case string:
		fmt.Printf("%q 长度为 %v 字节\n", v, len(v))
	default:
		fmt.Printf("我不知道类型 %T!\n", v)
	}
    // 几种基本类型的判断
}

func main() {
	do(21)
	do("hello")
	do(true)
}

```

### Stringer

`fmt`包中定义的 `Stringer`是最普遍的接口之一。

```
type Stringer interface {
    String() string
}
// Stringer接口的实现
```

`Stringer` 是一个可以用字符串描述自己的类型。`fmt` 包（还有很多包）都通过此接口来打印值。

示例代码：
```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
// 创建结构体  ： 当然这个结构体看着没啥问题

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
// 创建方法 ： 当然这个方法看着也是没啥问题的

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}  // 创建结构体对象
    fmt.Println(a, z)  // 调用方法：：由于Stringer接口  默认i盗用了String()方法：将两个字符串拼接起来了
}
// 别想太多 就是普通的输出结果的链式凭借罢了 也就是 +  +  +  +  效果是一样的
```

运行结果如下：
```go
Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

解释如下：
```go
因为 Person 类型实现了 Stringer 接口，所以 fmt.Println 会调用 a 和 z 的 String 方法，然后打印返回的字符串。
```

有点像是隐式调用方法：虽然不是直接调用：但是间接调用了

示例代码：
```go
package main

import "fmt"

type IPAddr [4]byte

// 为 IPAddr 添加一个 "String() string" 方法。
func (ip IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
```

### 错误类型

Go 程序使用 `error` 值来表示错误状态。

与 `fmt.Stringer` 类似，`error` 类型是一个内建接口：

```
type error interface {
    Error() string
}
```

（与 `fmt.Stringer` 类似，`fmt` 包也会根据对 `error` 的实现来打印值。）

通常函数会返回一个 `error` 值，调用它的代码应当判断这个错误是否等于 `nil` 来进行错误处理。

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

`error` 为 nil 时表示成功；非 nil 的 `error` 表示失败。

示例代码：

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}
// 创建简单MyError的结构体

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
// 创建一个深用方法 ：输出结果是一个string


// run函数的定义在这里 ：放回置是一个error类型
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}  // 返回的是一个MyError结构体的指针
}
// 创建一个普通的函数 ：： 注意返回值：

func main() {
	if err := run(); err != nil {  // 创建if语句的时候赋值：和将赋值语句移出去效果一直
		fmt.Println(err)
	}
    // 先调用run函数 然后检查err是不是错误：如果是错入就答应err的错误信息i
}

```

示例错误：

```go
at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work

```

### Readers

> 关于输入输出流相关的接口

`io` 包指定了 `io.Reader` 接口，它表示数据流的读取端。

Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。

`io.Reader` 接口有一个 `Read` 方法：

```
func (T) Read(b []byte) (n int, err error)
```

`Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

示例代码创建了一个 `strings.Reader` 并以每次 8 字节的速度读取它的输出。

示例代码：

```go
package main

import (
	"fmt"
	"io"  // 引入io包
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")  // 创建String

	b := make([]byte, 8)  // 创建byte切片 共8个容量
	for {
		n, err := r.Read(b)  // 将r的内容写入b中
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)  // 每次读取操作的结果：：
		fmt.Printf("b[:n] = %q\n", b[:n])  // 输出实际读取的值
		if err == io.EOF {
			break
		}
	}
    // 无限循环：从r中取值：直到EOF
}

```

运行结果是：
```go
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""
```

稍微区分下字符的二进制表示以及b的实际含义：以及每次读取的n的意思是什么：：：：：b储存的是字符的二进制表示

### 图像

> go语言中的图像接口

`image` 包定义了 `Image` 接口：

```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**注意:** `Bounds` 方法的返回值 `Rectangle` 实际上是一个 `image.Rectangle`)，它在 `image` 包中声明。

`color.Color` 和 `color.Model` 类型也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由 `image/color` 包定义。

示例代码：

```go
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
// 调用图像接口推至Rectangle
```

```go
关于为什么没有绘制出一个矩形出来：？
你的代码创建了一个100x100的RGBA图像，并打印了图像的边界和(0,0)位置的颜色。但是，它并没有绘制出图像。

在Go语言中，image包用于处理图像，但它并不包含绘制图像的功能。要在屏幕上显示图像，你需要使用其他的库，例如image/draw包可以用于绘制图像，image/png或image/jpeg包可以用于将图像保存为文件。
```

下面是一个完整的可绘制出一个矩形的代码：

```go
package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
)

func main() {
    // 创建一个100x100的RGBA图像
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))

    // 创建一个红色
    red := color.RGBA{255, 0, 0, 255}

    // 遍历图像的每个像素，将其设置为红色
    for y := 0; y < img.Bounds().Dy(); y++ {
        for x := 0; x < img.Bounds().Dx(); x++ {
            img.Set(x, y, red)
        }
    }

    // 创建一个文件
    file, err := os.Create("rectangle.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // 将图像编码为PNG并写入文件
    png.Encode(file, img)
}
// 绘制图像并且将图像存入文件中
```

> 绘制图像到文件中：终端中不可能直接绘制图像：下面是创建web服务器进行图像绘制

示例代码：
```go
package main

import (
    "image"
    "image/color"
    "image/png"
    "net/http"
)
// 导包

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {  // 处理函数：当用户访问 / 时会被调用： 
        // 创建一个100x100的RGBA图像
        img := image.NewRGBA(image.Rect(0, 0, 100, 100))
        // 像是绘制像素点：再由后面的来填充颜色

        red := color.RGBA{255, 0, 0, 255}
        // 填充红色

        // 遍历图像的每个像素，将其设置为红色
        for y := 0; y < img.Bounds().Dy(); y++ {
            for x := 0; x < img.Bounds().Dx(); x++ {
                img.Set(x, y, red)
            }
        }
        // 遍历将每个像素点绘制为红色

        // 将图像编码为PNG并写入响应
        w.Header().Set("Content-Type", "image/png")  // 编码并且写入响应
        png.Encode(w, img)
    })

    // 启动web服务器
    http.ListenAndServe(":8080", nil)  // 监听端口
}
```

> 顺便学下这里的web服务器编程
>
> 关于这比较重要的两句：
>
> 1. `w http.ResponseWriter`：这是一个接口，它提供了一组方法，用于构造HTTP响应。你可以通过这个接口写入响应头（header）和响应体（body）。
> 2. `r *http.Request`：这是一个指向http.Request结构体的指针，它包含了HTTP请求的所有信息，如请求方法（GET、POST等）、URL、头部信息、请求体等。

## 泛型

### 类型参数

可以使用类型参数编写 Go 函数来处理多种类型。 函数的类型参数出现在函数参数之前的方括号之间。

```
func Index[T comparable](s []T, x T) int
// 创建一个函数：Index 参数是T的切片s和一个类型为T的值x。   返回值是int
```

此声明意味着 `s` 是满足内置约束 `comparable` 的任何类型 `T` 的切片。 `x` 也是相同类型的值。

`comparable` 是一个有用的约束，它能让我们对任意满足该类型的值使用 `==` 和 `!=` 运算符。在此示例中，我们使用它将值与所有切片元素进行比较，直到找到匹配项。 该 `Index` 函数适用于任何支持比较的类型。

示例代码：

```go
package main

import "fmt"

// Index 返回 x 在 s 中的下标，未找到则返回 -1。
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v 和 x 的类型为 T，它拥有 comparable 可比较的约束，
		// 因此我们可以使用 ==。
		if v == x {
			return i
		}  //有点像是查找函数  ；； 找对对应的下标i 然后返回 否则返回 -1
	}
	return -1
}

func main() {
	// Index 可以在整数切片上使用
	si := []int{10, 20, 15, -10}  // 创建切片数组
	fmt.Println(Index(si, 15))

	// Index 也可以在字符串切片上使用
	ss := []string{"foo", "bar", "baz"}  // 创建string切片数组
	fmt.Println(Index(ss, "hello"))  
}

```

```go
// 简略解释如下：
这个程序定义了一个泛型函数 Index，并在 main 函数中对其进行了调用。这个函数使用了 Go 1.18 引入的泛型特性。

Index 函数接受两个参数：一个类型为 T 的切片 s 和一个类型为 T 的值 x。这里的 T 是一个类型参数，表示任何可比较的类型。函数返回一个 int 类型的值，表示 x 在 s 中的索引，如果 s 中没有 x，则返回 -1。

在 Index 函数中，使用 range 关键字遍历切片 s，对于每个元素 v，如果 v 等于 x，则返回当前的索引 i。如果遍历完整个切片都没有找到等于 x 的元素，就返回 -1。

在 main 函数中，首先创建了一个整数切片 si 和一个字符串切片 ss，然后调用 Index 函数查找 15 在 si 中的索引和 "hello" 在 ss 中的索引，然后打印这两个索引。因为 15 在 si 中的索引为 2，而 ss 中没有 "hello"，所以输出将是 2 和 -1。
```

关于[T comparable]的解释如下：
```go
[T comparable] 是一个类型参数列表。这是Go 1.18版本引入的泛型特性的一部分。

在这个列表中，T 是一个类型参数，你可以把它看作是一个占位符，它可以代表任何类型。在函数调用时，调用者可以用实际的类型来替换它。

comparable 是一个类型约束，它限制了 T 可以代表的类型。comparable 是Go语言预定义的一个约束，表示任何可以进行 == 和 != 操作的类型。这包括所有的基本类型，如int、float64、string等，以及由这些类型组成的数组和结构体。

所以，[T comparable] 的意思是，T 是一个可以代表任何可比较类型的类型参数。
// 简言之就是泛型比价参数  -->  泛型函数
```

### 泛型类型

除了泛型函数之外，Go 还支持泛型类型。 类型可以使用类型参数进行参数化，这对于实现通用数据结构非常有用。

此示例展示了能够保存任意类型值的单链表的简单类型声明。

作为练习，请为此链表的实现添加一些功能。

示例代码：

```go
package main

import "fmt"

// List 表示一个可以保存任何类型的值的单链表。
type List[T any] struct {
    next *List[T]  // 表示链表的next
    val  T  // 表示当前值
}
// 创建结构体 -->  用于表示链表

// New 创建一个新的链表。
func New[T any](val T) *List[T] {  // T any 可以看作是任意参数：也就是不确定当创建的链表是什么类型的
    return &List[T]{val: val}
}

// Add 在链表的末尾添加一个新的元素。
func (l *List[T]) Add(val T) {  // 对已有的链表头 的尾部添加 一个新节点
    for l.next != nil {
        l = l.next
    }
    l.next = &List[T]{val: val}
}

// Print 打印链表的所有元素。
func (l *List[T]) Print() {
    for l != nil {
        fmt.Println(l.val)
        l = l.next
    }
}

func main() {
    // 创建一个整数链表
    l := New[int](1)  // 创建int参数的链表
    l.Add(2)
    l.Add(3)
    l.Print() // 输出：1 2 3

    // 创建一个字符串链表
    s := New[string]("hello")
    s.Add("world")
    s.Print() // 输出：hello world
}
```

> 注意：泛型有点像是c++里面的模板创建各种类型的容器
>
> 注意：go语言里面是有链表的板子的  包名   "container/list"   

## 并发

### Go 协程  ： 简称 go程

Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

```
go f(x, y, z)   // 通过go来运行go程
```

会启动一个新的 Go 协程并执行

```
f(x, y, z)
```

`f`, `x`, `y` 和 `z` 的求值发生在当前的 Go 协程中，而 `f` 的执行发生在新的 Go 协程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。[`sync`](https://go-zh.org/pkg/sync/) 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。

示例代码：
```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)  // 每次休眠1s
		fmt.Println(s)
	}
    // for循环遍历5遍
}
// 创建一个go程

func main() {
	go say("world")  // 运行go程  // 当然这个go程和主线程main是并发执行的
	say("hello")  // 普通函数调用
}
// 运行结果可见是并发执行的
```

### 信道(先入先出)

信道是带有类型的管道，你可以通过它用信道操作符 `<-` 来发送或者接收值。

```
ch <- v    // 将 v 发送至信道 ch。    // 发送
v := <-ch  // 从 ch 接收值并赋予 v。  // 接收
```

（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

```
ch := make(chan int)
```

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

> 关于ch := make(chan int)这行代码：
> `make(chan int)` 是创建通道的语法，其中 `chan int` 表示这是一个整数类型的通道。由于没有指定缓冲区大小，所以这是一个无缓冲通道，也就是说，发送和接收操作都是阻塞的，直到另一方准备好为止。
>
> 没有指定缓冲区大小 是一个无缓冲的通道
>
> 下面是创建有缓冲区的通道：
> ch := make(chan int, 10)

示例代码：
```go
package main

import "fmt"

func sum(s []int, c chan int) {  // 接收参数是一个int数组s，一个通道c
	sum := 0
	for _, v := range s {  // 遍历数组s将结果累加到sum上
		sum += v
	}
	c <- sum // 发送 sum 到 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}  // 创建原始数组

	c := make(chan int)  // 创建int通道
	go sum(s[:len(s)/2], c)  // 开启go程  // 这里应该是17
	go sum(s[len(s)/2:], c)  // 开启go程  // 这里应该是-5
	x, y := <-c, <-c // 从 c 接收 
	// x的结果是-5  y的结果是17  x+y的结果是12  // 可见信道的读取像是读去堆栈一样 X ：：实际原因是因为go程的执行顺序不定
	fmt.Println(x, y, x+y)
}

```

### 带缓冲的信道

信道可以是 **带缓冲的**。将缓冲长度作为第二个参数提供给 `make` 来初始化一个带缓冲的信道：

```
ch := make(chan int, 100)  // 有初始化长度的信道就是带缓冲区的信道
```

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

修改示例填满缓冲区，然后看看会发生什么。

示例代码：

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)  // 创建长度为2的信道
	ch <- 1
	ch <- 2  // 向信道中写入两个数
	fmt.Println(<-ch) 
	fmt.Println(<-ch)
}

```

### range 和 close

发送者可通过 `close` 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

```
v, ok := <-ch
```

此时 `ok` 会被设置为 `false`。

循环 `for i := range c` 会不断从信道接收值，直到它被关闭。

**注意**： 只应由发送者关闭信道，而不应油接收者关闭。向一个已经关闭的信道发送数据会引发程序 panic。

**还要注意**： 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 `range` 循环。

示例代码：
```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x  // 每次将x读入到信道c中
		x, y = y, x+y
	}
	close(c)  // 关闭信道
}
// 创建函数

func main() {
	c := make(chan int, 10)  // 创建10个缓存空间
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
    // 将信道中的数据全部取出来  实际上取出来的顺序是顺序的：“那么之前说像堆栈一样的说法是错误的
}

```

### select 语句

`elect` 语句使一个 Go 程可以等待多个通信操作。

`select` 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

示例代码：

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {  // 传入两个信道c和quit
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
// 创建函数

func main() {
	c := make(chan int)
	quit := make(chan int)  // 创建两个无缓冲区的信道
    
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
            // 输出信道c中的值
		}
		quit <- 0  // 表示终止
	}()  //() 是对这个匿名函数的调用。这意味着，当 Go 运行到这一行代码时，它会立即执行这个匿名函数的函数体。
    // 创建go程
    
	fibonacci(c, quit)  
    // 无限制的计算 ；； 知道select 选择了quit
}

```

> 关于go程的理解有很大作用，稍微多家理解下

### 默认选择

当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

```
select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
```

示例代码：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // 定时器100ms触发
	boom := time.After(500 * time.Millisecond) // 定时器500ms触发
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

```

输出结果是：
```go
    .
    .
tick.
    .
    .
tick.
    .
    .
    .
tick.
    .
tick.
    .
    .
tick.
BOOM!
// tick是没100ms执行一次，BOOM是每1s执行一次  Tick和After应该是用来表示是否继续的状态码？
```

### sync.Mutex

我们已经看到信道非常适合在各个 Go 程间进行通信。

但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？

这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。

Go 标准库中提供了 `sync.Mutex`互斥锁类型及其两个方法：

- `Lock`
- `Unlock`

我们可以通过在代码前调用 `Lock` 方法，在代码后调用 `Unlock` 方法来保证一段代码的互斥执行。参见 `Inc` 方法。

我们也可以用 `defer` 语句来保证互斥锁一定会被解锁。参见 `Value` 方法。

示例代码：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 是并发安全的
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc 对给定键的计数加一
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// 锁定使得一次只有一个 Go 协程可以访问映射 c.v。
	c.v[key]++
	c.mu.Unlock()
}

// Value 返回给定键的计数的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// 锁定使得一次只有一个 Go 协程可以访问映射 c.v。
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

```



### Web爬虫

在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 `Crawl` 函数来并行地抓取 URL，并且保证不重复。

*提示：* 你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！

示例代码：

```
package main

import (
    "fmt"
)

// Fetcher 是一个接口，定义了获取网页内容和链接的方法
type Fetcher interface {
    // Fetch 方法接收一个 URL，返回该 URL 页面的内容，页面上的所有 URL，以及可能的错误
    Fetch(url string) (body string, urls []string, err error)
}

// Crawl 函数接收一个 URL，一个深度限制和一个 Fetcher 实例
// 它开始于给定的 URL，递归地爬取链接，直到达到给定的深度
func Crawl(url string, depth int, fetcher Fetcher) {
    // 如果深度小于等于0，停止爬取
    if depth <= 0 {
        return
    }
    // 使用 fetcher 获取页面内容和链接
    body, urls, err := fetcher.Fetch(url)
    // 如果获取过程中出现错误，打印错误并停止爬取
    if err != nil {
        fmt.Println(err)
        return
    }
    // 打印找到的页面内容
    fmt.Printf("found: %s %q\n", url, body)
    // 对找到的每个链接，递归地进行爬取
    for _, u := range urls {
        Crawl(u, depth-1, fetcher)
    }
    return
}

func main() {
    // 从 "https://golang.org/" 开始，使用 fetcher 进行爬取，深度限制为 4
    Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是一个模拟的 Fetcher，它使用一个 map 来存储预定义的爬取结果
type fakeFetcher map[string]*fakeResult

// fakeResult 是一个结构体，用于存储一个页面的内容和链接
type fakeResult struct {
    body string
    urls []string
}

// Fetch 方法实现了 Fetcher 接口，它从预定义的结果中查找给定的 URL
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    // 如果 URL 在预定义的结果中，返回对应的内容和链接
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    // 如果 URL 不在预定义的结果中，返回一个错误
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是一个已经填充了预定义结果的 fakeFetcher
var fetcher = fakeFetcher{
    "https://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "https://golang.org/pkg/",
            "https://golang.org/cmd/",
        },
    },
    "https://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "https://golang.org/",
            "https://golang.org/cmd/",
            "https://golang.org/pkg/fmt/",
            "https://golang.org/pkg/os/",
        },
    },
    "https://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
    "https://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "https://golang.org/",
            "https://golang.org/pkg/",
        },
    },
}
//这段代码定义了一个名为 fetcher 的变量，它是 fakeFetcher 类型的一个实例。fakeFetcher 是一个映射（map），它将字符串（URL）映射到 *fakeResult 类型的值。

fakeResult 是一个结构体，包含一个表示页面内容的字符串 body 和一个表示页面上所有链接的字符串切片 urls。
```







