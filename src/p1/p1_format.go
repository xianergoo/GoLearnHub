package main

import "fmt"

type WebSite struct {
	name string
}

func main() {
	/*
	   	占位符						说明						举例										输出
	   %v		相应值的默认格式。								Printf("%v", site)，Printf("%+v", site)	{duoke360}，{Name:duoke360}
	   		在打印结构体时，“加号”标记（%+v）会添加字段名
	   %#v		相应值的Go语法表示							Printf("#v", site)						main.Website{Name:"duoke360"}
	   %T		相应值的类型的Go语法表示						Printf("%T", site)						main.Website
	   %%		字面上的百分号，并非值的占位符					Printf("%%")							%
	*/

	/* site := WebSite{name: "hello"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site)
	fmt.Printf("site: %T\n", site)

	a := 100
	fmt.Printf("a: %T\n", a)
	fmt.Println("%%\n")

	b := true
	fmt.Printf("b: %t\n", b)
	fmt.Printf("b: %b\n", b) */

	/*
	   	占位符						说明						举例									输出
	   %b		二进制表示									Printf("%b", 5)						101
	   %c		相应Unicode码点所表示的字符					Printf("%c", 0x4E2D)				中
	   %d		十进制表示									Printf("%d", 0x12)					18
	   %o		八进制表示									Printf("%o", 10)					12
	   %q		单引号围绕的字符字面值，由Go语法安全地转义		Printf("%q", 0x4E2D)				'中'
	   %x		十六进制表示，字母形式为小写 a-f				Printf("%x", 13)					d
	   %X		十六进制表示，字母形式为大写 A-F				Printf("%x", 13)					D
	   %U		Unicode格式：U+1234，等同于 "U+%04X"			Printf("%U", 0x4E2D)				U+4E2D
	*/
	i := 8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)

	i = 96
	fmt.Printf("i: %c\n", i)
	fmt.Printf("i: %x\n", 100)

	x := 100
	p := &x
	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %t\n", p)
	fmt.Printf("p: %p\n", p)

}
