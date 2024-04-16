package main

import (
	"fmt"
)

func main() {
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	// // Get a greeting message and print it.
	// message, err := greetings.Hello("")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)

	var x [5]int = [5]int{1, 2, 3, 4, 5}

	var s []int = x[1:2]
	fmt.Println(x)
	fmt.Println(s)

	var m map[string]int = map[string]int{
		"apple":  1,
		"mango":  2,
		"orange": 5,
	}

	fmt.Println(m)

	v, ok := m["papapya"]

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}
}
