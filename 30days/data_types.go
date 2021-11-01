package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
      
    /*
    Complete the code in the editor below. The variables i, d, & s are already
    declared and initialized for you. You must: Declare 3 variables: one of type
    int, one of type double, and one of type String. Read 3 lines of input from
    stdin (according to the sequence given in the Input Format section below)
    and initialize your  variables.  Use the  operator to perform the following
    operations: Print the sum of i plus your int variable on a new line. Print
    the sum of d plus your double variable to a scale of one decimal place on a
    new line.  Concatenate d with the string you read as input and print the
    result on a new line.
    */

    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    var i2 uint64
    var d2 float64
    var s2 string
     
    fmt.Scan(&i2)
    fmt.Scan(&d2)
    scanner.Scan()
    s2 = scanner.Text()
    
    fmt.Println(i + i2)
    fmt.Printf("%.1f\n", d + d2)
    fmt.Printf("%s%s", s, s2)
}
