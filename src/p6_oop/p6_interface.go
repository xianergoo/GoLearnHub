package main

import (
	"errors"
	"fmt"
)

// /* 定义接口 */
// type interface_name interface {
// 	method_name1 [return_type]
// 	method_name2 [return_type]
// 	method_name3 [return_type]
// 	...
// 	method_namen [return_type]
//  }

//  /* 定义结构体 */
//  type struct_name struct {
// 	/* variables */
//  }

//  /* 实现接口方法 */
//  func (struct_name_variable struct_name) method_name1() [return_type] {
// 	/* 方法实现 */
//  }
//  ...
//  func (struct_name_variable struct_name) method_namen() [return_type] {
// 	/* 方法实现*/
//  }

type Account struct {
	firstName string
	lastName  string
}

type EmployeeA struct {
	Account
	credits float32
}

func (a *Account) ChangeName(name string) {
	a.firstName = name
}

func (e EmployeeA) String() string {
	return fmt.Sprintf("name: %s %s\nCreits: %.2f\n", e.firstName, e.lastName, e.credits)
}

func (e *EmployeeA) AddCredits(amount float32) (float32, error) {
	if amount > 0.0 {
		e.credits += amount
		return e.credits, nil
	}
	return 0.0, errors.New("Invalid credit amount.")
}

func (e *EmployeeA) RemoveCredits(amount float32) (float32, error) {
	if amount > 0.0 {
		if amount < e.credits {
			e.credits -= amount
			return e.credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has.")
	}
	return 0.0, errors.New("Invalid credit amount.")
}

func (e EmployeeA) CheckCredits() float32 {
	return e.credits
}

func main() {
	emp1 := EmployeeA{Account{"john", "snow"}, 400}
	fmt.Printf("emp1: %v\n", emp1)
	fmt.Printf("emp1 account: %v\n", emp1.CheckCredits())
	credits, error := emp1.AddCredits(200)
	if error != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("credits: %v\n", credits)
	}
	fmt.Printf("emp1: %v\n", emp1)
	_, err := emp1.RemoveCredits(1000)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("emp1: %v\n", emp1)
	emp1.ChangeName("mike")
	fmt.Printf("emp1: %v\n", emp1)

}
