package main

import (
	"fmt"
	"DataProc/data"
)


func main() {
	var bs data.Base
	fmt.Println("Hello")
	bs.Id = 34
	bs.Cat_desc = "Emotion"
	data.ReadJSONFile("users.json")
}
