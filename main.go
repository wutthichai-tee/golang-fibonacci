package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Begin")


	f := findFibonanci()
	fmt.Println(f(), f(), f(), f() , f())


	fmt.Println("End")
}


func findFibonanci() func() int {
	/*
		0 (before)
		1 (next)
		1 = 0 (before) + 1 (next), set before = 1 (replace before with next), set next = 1 (0 + 1)
		2 = 1 (before) + 1 (next), set before = 1 (replace before with next), set next = 2 (1 + 1)
		3 = 1 (before) + 2 (next), set before = 2 (replace before with next), set next = 3 (1 + 2)
		5 = 2 (before) + 3 (next), set before = 3 (replace before with next), set next = 5 (2 + 3)
		8 = 3 (before) + 5 (next), set before = 5 (replace before with next), set next = 8 (3 + 5)
	*/
	before := 0
	next := 1
	return func() int {
		sumResult := before + next
		before = next
		next = sumResult
		return sumResult
	}
}