/*
	This contains essential data
	types for Nwlag
*/

package main

// Function not implemented yet,
// but there for completedness
// type Function struct {
// 	Name      string
// 	Arguments []interface{}
// 	// A function can contain
// 	// both expressions and
// 	// statements
// 	Code []Mixed
// }

// Mixed can be either Statement
// or Expression
type Mixed struct {
	IsStatement  bool
	IsExpression bool
	Abs          interface{}
}

// Statement represents
// statements. How a statement
// is used is entirely up to
// the interpreter.
type Statement struct {
	Name      string
	Arguments []interface{}
}

// Since an Expression
// is calculated beforehand,
// store the value
type Expression struct {
	Name      string
	Arguments []interface{}
	Val       interface{}
}
