package main

import "fmt"

//maps := make(map[string]int)

func main() {
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}

	for _, age := range studentsAge {
		fmt.Printf("age: %v\n", age)
	}

	studentsAge2 := make(map[string]int)
	studentsAge2["john"] = 32
	studentsAge2["bob"] = 31
	fmt.Println(studentsAge2)

	/* var studentsAge3 map[string]int
	studentsAge3["john"] = 32
	studentsAge3["bob"] = 31
	fmt.Println(studentsAge3)
	//panic: assignment to entry in nil map
	*/

	studentsAge4 := make(map[string]int)
	studentsAge4["john"] = 32
	studentsAge4["bob"] = 31
	fmt.Println("Christy's age is", studentsAge4["christy"])
	//Christy's age is 0

	studentsAge5 := make(map[string]int)
	studentsAge5["john"] = 32
	studentsAge5["bob"] = 31

	age, exist := studentsAge5["AAA"]
	if exist {
		fmt.Println("AAA's age is", age)
	} else {
		fmt.Println("AAA's age could not be found")
	}

	//delete item
	studentsAge6 := make(map[string]int)
	studentsAge6["john"] = 32
	studentsAge6["bob"] = 31
	delete(studentsAge6, "JOHN")
	delete(studentsAge6, "john")
	fmt.Println(studentsAge6)
	// map[bob:31]

}
