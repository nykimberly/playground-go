package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println("Would you prefer to use scan (s) or use bufio (b)?")
	var userOption string
	fmt.Scan(&userOption)
	userOptionLower := strings.ToLower(userOption)
	if userOptionLower != "s" && userOptionLower != "b" {
		fmt.Print("Invalid option.")
	} else {
		fmt.Println("What would you like me to parrot back to you?")
		if userOptionLower == "s" {
			var inputStr string
			fmt.Scanln(&inputStr)
			fmt.Println(inputStr)
		} else {
			reader := bufio.NewReader(os.Stdin)
			inputLn, _ := reader.ReadString('\n')
			fmt.Print(inputLn)
		}
	}

}
