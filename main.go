package main

import (
	"fmt"
	l "github.com/willowtreeapps/gh-issue/log"
)

// UselessImplementation is useless
type UselessImplementation struct{}

// FunctionThatLogsSomething is an implementation
func (*UselessImplementation) FunctionThatLogsSomething() {
	fmt.Printf("Log This Implementation")
}

func main() {
	var logger l.UselessLogger
	logger = &UselessImplementation{}
	logger.FunctionThatLogsSomething()
}
