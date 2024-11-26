package main
import (
	"fmt"
)
func main() {
	// const name, age = "Kim", 22
	const name = "Kim"
	const age = 22
	// fmt.Printf("%s is %d years old.\n", name, age)
	// fmt.Printf("%s is %d years old.\t The type is %T and %T", name, age, name, age)
	fmt.Printf("%v is %v years old.\t The type is %T and %T", name, age, name, age)
}
