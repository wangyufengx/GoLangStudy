# fmt格式化指令

golang的fmt包格式化I/O函数的实现与类似C语言的`printf`和`scanf`类似。格式的`verbs`源于C,但是更简单。以下在路径`src/fmt/doc.go`可找到。

### 一般的指令

| 格式化指令 | 说明                              |
| ---------- | --------------------------------- |
| %v         | 该值的默认格式                    |
| %+v        | 当打印结构体时，+表示添加字段名称 |
| %#v        | 该值的go语法表示形式              |
| %T         | 该值类型                          |
| %%         | 输出%，无任何值                   |

#### 实例

```go
type People struct {
	Name string
}

func main() {
	fmt.Printf("%%\n")			//%
	t := People{Name: "golang"}
    fmt.Printf("%v\n", t)		//{golang}
    fmt.Printf("%+v\n", t)		//{Name:golang}
    fmt.Printf("%#v\n", t)		//main.People{Name:"golang"}
	fmt.Printf("%T\n", t)		//main.People
	
	//#v
	fmt.Printf("%#v\n", 1)                                    //1
	fmt.Printf("%#v\n", "hello")                              //“hello"
	fmt.Printf("%#v\n", []int{1, 2})                          //[]int{1,2}
	fmt.Printf("%#v\n", map[string]interface{}{})             //map[string]interface {}{}
	fmt.Printf("%#v\n", struct{ Hello string }{Hello: "tom"}) //struct { Hello string }{Hello:"tom"}
}
```

### 布尔类型占位符

| 格式化指令 | 说明          |
| ---------- | ------------- |
| %t         | true或者false |

### 实例

```go
fmt.Printf("%t", true)	//true
```

### Integer类型的占位符

| 格式化指令 | 说明                                              |
| ---------- | ------------------------------------------------- |
| %b         | 一个二进制整数，将一个整数格式化为二进制的表达式  |
| %c         | 一个Unicode的字符                                 |
| %d         | 十进制数值                                        |
| %o         | 八进制数值                                        |
| %O         | 以0o为前置的八进制数值                            |
| %q         | 单引号围绕的字符字面量的值，由Go语法安全地转义    |
| %x         | 小写的十六进制数值                                |
| %X         | 大写的十六进制数值                                |
| %U         | 一个Unicode表示法表示的整型码值,默认是4个数字字符 |

#### 实例

```go
fmt.Printf("%b\n", 12)     //1100
fmt.Printf("%c\n", 97)     //a
fmt.Printf("%d\n", 010)    //8
fmt.Printf("%o\n", 10)     //12
fmt.Printf("%O\n", 10)     //0o12
fmt.Printf("%q\n", 0x4E2D) //'中'
fmt.Printf("%x\n", 10)     //a
fmt.Printf("%X\n", 10)     //A
fmt.Printf("%U\n", 97)     //U+0061
```

### 浮点数和复数的组成部分

| 格式化指令 | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| %b         | 无小数部分，指数为二的幂的科学计数法，与strconv.FormatFloat的‘b’转换格式一致 |
| %e         | 科学计数法                                                   |
| %E         | 科学计数法                                                   |
| %f         | 有小数点无指数                                               |
| %F         | 有小数点无指数，注：是%f的同义词                             |
| %g         | 根据情况选择%e或%f,大指数选%e,否则选%f                       |
| %G         | 根据情况选择%E或%F,大指数选%E,否则选%F                       |
| %x         | 十六进制表示法（十进制幂为两个指数）                         |
| %X         | 大写的十六进制表示法                                         |

#### 实例

```go
fmt.Printf("%b\n", 123456.789)	//4953959590107546p-52
fmt.Printf("%e\n", 123456.789)	//1.234568e+05
fmt.Printf("%E\n", 123456.789)	//1.234568E+05
fmt.Printf("%f\n", 123456.789)	//123456.789000
fmt.Printf("%F\n", 123456.789)	//123456.789000
fmt.Printf("%g\n", 123456.789)	//123456.789
fmt.Printf("%G\n", 123456.789)	//123456.789
fmt.Printf("%x\n", 123456.789)	//0x1.e240c9fbe76c9p+16
fmt.Printf("%X\n", 123456.789)	//0x1.e240c9fbe76c9p+16
```

### 字符串与字节切片

| 格式化指令 | 说明                                   |
| ---------- | -------------------------------------- |
| %s         | 输出字符串(string类型或[]byte)         |
| %q         | 双引号围绕的字符串，由Go语法安全地转义 |
| %x         | 十六进制，小写字母，每个字节两个字符   |
| %X         | 十六进制，大写字母，每个字节两个字符   |

#### 实例

```go
str := "Golang"
str1 := []byte(str)
fmt.Printf("%s\n", str)  //Golang
fmt.Printf("%s\n", str1) //Golang
fmt.Printf("%q\n", str)  //"Golang"
fmt.Printf("%x\n", str)  //476f6c616e67
fmt.Printf("%X\n", str)  //476F6C616E67
```

### 切片

| 格式化指令 | 说明                                        |
| ---------- | ------------------------------------------- |
| %p         | 以16进制表示的切片第0个元素的地址，开头为0x |

#### 实例

```go
slices := []int{1, 2, 3, 4}
fmt.Printf("%p\n", slices) //0xc00000e3a0
fmt.Println(&slices[0])    //0xc00000e3a0
```

### 指针

| 格式化指令 | 说明                   |
| ---------- | ---------------------- |
| %p         | 以16进制表示，前缀为0x |

注：**%b, %d, %o, %x and %X这些指令也可与指针配合使用,将值格式化为整数.**

#### 实例

```go
x := 1
fmt.Printf("%p\n", &x) //0xc0000160f0
```

### %v的不同类型下的默认的格式化指令

| 类型                   | 格式化指令               |
| ---------------------- | ------------------------ |
| bool                   | %t                       |
| int、int8 etc          | %d                       |
| uint、uint8 etc        | %d(%#x如果是使用%#v打印) |
| float32、complex64 etc | %g                       |
| string                 | %s                       |
| chan                   | %p                       |
| pointer                | %p                       |

#### 实例

```go
//uint
fmt.Printf("%#v\n", int(15))  //15
fmt.Printf("%v\n", uint(15))  //15
fmt.Printf("%#v\n", uint(15)) //0xf
fmt.Printf("%#x\n", int(15))  //0xf
fmt.Printf("%#x\n", uint(15)) //0xf
//chan
ch := make(chan string, 1)
ch <- "golang"
fmt.Printf("%v\n", ch) //0xc000046060
fmt.Printf("%p\n", ch) //0xc000046060
```

**对于复合对象，使用以上规则递归的打印对象**