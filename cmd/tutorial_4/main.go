package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	// arrays are contiguous in memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intSlice := []int32{4, 5, 6}
	printIntSlices(intSlice)

	// append doubles capacity when exceeding existing capacity
	intSlice = append(intSlice, 7)
	printIntSlices(intSlice)

	intSlice2 := []int32{7, 8}
	intSlice = append(intSlice, intSlice2...)
	printIntSlices(intSlice)

	intSlice = append(intSlice, 9)
	printIntSlices(intSlice)

	// make takes params type, length, capacity
	intSlice3 := make([]int32, 3, 10)
	printIntSlices(intSlice3)

	// map where keys are strings and values are uint8s
	myMap := make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	// jason doesn't exist so this returns the default value of uint8 = 0
	// ok gives us whether it exists or not
	var age, ok = myMap2["Jason"]
	if !ok {
		fmt.Println("Does not exist")
	} else {
		fmt.Printf("The age is %v", age)
	}

	// works without a pointer because maps are reference variables in the first place
	delete(myMap2, "Adam")
	fmt.Println(myMap2["Adam"])

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printIntSlices(i []int32) {
	fmt.Printf("The length is %v with capacity %v and values %v\n", len(i), cap(i), i)
}
