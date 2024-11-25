package main

import "fmt"

func main() {
	// fmt.Printf: requires a format string and can format the output in various ways.
	// fmt.Println: automatically handles formatting for default types without needing a format string.

	// different types of declaration: 1. var 2. short declaration 3. variable with zero values 4. var when specificity  required + multiple declaration
	fmt.Printf("\t\tVARIABLE WITH VAR\n")
	// 1. var
	// DECLARE a variable to hold a VALUE of a certain TYPE
	// then ASSIGN a VALUE of that TYPE to the variable
	// initializing a variable
	var g int = 84
	fmt.Printf("g = %d\n", g)

	// 2. short declaration
	// := means short declaration by which Go will automatic understand the data type
	a := 5
	fmt.Println(a)

	// 3. variable with zero values
	fmt.Printf("\n\t\tVARIABLE WITH ZERO VALUES\n")
	// zero values: by default the value of int/float/string/boolean will be set to 0/0/""/false
	var h int
	fmt.Println(h)
	h = 184
	fmt.Println(h)

	var i float64
	fmt.Println(i)
	i = 184.19
	fmt.Println(i)

	var j string
	fmt.Println(j)
	j = "Gophers"
	fmt.Println(j)

	var k bool
	fmt.Println(k)
	k = true
	fmt.Println(k)

	// 3. var when specificity required + multiple declaration
	fmt.Printf("\n\t\tVARIABLE WITH SPECIFICITY REQUIRED\n")
	aa, bb, cc := 1, 2, "chill"
	fmt.Println(aa, bb, cc)

	fmt.Printf("\n\t\tBLANK IDENTIFIER\n")
	// multiple value ASSIGN and BLANK IDENTIFIER
	// BLANK IDENTIFIER (_): a variable that is declared but not assigned a value. in another words, a variable that is declared to skip
	a, b, c, _, d := 0, 1, 2, 3, "chill"
	fmt.Println(a, b, c, d)

	// this will not work, if any variable declared and not used
	// you cannot have unused variables in your code, the compiler doesn't allow it - this is called "code pollution"
	/*
		aa, bb, cc, ee := 0, 1, 2, 3
		fmt.Println(aa, bb, cc)
	*/
}
