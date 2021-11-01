package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    /*
    To complete this challenge, you must save a line of input from stdin to a
    variable, print Hello, World. on a single line, and finally print the value
    of your variable on a second line.
    */
    reader := bufio.NewReader(os.Stdin)
    message, _ := reader.ReadString('\n')
    fmt.Printf("Hello, World.\n%s", message)
}
