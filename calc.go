package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/mdwhatcott/calc-lib/calc"
)

type Calculator interface{ Calculate(a, b int) int }

func main() {
	var (
		inputs     []string   = os.Args[1:]
		calculator Calculator = calc.Addition{}
		output     io.Writer  = os.Stdout
	)

	if len(inputs) != 2 {
		panic("usage: <a> <b>")
	}
	a, err := strconv.Atoi(inputs[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(inputs[1])
	if err != nil {
		panic(err)
	}
	c := calculator.Calculate(a, b)
	_, err = fmt.Fprintln(output, c)
	if err != nil {
		panic(err)
	}
}
