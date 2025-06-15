package data
import (
	"fmt"
	"os"
	"encoding/json"
)
/**
	This is a test structure, make the members starting with upper case
	if this need to exported to other packages, otherwise it remains private.
**/
type Base struct {
	Id int
	Cat_desc string
}

/**
	This is a test function
**/
func HelloString() {
	fmt.Println("Hello")
}

/**

	Structures within the uses data structure definitions.

**/
type Users struct { // Structure of an array of users
	Users []User `json:"users"`
}

type User struc { // per user structre 
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"` // Has another structure on social
}








/**
	readJSONfile(fileName string)
	Function that reads a JSON file, reads specific data and updates a
	type struct interface. Similar implementation for configuration to control
	various functions 
**/
func ReadJSONFile(fileName string) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file open success!!!")
	defer jsonFile.Close()
} // End of ReadJSONFile()


