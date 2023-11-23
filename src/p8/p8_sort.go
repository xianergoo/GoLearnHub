package main

import (
	"fmt"
	"sort"
)

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	fmt.Println(n[i], n[j])
	n[i], n[j] = n[j], n[i]
}

type testSlice [][]int

func (l testSlice) Len() int           { return len(l) }
func (l testSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice) Less(i, j int) bool { return l[i][1] < l[j][1] }

type People struct {
	Name string
	Age  int
}

type testSlice2 []People

func (l testSlice2) Len() int           { return len(l) }
func (l testSlice2) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l testSlice2) Less(i, j int) bool { return l[i].Age < l[j].Age }

func main() {

	lp := testSlice2{
		{Name: "n1", Age: 12},
		{Name: "n2", Age: 11},
		{Name: "n3", Age: 10},
	}

	fmt.Println(lp) //[{n1 12} {n2 11} {n3 10}]
	sort.Sort(lp)
	fmt.Println(lp) //[{n3 10} {n2 11} {n1 12}]

	n := []int{1, 3, 2}
	fmt.Printf("%p\n", &n)
	sort.Ints(n)
	fmt.Printf("%p\n", &n)
	fmt.Println(n)

	li := sort.IntSlice{
		12,
		33,
		4,
		4,
		5,
		7,
	}
	fmt.Printf("li: %v\n", li)
	sort.Ints(li)
	fmt.Printf("li: %v\n", li)
	li = []int{1, 3, 4, 3, 4}
	fmt.Printf("li: %v\n", li)
	sort.Ints(li)
	fmt.Printf("li: %v\n", li)

	//字符串排序，现比较高位，相同的再比较低位
	// [] string
	ls := sort.StringSlice{
		"100",
		"42",
		"41",
		"3",
		"2",
	}
	fmt.Println(ls) //[100 42 41 3 2]
	sort.Strings(ls)
	fmt.Println(ls) //[100 2 3 41 42]

	//字符串排序，现比较高位，相同的再比较低位
	ls = sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println(ls) //[d ac c ab e]
	sort.Strings(ls)
	fmt.Println(ls) //[ab ac c d e]

	//汉字排序，依次比较byte大小
	ls = sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}
	fmt.Println(ls) //[啊 博 次 得 饿 周]
	sort.Strings(ls)
	fmt.Println(ls) //[博 周 啊 得 次 饿]

	for _, v := range ls {
		fmt.Println(v, []byte(v))
	}

	//博 [229 141 154]
	//周 [229 145 168]
	//啊 [229 149 138]
	//得 [229 190 151]
	//次 [230 172 161]
	//饿 [233 165 191]

	lt := testSlice{
		{1, 4},
		{9, 3},
		{7, 5},
	}

	fmt.Println(lt) //[[1 4] [9 3] [7 5]]
	sort.Sort(lt)
	fmt.Println(lt) //[[9 3] [1 4] [7 5]]

}
