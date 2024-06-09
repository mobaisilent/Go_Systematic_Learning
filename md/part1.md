## 基础部分

### 基本形式

```go
package main
// 当前包名

import "fmt"
// 导入单个包：后面加不加;都问题不大

func main() {
	fmt.Println("hello go")
}
// 使用包里面的Println进行简单的字符串的打印
```

> 注意 函数后面的{ 必须与函数在同一行

### go命令行的使用

#### 直接运行

例如：

```go
go run hello.go
```

然后直接就可以看见结果

#### 编译运行

需要进入对应的文件夹中，前面也一样

```go
go build hello.go
```

##### 指定生成exe的名字

```go
go build -o heiheihei.exe
```

### 关于无法使用vscode用f5进行调试

#### 先安装下列包文件

```go
go install github.com/go-delve/delve/cmd/dlv@latest
go install github.com/mdempsky/gocode@latest
go install github.com/rogpeppe/godef@latest
go install golang.org/x/lint/golint@latest
go install github.com/ramya-rao-a/go-outline@latest
go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go install golang.org/x/tools/cmd/gorename@latest
go install github.com/sqs/goreturns@latest
go install github.com/acroca/go-symbols@latest
go install golang.org/x/tour@gotour
go install golang.org/x/tools/cmd/guru@latest
```

#### 报错信息如下

```go
Starting: C:\Users\27892\Desktop\Golang\bin\dlv.exe dap --listen=127.0.0.1:6895 from c:\Users\27892\Desktop\Golang\src\first
DAP server listening at: 127.0.0.1:6895
Build Error: go build -o c:\Users\27892\Desktop\Golang\src\first\__debug_bin219222439.exe -gcflags all=-N -l .
go: go.mod file not found in current directory or any parent directory; see 'go help modules' (exit status 1)
```

> 从信息中知道缺少module模块

#### 解决方案

```go
cd C:\Users\27892\Desktop\Golang\src\first

go mod init first
```

> 也就是将first模块化

#### 查看mod

```go
module first

go 1.22.4
```

然后就可以快乐调试咯

![image-20240609125609716](./images/image-20240609125609716.png)



## 变量的声明

### 4种基础变量的声明

```go
package main

import "fmt"

func main() {
	//第一种 使用默认值0/...
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
```

> 全局变量的声明只能使用前面三种
>
> := 冒等不适用全局变量的声明

### 多变量的声明

```go
package main


import "fmt"


var x, y int
var ( //这种分解的写法,一般用于声明全局变量
        a int
        b bool
)


var c, d int = 1, 2
var e, f = 123, "liudanbing"


//这种不带声明格式的只能在函数体内声明
//g, h := 123, "需要在func函数体内实现"


func main() {
        g, h := 123, "需要在func函数体内实现"
        fmt.Println(x, y, a, b, c, d, e, f, g, h)


        //不能对g变量再次做初始化声明
        //g := 400


        _, value := 7, 5  //实际上7的赋值被废弃，变量 _  不具备读特性
        //fmt.Println(_) //_变量的是读不出来的
        fmt.Println(value) //5
}
```



## 常量的声明

常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

```go
const identifier [type] = value
```

你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

- 显式类型定义：

```go
const b string = "abc"
```

- 隐式类型定义：

```go
const b = "abc"
```

例如:

```go
package main


import "fmt"


func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" //多重赋值


   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d\n", area)
   println(a, b, c)   
}
```

以上实例运行结果为：

```go
面积为 : 50
1 false str
```

常量还可以用作枚举：

```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

数字 0、1 和 2 分别代表未知性别、女性和男性。

常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：

```go
package main


import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)


func main(){
    println(a, b, c)
}
```

输出结果为：abc, 3, 16

unsafe.Sizeof(a)输出的结果是16 。

字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。

### 优雅的常量 iota

有些概念有名字，并且有时候我们关注这些名字，甚至（特别）是在我们代码中。

```go
const (
    CCVisa            = "Visa"
    CCMasterCard      = "MasterCard"
    CCAmericanExpress = "American Express"
)
```

在其他时候，我们仅仅关注能把一个东西与其他的做区分。有些时候，有些时候一件事没有本质上的意义。比如，我们在一个数据库表中存储产品，我们可能不想以 string 存储他们的分类。我们不关注这个分类是怎样命名的，此外，该名字在市场上一直在变化。

我们仅仅关注它们是怎么彼此区分的。

```go
const (
    CategoryBooks    = 0
    CategoryHealth   = 1
    CategoryClothing = 2
)
```

使用 0, 1, 和 2 代替，我们也可以选择 17， 43， 和 61。这些值是任意的。

在 Go，常量有许多微妙之处。当用好了，可以使得代码非常优雅且易维护的。

### 自增长

在 golang 中，一个方便的习惯就是使用`iota`标示符，它简化了常量用于增长数字的定义，给以上相同的值以准确的分类。

```go
const (
    CategoryBooks = iota // 0
    CategoryHealth       // 1
    CategoryClothing     // 2
)
```

### iota和表达式

`iota`可以做更多事情，而不仅仅是 increment。更精确地说，`iota`总是用于 increment，但是它可以用于表达式，在常量中的存储结果值。

```go
type Allergen int


