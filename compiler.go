/*
	this is the compiler
	for the Nwlag programming language.
*/

package main

import (
	"encoding/json"
	"strconv"
	"strings"
)

// commands contains commands to
// execute later.
// It is of type Mixed because Mixed
// can be of type Expression or Statement
var commands = make([]*Mixed, 0)

var availableStatements = []string{
	"print",
}
var availableExpressions = []string{
	"+",
	"-",
	"*",
	"%",
	"/",
}

// Compile takes in code, and returns JSON
// bytes and an error if there was one
func Compile(in string) ([]byte, error) {
	var cmd *Mixed
	// iterate through every line in the code
	for _, line := range strings.Split(in, "\n") {
		// Divide the line into sections
		// for easier processing
		tokens := strings.Split(line, " ")
		for _, a := range availableStatements {
			for _, b := range availableExpressions {
				if tokens[1] == b {
					// It is an expression, so
					// calculate the value of
					// the expression
					var val int
					num1, err := strconv.Atoi(tokens[0])
					if err != nil {
						return nil, err
					}
					num2, err := strconv.Atoi(tokens[2])
					if err != nil {
						return nil, err
					}
					args := []int{num1, num2}
					switch tokens[1] {
					case "+":
						val = args[0] + args[1]
					case "*":
						val = args[0] * args[1]
					case "-":
						val = args[0] - args[1]
					case "/":
						val = args[0] / args[1]
					case "%":
						val = args[0] % args[1]
					}
					cmd = &Mixed{
						IsExpression: true,
						IsStatement:  false,
						Abs: &Expression{
							Name: b,
							// for some reason go doesn't
							// let me put args directly
							Arguments: []interface{}{args},
							Val:       val,
						}}
				} else if tokens[0] == a {
					// argument is a statement, don't bother
					// to evaluate the value
					arguments := strings.Join(tokens[1:], " ")
					cmd = &Mixed{
						IsStatement:  true,
						IsExpression: false,
						Abs: &Statement{
							Name:      a,
							Arguments: []interface{}{arguments},
						}}
				}
			}
		}
		commands = append(commands, cmd)
	}
	gen, err := json.MarshalIndent(commands, "", "    ")
	// empty the commands
	commands = make([]*Mixed, 0)
	if err != nil {
		return nil, err
	}
	return gen, nil
}
