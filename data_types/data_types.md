## 数据类型
### 定义方式

```go
// A
// var firstName, lastName string
// var age int
```

```go
// B
// var (
// 	firstName = "John"
// 	lastName  = "Doe"
// 	age       = 23
// )
```

```go
// C
// var (
// 	firstName, lastName, age = "John", "Doe", 34
// )
```

```go
// D
//	firstName, lastName := "John", "Doe"
//	age := 23
```

### string

```go
// var firstName, lastName string
// var age int
```

### integer

一般来说，定义整数类型的关键字是 int。 但 Go 还提供了 int8、int16、int32 和 int64 类型，其大小分别为 8、16、32 或 64 位的整数。 使用 32 位操作系统时，如果只是使用 int，则大小通常为 32 位。 在 64 位系统上，int 大小通常为 64 位。 但是，此行为可能因计算机而不同。 可以使用 uint。 但是，只有在出于某种原因需要将值表示为无符号数字的情况下，才使用此类型。 此外，Go 还提供 uint8、uint16、uint32 和 uint64 类型。

```go
var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807
fmt.Println(integer8, integer16, integer32, integer64)
```

整数类型的范围是根据其位数和有符号（signed）或无符号（unsigned）属性来确定的。对于有符号整数类型，它的范围是从最小负数到最大正数的取值范围。

在 Go 语言中，整数类型的范围如下：

int8：有符号 8 位整数，范围从 -128 到 127。
int16：有符号 16 位整数，范围从 -32768 到 32767。
int32（也被称为 rune）：有符号 32 位整数，范围从 -2147483648 到 2147483647。
int64：有符号 64 位整数，范围从 -9223372036854775808 到 9223372036854775807。
uint8（也被称为 byte）：无符号 8 位整数，范围从 0 到 255。
uint16：无符号 16 位整数，范围从 0 到 65535。
uint32：无符号 32 位整数，范围从 0 到 4294967295。
uint64：无符号 64 位整数，范围从 0 到 18446744073709551615。
范围的计算是基于整数类型在内存中使用的位数。有符号整数的最高位（最左侧的位）用于表示符号，而无符号整数则使用所有位来表示非负数。

例如，对于 int16 类型，它是一个有符号的 16 位整数，其中 1 位用于表示符号，因此剩下的 15 位用于表示值。这导致 int16 类型的范围为 -32768 到 32767。

大小溢出时就会有类似下面的报错,这里就是一个int16赋值了一个超过32767的值
cannot use 32768 (untyped int constant) as int16 value in variable declaration (overflows)

var integer16 int16 = 127
var integer32 int32 = 32767
fmt.Println(integer16 + integer32)

# command-line-arguments
.\main.go:48:14: invalid operation: integer16 + integer32 (mismatched types int16 and int32)
fmt.Println(integer16 + integer32)



### boolean