const (
    IgEggs Allergen = 1 << iota         // 1 << 0 which is 00000001
    IgChocolate                         // 1 << 1 which is 00000010
    IgNuts                              // 1 << 2 which is 00000100
    IgStrawberries                      // 1 << 3 which is 00001000
    IgShellfish                         // 1 << 4 which is 00010000
)
```

这个工作是因为当你在一个`const`组中仅仅有一个标示符在一行的时候，它将使用增长的`iota`取得前面的表达式并且再运用它，。在 Go 语言的[spec](https://legacy.gitbook.com/book/aceld/how-do-go/edit#)中， 这就是所谓的隐性重复最后一个非空的表达式列表.

如果你对鸡蛋，巧克力和海鲜过敏，把这些 bits 翻转到 “on” 的位置（从左到右映射 bits）。然后你将得到一个 bit 值`00010011`，它对应十进制的 19。

```go
fmt.Println(IgEggs | IgChocolate | IgShellfish)


// output:
// 19
type ByteSize float64


const (
    _           = iota     // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)       // 1 << (10*1)
    MB                                   // 1 << (10*2)
    GB                                   // 1 << (10*3)
    TB                                   // 1 << (10*4)
    PB                                   // 1 << (10*5)
    EB                                   // 1 << (10*6)
    ZB                                   // 1 << (10*7)
    YB                                   // 1 << (10*8)
)
```

当你在把两个常量定义在一行的时候会发生什么？

Banana 的值是什么？ 2 还是 3？ Durian 的值又是？

```go
const (
    Apple, Banana = iota + 1, iota + 2
    Cherimoya, Durian
    Elderberry, Fig
)
```

在下一行增长，而不是立即取得它的引用。

```go
// Apple: 1
// Banana: 2
// Cherimoya: 2
// Durian: 3
// Elderberry: 3
// Fig: 4
```

### 小结

常量的定义和iota的使用大概就是这样咯：不要想复杂了

- const{}
- iota自增



## 函数

### 定义

#### 简单定义

```go
package main


import "fmt"


func swap(x, y string) (string, string) {
   return y, x
}


func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}
```

> 简单的定义了一个swap函数

#### 多个返回值的定义

```go
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
```

```go
func fun1(x string, y int) (r1 ,r2 int) {
	return
}
// 这样进行初始化也是正确的
```

### 多个包的执行顺序

#### main同时导入两个包

![image-20240609135819765](./images/image-20240609135819765.png)

##### 创建包1

> Lib1.go 当然放到文件夹 Lab1下

```go
package InitLib1

import "fmt"

func init() {
    fmt.Println("lib1")
}
```

##### 创建包2

> Lib2.go 当然放到文件夹 Lab2下

```go
package InitLib2

import "fmt"

func init() {
    fmt.Println("lib2")
}
```

##### 创建main

```go
package main

import (
    "fmt"
    _ "GolangTraining/InitLib1"
    _ "GolangTraining/InitLib2"
    // 匿名导包：不使用也不会报错
    // 前面还还可以起别名
)

func init() {
    fmt.Println("libmain init")
}

func main() {
    fmt.Println("libmian main")
}
```

#### main导入一个嵌套包

> 一个包嵌套另一个包

略

执行过程就是上面那张图

#### 函数参数



函数如果使用参数，该变量可称为函数的形参。



形参就像定义在函数体内的局部变量。



调用函数，可以通过两种方式来传递参数：



##### 值传递



值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。



默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

> 默认值传递：不影响其对应的实际值

以下定义了 swap() 函数：



```go
/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int


   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/


   return temp;
}
```



接下来，让我们使用值传递来调用 swap() 函数：



```go
package main


import "fmt"


func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200


   fmt.Printf("交换前 a 的值为 : %d\n", a )
   fmt.Printf("交换前 b 的值为 : %d\n", b )


   /* 通过调用函数来交换值 */
   swap(a, b)


   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}


/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int


   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/


   return temp;
}
```



以下代码执行结果为：



交换前 a 的值为 : 100



交换前 b 的值为 : 200



交换后 a 的值 : 100



交换后 b 的值 : 200

------

##### 引用传递(指针传递)



**指针**



Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务。



接下来让我们来一步步学习 Go 语言指针。



我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。



Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。



以下实例演示了变量在内存中地址：



```go
package main


import "fmt"


func main() {
   var a int = 10   


   fmt.Printf("变量的地址: %x\n", &a  )
}
```



执行以上代码输出结果为：



```bash
变量的地址: 20818a220
```

------

现在我们已经了解了什么是内存地址和如何去访问它。接下来我们将具体介绍指针。



引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。



引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递：



```go
/* 定义交换值函数*/
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保持 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```



以下我们通过使用引用传递来调用 swap() 函数：



```go
package main


