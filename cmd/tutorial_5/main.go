package main

import "fmt"

func main() {
	myString := []rune("résumé")
	indexed := myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	// strings are stored as uint8 arrays
	// é is 8 bytes, and therefore takes up 2 indexes
	// the value is the ascii representation of the character
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "a", "m"}
	var catStr = ""

	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)
	// strings are immutable, you can't change catStr after its been created
}
