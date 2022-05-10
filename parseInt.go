package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := "100"

	result, error1 := strconv.ParseInt(str1, 0, 8)

	if error1 == nil {
		fmt.Println("Parsed values is:", result)
	} else {
		fmt.Println("The string could not be parsed to int8")
	}

	//NOTE - PARSING STRING VALUE GREATER THAN THE BITSIZE SPECIFIED RESULT IN OVERFLOW AND OUT OF RANGE ERROR
}
