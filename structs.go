package main

import "fmt"

// type is used to declare a new data type. Here that datatype is a struct of type stack containing an int and an array of int of vals 5
type stack struct {
	stack_top int
	vals  [5]int
}

// push function which returns a pointer to stack datatype and takes in an integer parameter
func (s *stack) push(k int) {
	s.vals[s.stack_top] = k
	s.stack_top++
}

// pop funciton which returns int and takes in a pointer to stack datatype as a parameter
func (s *stack) pop() int {
	s.stack_top--
	return s.vals[s.stack_top]
}

func main() {
	
	var s *stack = new(stack)  // s is pointer to stack data type
	s.push(443)
	s.push(80)
	s.push(22)
	s.push(1023)

	fmt.Printf("first element of stack : %v\n", s.vals[0])
	fmt.Printf("stack: %v\n", s.vals)
	fmt.Printf("%v\n", s.stack_top)

	// print the entire stack data type
	fmt.Printf("stack overall : %v\n", *s)
}