import "fmt"


func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200


   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )


   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)
   // 由于是指针传递：所以实际上修改了a，b其对应的值


   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}


func swap(x *int, y *int) {  // 传递两个指针
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```



以上代码执行结果为：



交换前，a 的值 : 100



交换前，b 的值 : 200



交换后，a 的值 : 200



交换后，b 的值 : 100

> 所以说  go语言的指针的使用和c语言基本一致了





## defer

defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。

 执行机制：就是 栈操作

defer作用：

- 释放占用的资源
- 捕捉处理异常
- 输出日志

### 简单的展示

如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。

```go
func Demo(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}
func main() {
	Demo()
}
```

```go
4
3
2
1
```



###  recover错误拦截

运行时panic异常一旦被引发就会导致程序崩溃。

Go语言提供了专用于“拦截”运行时panic的内建函数“recover”。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。

**注意：**recover只有在defer调用的函数中有效。

如果程序没有异常，不会打印错误信息。

```go
func recover interface{}
package main

import "fmt"

func Demo(i int) {
	//定义10个元素的数组
	var arr [10]int
	//错误拦截要在产生错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产生panic异常  打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()
	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10

}

func main() {
	Demo(10)
	//产生错误后 程序继续
	fmt.Println("程序继续执行...")
}
```

```go
runtime error: index out of range
程序继续执行...
```

### 执行顺序

```go
package main

import "fmt"

func deferFunc() int {
    fmt.Println("defer func called ... ")
    return 0
}
// 定义defer函数

func returnFunc() int {
    fmt.Println("return func called ... ")
    return 0
}
// 定义retuen函数

func returnAndDefer() int {
    defer deferFunc()  // defer压入defer函数
    return returnFunc()  // 返回return函数
}

func main() {
    returnAndDefer()
}
```

执行结果：

```go
return func called ... 
defer func called ... 
```



## 切片 slice

**slice**

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片`("动态数组")`,与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

#### 定义切片

你可以声明一个未指定大小的数组来定义切片：

```go
var identifier []type
```

切片不需要说明长度。

或使用make()函数来创建切片:

```go
var slice1 []type = make([]type, len)


也可以简写为


slice1 := make([]type, len)
```

也可以指定容量，其中capacity为可选参数。

```go
make([]T, length, capacity)
```

这里 len 是数组的长度并且也是切片的初始长度。

#### 切片初始化

> 总结就是 用不用var  ||  make  ||  是否初始化空间

```go
s :=[] int {1,2,3 }
```

直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

```go
s := arr[:]
```

初始化切片s,是数组arr的引用

```go
s := arr[startIndex:endIndex]
```

将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

```go
s := arr[startIndex:]
```

缺省endIndex时将表示一直到arr的最后一个元素

```go
s := arr[:endIndex]
```

缺省startIndex时将表示从arr的第一个元素开始

```go
s1 := s[startIndex:endIndex]
```

通过切片s初始化切片s1

```go
s :=make([]int,len,cap)
```

通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片



#### 切片的好处

##### 简单使用数组：

实例代码：

```go
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
```



##### 使用切片数组

```go
package main

import "fmt"

func printArray(myArray []int) {
	// 引用传递
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4}
	// 动态数组 切片 slice
	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}

```



#### len() 和 cap() 函数

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

以下为具体实例：

```go
package main


import "fmt"


func main() {
   var numbers = make([]int,3,5)


   printSlice(numbers)
}


func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

以上实例运行输出结果为:

```go
len=3 cap=5 slice=[0 0 0]
```

#### 空(nil)切片

一个切片在未初始化之前默认为 nil，长度为 0，实例如下：

```go
package main


import "fmt"


func main() {
   var numbers []int


   printSlice(numbers)


   if(numbers == nil){
      fmt.Printf("切片是空的")
   }
}


func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

以上实例运行输出结果为:

```go
len=0 cap=0 slice=[]
切片是空的
```

#### 切片截取

可以通过设置下限及上限来设置截取切片*[lower-bound:upper-bound]*，实例如下：

```go
package main


import "fmt"


func main() {
   /* 创建切片 */
   numbers := []int{0,1,2,3,4,5,6,7,8}   
   printSlice(numbers)


   /* 打印原始切片 */
   fmt.Println("numbers ==", numbers)


   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])


   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])


   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])


   numbers1 := make([]int,0,5)
   printSlice(numbers1)


   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2)


   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)


}


func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

执行以上代码输出结果为：



```go
len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
numbers == [0 1 2 3 4 5 6 7 8]
numbers[1:4] == [1 2 3]
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
len=0 cap=5 slice=[]
len=2 cap=9 slice=[0 1]
len=3 cap=7 slice=[2 3 4]
```

#### apped() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。

