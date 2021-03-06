package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	s := make([]int, 2)
	
	//sets initial two values, required for first addition
	s[0] = 0
	s[1] = 1
	
	//takes the last two numbers and adds them
	return func() int {
		fibReturn := s[len(s)-1]
		s = append(s, s[len(s)-1] + s[len(s)-2])
		return fibReturn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
