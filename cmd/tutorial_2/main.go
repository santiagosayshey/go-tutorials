package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// === Integer Types ===
	var intNum int16 = 32767
	fmt.Println(intNum)
	// intNum = intNum + 1 // overflows: int16 wraps 32767 → -32768. Compiler can't catch this since the value isn't known until runtime.
	// var intNumOverflow int16 = 32767 + 1 // compiler catches this because the overflow is in a constant expression

	// === Integer Division ===
	var a int = 3
	var b int = 2
	fmt.Println(a / b) // 1 — integer division truncates, no rounding
	fmt.Println(a % b) // 1 — modulo gives the remainder

	// === Floats ===
	var floatNum float32 = 12345678.9
	fmt.Println(floatNum) // precision loss: float32 only has ~7 significant digits

	// === Type Casting ===
	var floatVal float32 = 10.1
	var intVal int32 = 1
	var result = floatVal + float32(intVal) // must explicitly cast to do mixed-type arithmetic
	fmt.Println(result)

	// === Strings ===
	var escaped string = "My \nString" // escape sequences work in double quotes
	var raw string = `My\nString`      // backticks are raw — no escape processing, supports multiline
	var concat string = "My" + " " + "String"
	fmt.Println(escaped)
	fmt.Println(raw)
	fmt.Println(concat)

	// === Bytes vs Runes ===
	fmt.Println(len("γ"))                    // 2 — len() counts bytes, and γ is 2 bytes in UTF-8
	fmt.Println(utf8.RuneCountInString("γ")) // 1 — counts actual characters (Unicode code points)

	var myRune rune = 'a'      // single quotes = rune (a single Unicode code point, alias for int32)
	fmt.Println(myRune)        // 97 — prints the numeric code point
	fmt.Printf("%c\n", myRune) // a  — %c prints the character

	// === Booleans ===
	var myBool bool = true
	fmt.Println(myBool)
	// no truthy/falsy in Go — conditions must be explicit: if x != 0, not if x

	// === Zero Values ===
	var zeroInt int    // 0
	var zeroStr string // ""
	var zeroBool bool  // false
	fmt.Println(zeroInt, zeroStr, zeroBool)

	// === Variable Declaration ===
	var explicit string = "typed" // explicit type
	var inferred = "inferred"     // type inferred from value
	shorthand := "shorthand"      // := declares and assigns, most common in practice
	fmt.Println(explicit, inferred, shorthand)

	// := only for new variables, = for reassignment
	shorthand = "reassigned"
	fmt.Println(shorthand)

	// multiple assignment
	x, y := 1, 2
	fmt.Println(x, y)

	// === Constants ===
	const pi float64 = 3.14159
	const greeting = "hello" // type inferred for constants too
	// constants must be assigned at declaration — no default zero value
	// constants can't use :=
}
