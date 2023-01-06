package pointers

import (
	"fmt"
)


func Foo(ps *string) {
	fmt.Println("ps", ps)
	fmt.Println("*ps", *ps)
	fmt.Println("&ps", &ps)
	*ps = "hi"
	fmt.Println("ps", ps)
	fmt.Println("*ps", *ps)
	fmt.Println("&ps", &ps)
}

func main() {
	var someString string // zero value: ""
	fmt.Println("someString", someString)
	fmt.Println("&someString", &someString)
	Foo(&someString)
	fmt.Println("someString", someString)
	fmt.Println("&someString", &someString)

	var stringPtr *string = new(string) // zero value: nil
	fmt.Println("stringPtr", stringPtr)
	fmt.Println("*stringPtr", *stringPtr)
	fmt.Println("&stringPtr", &stringPtr)
	Foo(stringPtr)
	fmt.Println("stringPtr", stringPtr)
	fmt.Println("*stringPtr", *stringPtr)
	fmt.Println("&stringPtr", &stringPtr)
}
