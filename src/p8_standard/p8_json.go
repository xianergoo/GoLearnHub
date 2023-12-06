package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Age       int
	BirthDate time.Time `json:"birthdate"`
}

func main() {
	person := Person{
		Name:      "John Doe",
		Age:       30,
		BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	type customTime struct {
		time.Time
	}

	customFormat := "2006-01-02"
	jsonData, err := json.MarshalIndent(person, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonData = bytes.Replace(jsonData, []byte(`"birthdate"`), []byte(`"birthdate":"`+person.BirthDate.Format(customFormat)+`"`), 1)

	fmt.Println(string(jsonData))
}