```go
package main


import "fmt"


func main() {
   var numbers []int
   printSlice(numbers)


   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)


   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)


   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)


   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)


   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)   
}


func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

以上代码执行输出结果为：

```go
len=0 cap=0 slice=[]
len=1 cap=1 slice=[0]
len=2 cap=2 slice=[0 1]
len=5 cap=6 slice=[0 1 2 3 4]
len=5 cap=12 slice=[0 1 2 3 4]
```

> 当长度溢出时依旧append 结果就是导致cap的长度直接变为原来cap长度的两倍：倍增


------

#### **map**

map和slice类似，只不过是数据结构不同，下面是map的一些声明方式。	

```go
package main
import (
    "fmt"
)

func main() {
    //第一种声明
    var test1 map[string]string
    //在使用map前，需要先make，make的作用就是给map分配数据空间
    test1 = make(map[string]string, 10) 
    test1["one"] = "php"
    test1["two"] = "golang"
    test1["three"] = "java"
    fmt.Println(test1) //map[two:golang three:java one:php]


    //第二种声明
    test2 := make(map[string]string)
    test2["one"] = "php"
    test2["two"] = "golang"
    test2["three"] = "java"
    fmt.Println(test2) //map[one:php two:golang three:java]

    //第三种声明
    test3 := map[string]string{
        "one" : "php",
        "two" : "golang",
        "three" : "java",
    }
    fmt.Println(test3) //map[one:php two:golang three:java]


    
    language := make(map[string]map[string]string)
    language["php"] = make(map[string]string, 2)
    language["php"]["id"] = "1"
    language["php"]["desc"] = "php是世界上最美的语言"
    language["golang"] = make(map[string]string, 2)
    language["golang"]["id"] = "2"
    language["golang"]["desc"] = "golang抗并发非常good"
    
    fmt.Println(language) //map[php:map[id:1 desc:php是世界上最美的语言] golang:map[id:2 desc:golang抗并发非常good]]


    //增删改查
    // val, key := language["php"]  //查找是否有php这个子元素
    // if key {
    //     fmt.Printf("%v", val)
    // } else {
    //     fmt.Printf("no");
    // }

    //language["php"]["id"] = "3" //修改了php子元素的id值
    //language["php"]["nickname"] = "啪啪啪" //增加php元素里的nickname值
    //delete(language, "php")  //删除了php子元素
    fmt.Println(language)
}
```

map常见声明方式剪辑

```go
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

```



## 面向对象特征

### 方法

假设有两个方法，一个方法的接收者是指针类型，一个方法的接收者是值类型，那么：



- 对于值类型的变量和指针类型的变量，这两个方法有什么区别？
- 如果这两个方法是为了实现一个接口，那么这两个方法都可以调用吗？
- 如果方法是嵌入到其他结构体中的，那么上面两种情况又是怎样的？



```go
package main


import "fmt"


//定义一个结构体
type T struct {
    name string
}


func (t T) method1() {
    t.name = "new name1"
}


func (t *T) method2() {
    t.name = "new name2"
}


func main() {


    t := T{"old name"}


    fmt.Println("method1 调用前 ", t.name)
    t.method1()
    fmt.Println("method1 调用后 ", t.name)


    fmt.Println("method2 调用前 ", t.name)
    t.method2()
    fmt.Println("method2 调用后 ", t.name)
}
```



运行结果：



```bash
method1 调用前  old name
method1 调用后  old name
method2 调用前  old name
method2 调用后  new name2
```



当调用`t.method1()`时相当于`method1(t)`，实参和行参都是类型 T，可以接受。此时在`method1`()中的t只是参数t的值拷贝，所以`method1`()的修改影响不到main中的t变量。



当调用`t.method2()`=>`method2(t)`，这是将 T 类型传给了 *T 类型，go可能会取 t 的地址传进去：`method2(&t)`。所以 `method1`() 的修改可以影响 t。



T 类型的变量这两个方法都是拥有的。

------

### 方法值和方法表达式



##### 方法值



我们经常选择一个方法，并且在同一个表达式里执行，比如常见的p.Distance()形式，实际上将其分成两步来执行也是可能的。p.Distance叫作“选择器”，选择器会返回一个方法"值"`一个将方法(Point.Distance)绑定到特定接收器变量的函数`。这个函数可以不通过指定其接收器即可被调用；即调用时不需要指定接收器，只要传入函数的参数即可：



```go
package main


import "fmt"
import "math"


type Point struct{ X, Y float64 }


//这是给struct Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}


