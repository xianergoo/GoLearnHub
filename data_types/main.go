package main

import (
	"fmt"
	"math"
	"strconv"
)

// Basic type

// ## string
// A
// var firstName, lastName string
// var age int

// B
// var (
// 	firstName = "John"
// 	lastName  = "Doe"
// 	age       = 23
// )

// C
// var (
// 	firstName, lastName, age = "John", "Doe", 34
// )

// ## integer
var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807

// const type
const HTTPStatusOK = 200

const (
	StatusOK              = 0
	StatusConnectionReset = 1
	StatusOtherErro       = 2
)

func main() {
	// firstName = string('Z')
	// lastName = string('3')

	firstName, lastName := "John", "Doe"
	age := 23
	fmt.Printf("my name is %s.%s, age %d\n", firstName, lastName, age)

	fmt.Println(integer8, integer16, integer32, integer64)
	fmt.Println(integer16 + int16(integer32)) //强制转换类型

	arune := 'G'
	fmt.Println(arune)

	// rune := 'G'
	// fmt.Printf('%c \n', rune)
	// 修正后的代码使用双引号 " 包裹 %c\n，而不是单引号 '。在 fmt.Printf 中，%c 是用于格式化字符的占位符，需要使用双引号。

	fmt.Printf("%c\n", arune)
	// 	%c 是 Go 语言中用于格式化字符的占位符。当你使用 fmt.Printf 或 fmt.Sprintf 函数时，可以使用 %c 来表示一个字符。
	// 在 Go 中，字符是使用 Unicode 编码表示的。Unicode 是一种字符集，为世界上几乎所有的字符都分配了一个唯一的数字标识符，称为码点（code point）。每个字符都有一个对应的整数值，这个整数值就是码点。而 %c 占位符可以将一个码点转换为对应的字符。
	// 例如，rune := 'G' 将字符 'G' 赋值给 rune 变量。当你使用 fmt.Printf 和 %c 来格式化输出时，%c 会将码点 71 转换为对应的字符 'G'，并将字符输出。

	var integer32 int = 2147483648
	fmt.Println(integer32)

	// var integer uint = -10
	//无符号整数类型（uint）只能表示非负数，其范围从 0 到最大正整数。因此，无法使用负数来初始化无符号整数类型的变量。

	//如果你想要存储一个有符号的整数，可以使用有符号整数类型（int）或其他适当的有符号整数类型。
	var integer int = -10
	fmt.Println(integer)

	var float32 float32 = 2147483647
	var float64 float64 = 9223372036854775807
	fmt.Println(float32, float64)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	const e = 2.71828
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	var featureFlag bool = true
	fmt.Println(featureFlag)

	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)

	var defaultInt int
	// var defaultFloat32 float32
	// var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultBool, "_ "+defaultString+" _")

	// var integer16 int16 = 127
	// var integer32 int32 = 32767
	// fmt.Println(int32(integer16) + integer32)

	i, _ := strconv.Atoi("-42")
	// strconv.Atoi 函数，并将返回的整数值赋给变量 i。
	// 需要注意的是，我们使用了一个匿名变量 _ 来忽略 strconv.Atoi 函数的错误返回值。在这种情况下，我们假设字符串转换为整数的操作不会出错，因此我们忽略了可能出现的错误。
	s := strconv.Itoa(-42)
	// strconv.Itoa 函数。该函数接受一个整数作为参数，并将其转换为对应的字符串表示
	// 注意，strconv.Itoa 函数只能用于将整数转换为字符串。如果你想要将其他类型的值转换为字符串，可以使用 fmt.Sprintf 函数或字符串拼接等方法。
	fmt.Println(i, s)

}
