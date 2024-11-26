package main

import "fmt"

// conversion: one tyoe to another type. (in another programming language , it is called casting)
// scope: you can access variable in the declation to lower side in the declaration area
/*
	In Go, housekeeping generally refers to the set of tasks or processes aimed at maintaining
	the program's runtime efficiency, organization, and memory management. It isn't a formal
	concept specific to Go but is commonly used in software engineering to describe the background
	or maintenance activities needed to keep a program running smoothly.
*/
func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 42.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can •t take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type XT \n", z, z)
	*/

	/*	// conversion
		// this does work
		z = float64(m)
		fmt.Printf("%v of type %T \n", z, z)
	*/
}
