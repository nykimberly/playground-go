package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
Given an integer, , perform the following conditional actions:
If  is odd, print Weird
If  is even and in the inclusive range of  to , print Not Weird
If  is even and in the inclusive range of  to , print Weird
If  is even and greater than , print Not Weird
Complete the stub code provided in your editor to print whether or not  is weird.
*/

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    N := int32(NTemp)
    
    if (N % 2 == 1) {
        fmt.Println("Weird")
    } else {
        if (N <= 5) {
            fmt.Println("Not Weird")
        } else if (N <= 20) {
            fmt.Println("Weird")
        } else {
            fmt.Print("Not Weird")
        }
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

