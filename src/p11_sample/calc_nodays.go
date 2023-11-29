package main

import (
	"fmt"
)

func dayOfYear(y, m, d int) int {
	var days int
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	return days
}

func dayOfYear2(y, m, d int) int {
	isLeapYear := func(year int) bool {
		return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
	}
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLeapYear(y) {
		daysInMonth[1] = 29
	}

	day := d
	for i := 0; i < m-1; i++ {
		day += daysInMonth[i]
	}
	return day
}

func main() {

	var y, m, d int = 0, 0, 0
	var days int = 0
	fmt.Printf("请输入年月日\n")
	fmt.Scanf("%d%d%d\n", &y, &m, &d)
	fmt.Printf("%d 年 %d 月 %d 日", y, m, d)

	// days = dayOfYear(y, m, d)
	days = dayOfYear2(y, m, d)
	fmt.Printf("是今年的第 %d 天!\n", days)

}
