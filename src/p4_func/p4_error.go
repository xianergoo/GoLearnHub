package main

import (
	"errors"
	"fmt"
)

type EmployeeX struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {

	employee, err := getInformation(100)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("employee: %v\n", employee)
	}

}

var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*EmployeeX, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := EmployeeX{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

/* func getInformation(id int) (*EmployeeX, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(1000)
		if err == nil {
			return employee, nil
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
} */

func apiCallEmployee(id int) (*EmployeeX, error) {
	employee := EmployeeX{LastName: "Doe", FirstName: "john"}
	return &employee, nil
}
