package main

import (
	"errors"
	"fmt"
)

func prompt(prompt string) string {
	var answer string
	fmt.Print(prompt)
	fmt.Scanln(&answer)
	return answer
}

func validate(name string) error {
	if name != "kimberly" {
		return errors.New("you are not kimberly!")
	} else {
		return nil
	}
}

func cleanup() {
	fmt.Println("Goodbye!")
}

func final_cleanup() {
	fmt.Println("ok really goodbye!")
}

func main() {
	name := prompt("please type your name: ")
	defer final_cleanup()
	defer cleanup()
	if err := validate(name); err != nil {
		fmt.Println("Uh oh", err)
		return
	}
	fmt.Println("hello", name)
}
