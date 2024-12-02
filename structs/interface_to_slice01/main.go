package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Custom struct {
	Field1 uint64 `json:"Field1"`
	Field2 byte   `json:"Field2"`
	Field3 string `json:"Field3"`
	Field4 bool   `json:"Field4"`
}

type Data struct {
	Field1 string      `json:"Field1"`
	Field2 int         `json:"Field2"`
	Field3 interface{} `json:"Field3"`
}

func main() {
	input := []Custom{
		{
			Field1: 123456789,
			Field2: 234,
			Field3: "Custom string",
		},
		{
			Field1: 23456789,
			Field2: 120,
			Field3: "Another custom string",
		},
		{
			Field1: 3456789,
			Field2: 100,
			Field3: "ANOTHER CUSTOM STRING",
		},
	}

	data := Data{
		Field1: "A test string",
		Field2: 1234,
		Field3: input,
	}

	jStr, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("Failed to marshal data: %v\n", err)

		os.Exit(1)
	}

	fmt.Printf("data to JSON str: %s\n", jStr)

	if output, ok := data.Field3.([]Custom); ok {
		fmt.Println("Successfully converted data.Field3 to type `[]Custom`!")

		fmt.Print("Entries:\n")

		for k, v := range output {
			fmt.Printf("\tEntry #%d", k+1)

			fmt.Printf("\t\tField1 => %d\n", v.Field1)
			fmt.Printf("\t\tField2 => %d\n", v.Field2)
			fmt.Printf("\t\tField3 => %s\n", v.Field3)

		}

	} else {
		fmt.Println("Failed to convert data.Field3 to type `[]Custom`.")
	}
}
