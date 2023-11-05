package main

import "fmt"

func fibonacci(n int) []int {

	if n < 2 {
		return make([]int, 0)
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			res[i] = 1
		} else {
			res[i] = res[i-1] + res[i-2]
		}

	}
	return res
}

func romanConvert(s string) int {
	roman_maps := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	// fmt.Printf("roman_maps: %v\n", roman_maps)
	arabicVals := make([]int, len(s)+1)

	for index, digit := range s {
		if val := roman_maps[string(digit)]; val > 0 {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", s, digit)
			return 0
		}
	}

	result := 0
	for i := 0; i < len(arabicVals)-1; i++ {
		if arabicVals[i] < arabicVals[i+1] {
			arabicVals[i] = -arabicVals[i]
		}
		result += arabicVals[i]
	}
	return result

}

func main() {
	x := 6
	s := fibonacci(x)
	fmt.Printf("s: %v\n", s)
	// fmt.Printf("s: %v\n", s)

	//convert roman to int
	roman_number := ""
	fmt.Println("please input a roman number: ")
	fmt.Scanln(&roman_number)
	roman_number2 := romanConvert(roman_number)
	fmt.Printf("roman_number2: %v\n", roman_number2)

}
