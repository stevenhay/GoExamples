package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// If statements
	a := 2
	if a > 1 {
		fmt.Println("Standard if statement")
	}

	rand.Seed(time.Now().UnixNano())
	if n := rand.Int(); n%2 == 0 {
		fmt.Printf("Assignment and condition: %d is even\n", n)
	} else {
		fmt.Printf("Assignment and condition: %d is odd\n", n)

		n := 10
		fmt.Printf("Shadowing variables is allowed: %d\n", n)
	}

	// For loops
	for i := 0; i < 10; i++ {
		fmt.Printf("Standard for loop: %d\n", i)
	}

	i := 0
	for i < 10 {
		fmt.Printf("For loopa is also the while: %d\n", i)
	}

	for {
		fmt.Println("Infinite loop")
		break
	}

	for i := 0; ; i++ {
		fmt.Println("Infinite loop with counting variable")
		break
	}

	// Switch statements
	switch n := rand.Intn(10); n {
	case 0, 1, 2, 3:
		fmt.Println("Can have multiple variables in a single case")
		fallthrough // fallthrough is explicit unlike Java
	case 4:
		fmt.Println("0 or 1")
	default: // default not being the last case is a warning in go, but not an error
		fmt.Println("5 or 9")
	case 6, 7, 8:
		fmt.Println("2 or 3")
	}

}
