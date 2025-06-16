package utils

import "fmt"

func Utilfunc1(){
	fmt.Println("goutil1.go file :")
	fmt.Println("function Utilfunc1()")
	utilSupport()
} 


func utilSupport() {
	fmt.Println("Calling utilSupport()")
	fmt.Println("Internal call")
}
