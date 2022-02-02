package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OpFuncs func(i, j int) int

func add(i, j int) int       { return i + j }
func substract(i, j int) int { return i - j }
func multiply(i, j int) int  { return i * j }

func divide(i, j int) int {
	if j == 0 {
		fmt.Println("0 is not valid value for divider")
		return 0
	}
	return i / j
}

var Operators = map[string]OpFuncs{
	"+": add,
	"-": substract,
	"/": divide,
	"*": multiply,
}

func getNum(reader *bufio.Reader) (int, error) {
	fmt.Printf("Enter number: ")
	stVal, _ := reader.ReadString('\n')
	val, err := strconv.Atoi(strings.TrimSpace(stVal))
	if err != nil {
		return 0, errors.New("This is not valid number")
	}
	return val, nil

}
func getOp(reader *bufio.Reader) (OpFuncs, error) {
	fmt.Printf("Choose operator from (+ - * /) : ")
	op, _ := reader.ReadString('\n')
	opFun, ok := Operators[strings.TrimSpace(op)]
	if !ok {
		return nil, errors.New("Invalid operator")
	}
	return opFun, nil
}

func display(reader *bufio.Reader) {
	firstVal, err := getNum(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	secondVal, err := getNum(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	opFunc, err := getOp(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := opFunc(firstVal, secondVal)
	fmt.Println("Result: ", result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	display(reader)
}
