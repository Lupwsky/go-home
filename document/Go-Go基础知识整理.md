整理下 Go 语言的基础知识和 Java 中的异同点, 主要参考: [https://github.com/astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)

# 语言特性

* 自动垃圾回收机制
* 函数多返回值
* 高并发
* 反射
* 更加丰富的内置类型
* 错误异常处理

# 行分隔符

Go 语言中一行代表一个语句, 和 Java 不同, 不需要在末尾添加分号, 如果想在一行写多哥语句, 则需要使用分号分隔

# 数据类型

## 布尔型

| 类型 | 描述 |
| --- | --- |
| bool | 值仅 true 或者 false |

## 数字类型

整数类型:

| 类型 | 描述 |
| --- | --- |
| unit8 | 无符号 8 位整型 (0 到 255) |
| uint16 | 无符号 16 位整型 (0 到 65535) |
| uint32 | 无符号 32 位整型 (0 到 4294967295) |
| uint64 | 无符号 64 位整型 (0 到 18446744073709551615) |
| int8 | 有符号 8 位整型 (-128 到 127) |
| int16 | 有符号 16 位整型 (-32768 到 32767) |
| int32 | 有符号 32 位整型 (-2147483648 到 2147483647) |
| int64 | 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

浮点类型:

| 类型 | 描述 |
| --- | --- |
| float32 | 32位浮点型数 |
| float64 | 64位浮点型数 |
| complex64 | 32 位实数和虚数 |
| complex128 | 64 位实数和虚数 |

其他类型, 仅列出部分:

| 类型 | 描述 |
| --- | --- |
| byte | 字节类型 |
| uintptr | 无符号整型, 用于存放一个指针 |

## 字符串类型

| 类型 | 描述 |
| --- | --- |
| string | Go 的字符串由单个字节连接起来的, 字节使用 UTF-8 编码标识 Unicode 文本, 可以使用双引号("")或者反引号(``)来创建, 区别是双引号之间的转义符会被转义, 而反引号之间的字符保持不变 |

## 派生类型

包括指针类型 (Pointer), 数组类型, 接口类型, 切片类型(slice), Channel 类型, 函数类型, Map, 范围类型 (Range) 等

# 变量声明

注意全局变量, 局部变量和函数形参的区别

1. 一般格式: `var 变量名 变量类型 = 变量值`
2. 变量类型自动推到: `var 变量名 = 变量值`
3. 省略关键字 var: `变量名 := 变量值`

```go
package main
import "fmt"

func main() {
    var username string = "lpw"
    fmt.Println(username)
}
```

一次定义多个变量:

```go
package main
import "fmt"

func main() {
    // 方式一
    var username1, password1 = "lpw", "123"
    fmt.Println(username1)
    fmt.Println(password1)
    
    // 方式二
    username2, password2 := "lpw", "123"
    fmt.Println(username2)
    fmt.Println(password2)
     
    // 方式三
    var (
        username3 string = "lpw"
        password3 string = "123"
    )
    fmt.Println(username3)
    fmt.Println(password3)
}
```

一次性给多个变量赋值:

```go
package main
import "fmt"

func main() {
    // 一次性给多个变量赋值实现数据的交换, 在 Java 中需要借助一个中间变量来实现
    var username, password = "lpw", "123"
    username, password = password, username
    fmt.Println(username)
    fmt.Println(password)
}
```

# 常量声明

注意全局常量和局部常量的区别, 被定义为常量的变量其值不能被修改, 常量声明的格式: `const 变量名 [变量类型] = 变量值`

将常量枚举:

```go
const (
    GREEN = 0
    RED = 1
    BLACK = 3
)
```

注: `用作枚举时, 枚举的常量值可以是一个内置的函数`, 如 `BLACK = unsafe.Sizeof(a)`

# 运算符 & 和 *

Go 的运算符和 Java 多数是一样的, 但是 Go 提供了 `&` 和 `*` 两个运算符分别用于获取变量地址和取一个指针指向的值, 这和 C 语言是一样的

```go
package main
import "fmt"

func main() {
    // 定义一个变量
    var a int = 100

    // 定义一个 int 类型的指针
    var point *int

    // 使用指针类型保存指向变量的地址
    point = &a

    // 输出 a 的值, 这里为 100
    println("a = ", a);

    // 输出 point 指针指向的值, 使用 * 运算符取指针指向的值, 这里为 100
    println("*point = ", *point);

    // 输出变量 a 的地址值
    println("point = ", point);
}
```

# 条件语句

Go 语言提供的条件语句除了和 Java 中提供的 if else 常见的条件语句外, 还提供了 select 语句, select 语句类似于 switch 语句, 但是select会随机执行一个可运行的case, 如果没有case可运行, 它将阻塞, 直到有case可运行

参考资料:

* [Go 语言 select 语句](http://www.runoob.com/go/go-select-statement.html)

# 循环语句

* `for init; condition; post {}`: 普通 for 循环语句
* `for condition {}`: 如实现 while true 效果一样的语句 `for true {}`
* `for {}`: 死循环语句, 等价于 Java 中的 for(;;)

# 函数定义和返回多个值

函数定义的格式: `func 函数名(参数列表) 返回值类型 {}`, 返回多个值时 `func 函数名(参数列表) (返回值类型列表) {}`

```go
package main
import "fmt"

func main() {
    a, b := swap("a", "b")
    fmt.Println(a, b)
}

func swap(x string, y string) (string, string) {
    return y, x
}
```

函数除了可以返回多个值还可以作为值使用:

```go
package main
import "fmt"

func main() {
    funcSwap := func(x string, y string) (string, string){
        return y, x
    }
    fmt.Println(funcSwap("a", "b"))
}
```

参考资料:

* [Golang 函数闭包的理解](https://blog.csdn.net/netdxy/article/details/72054431)

# 切片和切片定义

Go 数组长度不可变, 因此 Go 提供了一种内置的类型 slice 类型, 切片, 可以认为是一个动态数组, 可以追加元素, 和 Java 的 Map 一样会有扩容机制, 下例包含切片的定义, append, copy 方法

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")

	// 1. 数组方式
	// 定义未初始化的数组创建
	var slice1 []string
	if slice1 == nil {
		fmt.Println("切片为空, len = " + strconv.Itoa(len(slice1)) + ", cap = " + strconv.Itoa(cap(slice1)))
	}

	// 数组初始化的时候创建, 这种方式创建时 `初始化元素的数量 = len = capacity`
	var slice2 = [] string {"0", "1", "2"}
	fmt.Println(slice2)

	// 通过子数组创建, 关于 array[startIndex:endIndex], 得到一个子数组
	// startIndex 和 endIndex 可以缺省
	// array[:] 包含全部元素的数组
	// array[startIndex:] 包含下标 startIndex 开始到最后一个元素的数组
	// array[:endIndex] 包含第一个元素到下标为 endIndex 的数组
	var array = [6] string {"0", "1", "2", "3", "4", "5"}
	var slice3 = array[0:3];
	fmt.Println(slice3)

	// 2. 使用 make([]T, length, capacity) 方法创建
	var slice4 = make([]string, 3)
	slice4[0] = "0"
	slice4[1] = "1"
	slice4[2] = "2"
	fmt.Println(slice4)

	// 3. append 方法追加元素, 如果 cap 值不够则会扩容
	fmt.Println("添加元素前, len = " + strconv.Itoa(len(slice4)) + ", cap = " + strconv.Itoa(cap(slice4)))
	slice4 = append(slice4, "3")
	fmt.Println(slice4)
	fmt.Println("添加元素后, len = " + strconv.Itoa(len(slice4)) + ", cap = " + strconv.Itoa(cap(slice4)))

	// 4. copy(dst []T, src []T), 将切片 src 复制到切片 dst, 但是要保证 dst 切片的 cap 不能小于 src 切片的 cap
	slice5 := make([]string, len(slice4), cap(slice4) * 2)
	copy(slice5, slice4)
	fmt.Println(slice5)
}
```