package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)
/*
    Given the meal price (base cost of a meal), tip percent (the percentage of
    the meal price being added as tip), and tax percent (the percentage of the
    meal price being added as tax) for a meal, find and print the meal's total
    cost. Round the result to the nearest integer.
*/

func solve(meal_cost float64, tip_percent int32, tax_percent int32) float64 {
    var tip_part float64 = float64(tip_percent) / 100.0
    var tax_part float64 = float64(tax_percent) / 100.0
    return math.Round(meal_cost * (1.0 + tip_part + tax_part))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    totalCost := solve(meal_cost, tip_percent, tax_percent)
    fmt.Println(totalCost)
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

