package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    
    // Read and save an integer, double, and String to your variables.
    
    // Print the sum of both integer variables on a new line.
    
    // Print the sum of the double variables on a new line.
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.

    var intVar uint64
    var doubleVar float64
    var stringVar string
     
    scanner.Scan()
    intVar, _ = strconv.ParseUint(scanner.Text(), 10, 64)
    scanner.Scan()
    doubleVar, _ = strconv.ParseFloat(scanner.Text(), 64)
    scanner.Scan()
    stringVar = scanner.Text()
    
    fmt.Println(i + intVar)
    fmt.Println(fmt.Sprintf("%.1f", d + doubleVar))
    fmt.Println(s + stringVar) 
}