func main() {


    p := Point{1, 2}
    q := Point{4, 6}


    distanceFormP := p.Distance   // 方法值(相当于C语言的函数地址,函数指针)
    fmt.Println(distanceFormP(q)) // "5"
    fmt.Println(p.Distance(q))    // "5"


    //实际上distanceFormP 就绑定了 p接收器的方法Distance


    distanceFormQ := q.Distance   //
    fmt.Println(distanceFormQ(p)) // "5"
    fmt.Println(q.Distance(p))    // "5"


    //实际上distanceFormQ 就绑定了 q接收器的方法Distance
}
```



在一个包的API需要一个函数值、且调用方希望操作的是某一个绑定了对象的方法的话，方法"值"会非常实用.



举例来说，下面例子中的time.AfterFunc这个函数的功能是在指定的延迟时间之后来执行一个(译注：另外的)函数。且这个函数操作的是一个Rocket对象r



```go
type Rocket struct { /* ... */ }
func (r *Rocket) Launch() { /* ... */ }
r := new(Rocket)
time.AfterFunc(10 * time.Second, func() { r.Launch() })
```



直接用方法"值"传入AfterFunc的话可以更为简短：



```go
time.AfterFunc(10 * time.Second, r.Launch)
```



省掉了上面那个例子里的匿名函数。



##### 方法表达式



和方法"值"相关的还有方法表达式。当调用一个方法时，与调用一个普通的函数相比，我们必须要用选择器(p.Distance)语法来指定方法的接收器。



当T是一个类型时，方法表达式可能会写作`T.f`或者`(*T).f`，会返回一个函数"值"，这种函数会将其第一个参数用作接收器，所以可以用通常(译注：不写选择器)的方式来对其进行调用：



```go
package main


import "fmt"
import "math"


type Point struct{ X, Y float64 }


//这是给struct Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}


func main() {


    p := Point{1, 2}
    q := Point{4, 6}




    distance1 := Point.Distance //方法表达式, 是一个函数值(相当于C语言的函数指针)
    fmt.Println(distance1(p, q))
    fmt.Printf("%T\n", distance1) //%T表示打出数据类型 ,这个必须放在Printf使用


    distance2 := (*Point).Distance //方法表达式,必须传递指针类型
    distance2(&p, q)
    fmt.Printf("%T\n", distance2)


}
```



执行结果



```go
5
func(main.Point, main.Point) float64
func(*main.Point, main.Point) float64
// 这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
// 但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
// 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
// 看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
```



当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。你可以根据选择来调用接收器各不相同的方法。下面的例子，变量op代表Point类型的addition或者subtraction方法，Path.TranslateBy方法会为其Path数组中的每一个Point来调用对应的方法：



```go
package main


import "fmt"
import "math"


type Point struct{ X, Y float64 }


//这是给struct Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
        return math.Hypot(q.X-p.X, q.Y-p.Y)
}


func (p Point) Add(another Point) Point {
        return Point{p.X + another.X, p.Y + another.Y}
}


func (p Point) Sub(another Point) Point {
        return Point{p.X - another.X, p.Y - another.Y}
}


func (p Point) Print() {
        fmt.Printf("{%f, %f}\n", p.X, p.Y)
}


//定义一个Point切片类型 Path
type Path []Point


//方法的接收器 是Path类型数据, 方法的选择器是TranslateBy(Point, bool)
func (path Path) TranslateBy(another Point, add bool) {
        var op func(p, q Point) Point //定义一个 op变量 类型是方法表达式 能够接收Add,和 Sub方法
        if add == true {
                op = Point.Add //给op变量赋值为Add方法
        } else {
                op = Point.Sub //给op变量赋值为Sub方法
        }


        for i := range path {
                //调用 path[i].Add(another) 或者 path[i].Sub(another)
                path[i] = op(path[i], another)
                path[i].Print()
        }
}


func main() {

        points := Path{
                {10, 10},
                {11, 11},
        }


        anotherPoint := Point{5, 5}


        points.TranslateBy(anotherPoint, false)


        fmt.Println("------------------")


        points.TranslateBy(anotherPoint, true)
}
```



运行结果：



```go
{5.000000, 5.000000}
{6.000000, 6.000000}
------------------
{10.000000, 10.000000}
{11.000000, 11.000000}
```





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
> 3. **赋值给接口变：** 





## 万能接口

> 空接口->断言->万能类型

Golang的语言中提供了断言的功能。golang中的所有程序都实现了interface{}的接口，这意味着，所有的类型如string,int,int64甚至是自定义的struct类型都就此拥有了interface{}的接口，这种做法和java中的Object类型比较类似。那么在一个数据通过func funcName(interface{})的方式传进来的时候，也就意味着这个参数被自动的转为interface{}的类型。



```go
func funcName(a interface{}) string {
     return string(a)
}
```



编译器会返回



```go
cannot convert a (type interface{}) to type string: need type assertion
```



此时，意味着整个转化的过程需要类型断言。类型断言有以下几种形式：



1）直接断言使用



```go
var a interface{}


fmt.Println("Where are you,Jonny?", a.(string))
```



但是如果断言失败一般会导致panic的发生。所以为了防止panic的发生，我们需要在断言前进行一定的判断



```go
value, ok := a.(string)
```



如果断言失败，那么ok的值将会是false,但是如果断言成功ok的值将会是true,同时value将会得到所期待的正确的值。示例：



```go
value, ok := a.(string)
if !ok {
    fmt.Println("It's not ok for type string")
    return
}
fmt.Println("The value is ", value)
```



完整例子如下：



```go
package main


import "fmt"


/*
func funcName(a interface{}) string {
        return string(a)
}
*/


