package main

import (
	"fmt"
	"strconv"
)

func regularCalculator(_first_numb int, _last_numb int, _operator string) string {
	var result int
	var error_message string

	if _operator == "*" {
		result = _first_numb * _last_numb
	} else if _operator == "/" {
		result = _first_numb / _last_numb
	} else if _operator == "+" {
		result = _first_numb + _last_numb
	} else if _operator == "-" {
		result = _first_numb - _last_numb
	} else {
		error_message = ":< I'm sorry still learner. Not understand the operator you mention."
		fmt.Println(error_message)
		return "error"
	}

	return strconv.Itoa(result)
}
func main() {
	var calculatorResult string

	calculatorResult = regularCalculator(20, 21, "*")

	if calculatorResult != "error" {
		fmt.Println("Calculator Result :", calculatorResult)
	}
}
