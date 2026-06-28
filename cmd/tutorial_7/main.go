package main

import "fmt"

func main() {
	var i int = 5
	var p *int = &i

	fmt.Printf("The value is %v at address %v\n", *p, p)

	intSlice := []int32{1, 2, 3}
	intSlice2 := intSlice
	// this also affects intSlice because slices are pointers themselves.
	intSlice2[0] = 10
	fmt.Println(intSlice)
	fmt.Println(intSlice2)

	var num = 2
	// num becomes 4
	square(&num)
	fmt.Println(num)
	// num stays as 4 since we passed in a copy
	squarePure(num)
	fmt.Println(num)
	// need to resave to get it to work
	num = squarePure(num)
	fmt.Println(num)

}

// pointers are useful when you want to call behavior on some variable that permanently changes it
// no need to return the new value and set it again
// pointers are also useful when you have large variables and don't want to waste memory creating another copy

func square(num *int) {
	*num = *num * *num
}

// compared to

func squarePure(num int) int {
	return num * num
}