func funcName(a interface{}) string {
        value, ok := a.(string)
        if !ok {
                fmt.Println("It is not ok for type string")
                return ""
        }
        fmt.Println("The value is ", value)


        return value
}


func main() {
        //      str := "123"
        //      funcName(str)
        //var a interface{}
        //var a string = "123"
        var a int = 10
        funcName(a)
}
```



2）配合switch使用



```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T", t)       // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```



或者如下使用方法



```go
func sqlQuote(x interface{}) string {
    if x == nil {
        return "NULL"
    } else if _, ok := x.(int); ok {
        return fmt.Sprintf("%d", x)
    } else if _, ok := x.(uint); ok {
        return fmt.Sprintf("%d", x)
    } else if b, ok := x.(bool); ok {
        if b {
            return "TRUE"
        }
        return "FALSE"
    } else if s, ok := x.(string); ok {
        return sqlQuoteString(s) // (not shown)
    } else {
        panic(fmt.Sprintf("unexpected type %T: %v", x, x))
    }
}
```





## **反射reflect**

#### 编程语言中反射的概念



在计算机科学领域，反射是指一类应用，它们能够自描述和自控制。也就是说，这类应用通过采用某种机制来实现对自己行为的描述（self-representation）和监测（examination），并能根据自身行为的状态和结果，调整或修改应用所描述行为的状态和相关的语义。



每种语言的反射模型都不同，并且有些语言根本不支持反射。Golang语言实现了反射，反射机制就是在运行时动态的调用对象的方法和属性，官方自带的reflect包就是反射相关的，只要包含这个包就可以使用。



多插一句，Golang的gRPC也是通过反射实现的。



#### interface 和 反射



在讲反射之前，先来看看Golang关于类型设计的一些原则



- 变量包括（type, value）两部分
- type 包括 `static type`和`concrete type`. 简单来说 `static type`是你在编码是看见的类型(如int、string)，`concrete type`是`runtime`系统看见的类型
- 类型断言能否成功，取决于变量的`concrete type`，而不是`static type`. 因此，一个 `reader`变量如果它的`concrete type`也实现了`write`方法的话，它也可以被类型断言为`writer`.



接下来要讲的`反射`，就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type），在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type），只有interface类型才有反射一说。



在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型:



```go
(value, type)
```



value是实际变量值，type是实际变量的类型。一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。



例如，创建类型为*os.File的变量，然后将其赋给一个接口变量r：



```go
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

var r io.Reader
r = tty
```



接口变量r的pair中将记录如下信息：(tty, *os.File)，这个pair在接口变量的连续赋值过程中是不变的，将接口变量r赋给另一个接口变量w:



```csharp
var w io.Writer
w = r.(io.Writer)
```



接口变量w的pair与r的pair相同，都是:(tty, *os.File)，即使w是空接口类型，pair也是不变的。



interface及其pair的存在，是Golang中实现反射的前提，理解了pair，就更容易理解反射。反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。



```go
package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)
	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}
```



再比如:



```go
package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book.")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book.")
}

func main() {
	b := &Book{}
    //Book指针

	var r Reader
	r = b
    //b被动作Reader使用

	r.ReadBook()

	var w Writer
	w = r.(Writer)
    //b被当作Writer使用
	w.WriteBook()
}
```



#### Golang的反射reflect



**reflect的基本功能TypeOf和ValueOf**



既然反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。那么在Golang的reflect反射包中有什么样的方式可以让我们直接获取到变量内部的信息呢？ 它提供了两种类型（或者说两个方法）让我们可以很容易的访问接口变量内容，分别是reflect.ValueOf() 和 reflect.TypeOf()，看看官方的解释



```go
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i.  ValueOf(nil) returns the zero 
func ValueOf(i interface{}) Value {...}

//ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0


// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type {...}

//TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
```



reflect.TypeOf()是获取pair中的type，reflect.ValueOf()获取pair中的value，示例如下：



```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var num float64 = 1.2345

    fmt.Println("type: ", reflect.TypeOf(num))
    fmt.Println("value: ", reflect.ValueOf(num))
}

运行结果:
type:  float64
value:  1.2345
```



说明



1. reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
2. reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
3. 也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种



**从relfect.Value中获取接口interface的信息**



当执行reflect.ValueOf(interface)之后，就得到了一个类型为”relfect.Value”变量，可以通过它本身的Interface()方法获得接口变量的真实内容，然后可以通过类型判断进行转换，转换为原有真实类型。不过，我们可能是已知原有类型，也有可能是未知原有类型，因此，下面分两种情况进行说明。



*已知原有类型*【进行“强制转换”】



已知类型后转换为其对应的类型的做法如下，直接通过Interface方法然后强制转换，如下：



```csharp
realValue := value.Interface().(已知的类型)
```



示例如下：



```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var num float64 = 1.2345

    pointer := reflect.ValueOf(&num)
    value := reflect.ValueOf(num)

    // 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
    // Golang 对类型要求非常严格，类型一定要完全符合
    // 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
    convertPointer := pointer.Interface().(*float64)
    convertValue := value.Interface().(float64)

    fmt.Println(convertPointer)
    fmt.Println(convertValue)
}

运行结果：
0xc42000e238
1.2345
```

