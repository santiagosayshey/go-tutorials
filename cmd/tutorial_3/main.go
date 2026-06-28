package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("--- If/Else Version ---")
	ifElseVersion(10, 2)

	fmt.Println("\n--- Switch Version ---")
	switchVersion(11, 3)

	fmt.Println("\n--- Guard Clause Version (idiomatic Go) ---")
	guardClauseVersion(10, 0)
	guardClauseVersion(10, 2)
	guardClauseVersion(11, 3)

	fmt.Println("\n--- Short-Circuit Evaluation ---")
	shortCircuitDemo()
}

// ============================================================
// Version 1: If/Else
// Works, but nesting gets ugly fast with more conditions.
// ============================================================
func ifElseVersion(numerator, denominator int) {
	quotient, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
	} else {
		if remainder == 0 {
			fmt.Printf("The quotient is %v\n", quotient)
		} else {
			fmt.Printf("The quotient is %v with remainder %v\n", quotient, remainder)
		}
	}
}

// ============================================================
// Version 2: Switch
// Cleaner than nested if/else, but overkill for simple conditions.
// Shines when you have many distinct cases.
// ============================================================
func switchVersion(numerator, denominator int) {
	quotient, remainder, err := intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Printf("The quotient is %v\n", quotient)
	default:
		fmt.Printf("The quotient is %v with remainder %v\n", quotient, remainder)
	}
}

// ============================================================
// Version 3: Guard Clauses (idiomatic Go)
// Handle errors first, return early, keep the happy path flat.
// This is what you'll see in real Go codebases.
// ============================================================
func guardClauseVersion(numerator, denominator int) {
	quotient, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		return
	}
	if remainder == 0 {
		fmt.Printf("The quotient is %v\n", quotient)
		return
	}
	fmt.Printf("The quotient is %v with remainder %v\n", quotient, remainder)
}

// ============================================================
// Short-Circuit Evaluation: && and ||
//
// Go evaluates left to right and stops as soon as it knows the answer.
//
// &&  (AND): if the first condition is false, the second is NEVER checked
//
//	false && anything = false — so why bother?
//
// ||  (OR):  if the first condition is true, the second is NEVER checked
//
//	true || anything = true — so why bother?
//
// ============================================================
func shortCircuitDemo() {
	// && example: prevents a division by zero crash
	// if x is 0, Go never evaluates 10/x — it already knows the result is false
	x := 0
	if x != 0 && (10/x > 2) {
		fmt.Println("this never prints")
	} else {
		fmt.Println("&& short-circuited: 10/x was never evaluated (would have crashed!)")
	}

	// || example: skips the function call entirely
	// len(name) > 0 is true, so Go already knows the whole thing is true
	name := "Sam"
	if len(name) > 0 || isValidName(name) {
		fmt.Println("|| short-circuited: isValidName() was never called")
	}

	// practical use: safe nil check before using a value
	_, remainder, err := intDivision(10, 5)
	if err == nil && remainder == 0 {
		fmt.Println("clean division — safe because if err != nil, remainder is never checked")
	}

	// without short-circuit you'd need ugly nesting:
	_, remainder2, err2 := intDivision(10, 5)
	if err2 == nil {
		if remainder2 == 0 {
			fmt.Println("clean division — but nested, less readable")
		}
	}
}

func intDivision(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func isValidName(name string) bool {
	fmt.Println("isValidName was called!") // you'd see this if it wasn't short-circuited
	return len(name) > 0
}
