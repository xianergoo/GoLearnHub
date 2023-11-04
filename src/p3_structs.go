package main

import (
	"encoding/json"
	"fmt"
)

/*
type Employee struct {
	ID        int
	FirstName string
	Lastname  string
	Address   string
}

func main() {
	emp := Employee{}
	fmt.Printf("emp: %v\n", emp)

	var john Employee
	john = Employee{1001, "john", "stone", "winterfall"}
	john2 := Employee{1001, "john", "stone", "winterfall"}
	fmt.Printf("john: %v\n", john)
	fmt.Printf("john2: %v\n", john2)
	john3 := Employee{FirstName: "john", Lastname: "stone"}
	fmt.Printf("john3: %v\n", john3)

	john4 := &john3
	fmt.Printf("john4: %v\n", john4)
	john4.FirstName = "john4"
	fmt.Printf("john3: %v\n", john3)
	fmt.Printf("john4: %v\n", john4)

}
*/

// 结构嵌入

type Person struct {
	ID        int
	FirstName string
	Lastname  string
	Address   string
	gender    int
}

type Employee struct {
	Person
	WorkerID int
}

type Contractor struct {
	Person
	CompanyID int
}

type QC struct {
	Person
	workingflow string `A1:flow`
	WorkerID    int
	qcType      int `1`
}

func main() {

	employee := Employee{
		Person: Person{
			FirstName: "John",
			Lastname:  "Snow",
		},
	}
	employee.gender = 1
	fmt.Printf("employee: %v\n", employee)

	data, _ := json.Marshal(employee)
	fmt.Printf("data: %s\n", data)

	var decode Employee
	json.Unmarshal(data, &decode)
	fmt.Printf("decode: %v\n", decode)

	qc := QC{
		Person: Person{
			FirstName: "John",
			Lastname:  "Snow",
		},
	}
	fmt.Printf("qc: %v\n", qc)

}