如果直接打印：

```go
fmt.Println(pointer)
fmt.Println(value)
```

这两个结果是：

```go
<pointer to float64 Value>
<float64 Value>
```

像是显示形式一样



说明



1.  转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！ 
2.  转换的时候，要区分是指针还是指 
3.  也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量” 



未知原有类型【遍历探测其Filed】



很多情况下，我们可能并不知道其具体类型，那么这个时候，该如何做呢？需要我们进行遍历探测其Filed来得知，示例如下:



```go
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) ReflectCallFunc() {
    fmt.Println("Allen.Wu ReflectCallFunc")
}

func main() {

    user := User{1, "Allen.Wu", 25}

    DoFiledAndMethod(user)

}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

    getType := reflect.TypeOf(input)
    fmt.Println("get Type is :", getType.Name())

    getValue := reflect.ValueOf(input)
    fmt.Println("get all Fields is:", getValue)

    // 获取方法字段
    // 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
    // 2. 再通过reflect.Type的Field获取其Field
    // 3. 最后通过Field的Interface()得到对应的value
    for i := 0; i < getType.NumField(); i++ {
        field := getType.Field(i)
        value := getValue.Field(i).Interface()
        fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
    }

    // 获取方法
    // 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
    for i := 0; i < getType.NumMethod(); i++ {
        m := getType.Method(i)
        fmt.Printf("%s: %v\n", m.Name, m.Type)
    }
}

运行结果：
get Type is : User
get all Fields is: {1 Allen.Wu 25}
Id: int = 1
Name: string = Allen.Wu
Age: int = 25
ReflectCallFunc: func(main.User)
```



说明



通过运行结果可以得知获取未知类型的interface的具体变量及其类型的步骤为：



1. 先获取interface的reflect.Type，然后通过NumField进行遍历
2. 再通过reflect.Type的Field获取其Field
3. 最后通过Field的Interface()得到对应的value



通过运行结果可以得知获取未知类型的interface的所属方法（函数）的步骤为：



1. 先获取interface的reflect.Type，然后通过NumMethod进行遍历
2. 再分别通过reflect.Type的Method获取对应的真实的方法（函数）
3. 最后对结果取其Name和Type得知具体的方法名
4. 也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
5. struct 或者 struct 的嵌套都是一样的判断处理方式



##### 通过reflect.Value设置实际变量的值



reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，即：要修改反射类型的对象就一定要保证其值是“addressable”的。



示例如下：



```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    var num float64 = 1.2345
    fmt.Println("old value of pointer:", num)

    // 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
    pointer := reflect.ValueOf(&num)
    newValue := pointer.Elem()

    fmt.Println("type of pointer:", newValue.Type())
    fmt.Println("settability of pointer:", newValue.CanSet())

    // 重新赋值
    newValue.SetFloat(77)
    fmt.Println("new value of pointer:", num)

    ////////////////////
    // 如果reflect.ValueOf的参数不是指针，会如何？
    pointer = reflect.ValueOf(num)
    //newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}

运行结果：
old value of pointer: 1.2345
type of pointer: float64
settability of pointer: true
new value of pointer: 77
```



说明



1. 需要传入的参数是* float64这个指针，然后可以通过pointer.Elem()去获取所指向的Value，**注意一定要是指针**。
2. 如果传入的参数不是指针，而是变量，那么 

- - 通过Elem获取原始值对应的对象则直接panic
  - 通过CanSet方法查询是否可以设置返回false

1. newValue.CantSet()表示是否可以重新设置其值，如果输出的是true则可修改，否则不能修改，修改完之后再进行打印发现真的已经修改了。
2. reflect.Value.Elem() 表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的
3. 也就是说如果要修改反射类型对象，其值必须是“addressable”【对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象】
4. struct 或者 struct 的嵌套都是一样的判断处理方式



##### 通过reflect.ValueOf来进行方法的调用



这算是一个高级用法了，前面我们只说到对类型、变量的几种反射的用法，包括如何获取其值、其类型、如果重新设置新值。但是在工程应用中，另外一个常用并且属于高级的用法，就是通过reflect来进行方法【函数】的调用。比如我们要做框架工程的时候，需要可以随意扩展方法，或者说用户可以自定义方法，那么我们通过什么手段来扩展让用户能够自定义呢？关键点在于用户的自定义方法是未可知的，因此我们可以通过reflect来搞定



示例如下：



```go
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
    fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
    fmt.Println("ReflectCallFuncNoArgs")
}

// 如何通过反射来进行方法的调用？
// 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
    user := User{1, "Allen.Wu", 25}
    
    // 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
    getValue := reflect.ValueOf(user)

    // 一定要指定参数为正确的方法名
    // 2. 先看看带有参数的调用方法
    methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
    args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
    methodValue.Call(args)

    // 一定要指定参数为正确的方法名
    // 3. 再看看无参数的调用方法
    methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
    args = make([]reflect.Value, 0)
    methodValue.Call(args)
}


运行结果：
ReflectCallFuncHasArgs name:  wudebao , age: 30 and origal User.Name: Allen.Wu
ReflectCallFuncNoArgs
```



说明



1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
2. reflect.Value.MethodByName这.MethodByName，需要指定准确真实的方法名字，如果错误将直接panic，MethodByName返回一个函数值对应的reflect.Value方法的名字。
3. []reflect.Value，这个是最终需要调用的方法的参数，可以没有或者一个或者多个，根据实际参数来定。
4. reflect.Value的 Call 这个方法，这个方法将最终调用真实的方法，参数务必保持一致，如果reflect.Value'Kind不是一个方法，那么将直接panic。
5. 本来可以用u.ReflectCallFuncXXX直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调用methodValue.Call



#### Golang的反射reflect性能



Golang的反射很慢，这个和它的API设计有关。在 java 里面，我们一般使用反射都是这样来弄的。



```csharp
Field field = clazz.getField("hello");
field.get(obj1);
field.get(obj2);
```



这个取得的反射对象类型是 java.lang.reflect.Field。它是可以复用的。只要传入不同的obj，就可以取得这个obj上对应的 field。



但是Golang的反射不是这样设计的:



```go
type_ := reflect.TypeOf(obj)
field, _ := type_.FieldByName("hello")
```



这里取出来的 field 对象是 reflect.StructField 类型，但是它没有办法用来取得对应对象上的值。如果要取值，得用另外一套对object，而不是type的反射



```go
type_ := reflect.ValueOf(obj)
fieldValue := type_.FieldByName("hello")
```



这里取出来的 fieldValue 类型是 reflect.Value，它是一个具体的值，而不是一个可复用的反射对象了，每次反射都需要malloc这个reflect.Value结构体，并且还涉及到GC。



Golang reflect慢主要有两个原因

 

1. 涉及到内存分配以及后续的GC；
2. reflect实现里面有大量的枚举，也就是for循环，比如类型之类的.



#### 总结



上述详细说明了Golang的反射reflect的各种功能和用法，都附带有相应的示例，相信能够在工程应用中进行相应实践，总结一下就是：



- 反射可以大大提高程序的灵活性，使得interface{}有更大的发挥余地 

- - 反射必须结合interface才玩得转
  - 变量的type要是concrete type的（也就是interface变量）才有反射一说

- 反射可以将“接口类型变量”转换为“反射类型对象” 

- - 反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息

- 反射可以将“反射类型对象”转换为“接口类型变量 

- - reflect.value.Interface().(已知的类型)
  - 遍历reflect.Type的Field获取其Field

- 反射可以修改反射类型对象，但是其值必须是“addressable” 

- - 想要利用反射修改对象状态，前提是 interface.data 是 settable,即 pointer-interface

- 通过反射可以“动态”调用方法
- 因为Golang本身不支持模板，因此在以往需要使用模板的场景下往往就需要使用反射(reflect)来实现



#### 反射的基本原理



![img](./images/1650529719193-8d81ac41-106c-40a5-b6a1-e54bbc9748eb.png)





## **结构体标签**

实例代码：

```go
import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, " doc: ", tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
```



应用实例：

```go
package main

import (
    "encoding/json"
    "fmt"
)
// 导包

type Movie struct {
    Title  string   `json:"title"`
    Year   int      `json:"year"`
    Actors []string `json:"actors"`
}
// 创建标签结构体

func main() {
    movie := Movie{"喜剧之王", 2000, []string{"xingye", "zhangbozhi"}}
    
    // 将结构体转换为JSON格式
    jsonStr, err := json.Marshal(movie)
    if err != nil {
        fmt.Println("json marshal error", err)
        return
    }
    // err用于判断转json是否成功
    
    fmt.Printf("JSON格式：%s\n", jsonStr)
    // 输出json格式
    
    
    // 解析json字符串 jsonStr ===> 结构体实例
    // jsonStr = {"title":"哪吒之魔童降世","year":2020,"rmb":10,"actors":["xinye","zhangbozhi"]}
    myMovie := Movie{}  // 创建对象
    err = json.Unmarshal(jsonStr, &myMovie)  // 反向转换
    if err != nil {
        fmt.Println("json unmarshal error ", err)
        return
    }
    // 失败判断
    
     fmt.Printf("结构体：%v\n", myMovie)
}
```



应用示例2

```go
package main

import (
    "fmt"
    "reflect"
)

type resume struct {
    Name string `json:"name" doc:"我的名字"`
}

func findDoc(stru interface{}) map[string]string {
    t := reflect.TypeOf(stru).Elem()
    doc := make(map[string]string)

    for i := 0; i < t.NumField(); i++ {
        doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
    }

    return doc

}

func main() {
    var stru resume
    doc := findDoc(&stru)
    fmt.Printf("name字段为：%s\n", doc["name"])
}
```
