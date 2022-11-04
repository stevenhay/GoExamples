package main

import "fmt"

func deferCalls() {
	defer hello()
	defer world()

	fmt.Println("12345")
}

func world() {
	fmt.Println("World!")
}

func hello() {
	fmt.Println("Hello!")
}

func sum(x, y int) (z int) {
	defer func() {
		z++
	}()

	z = x + y
	return
}

func appendDefer() []int {
	i := []int{1, 4, 5}

	// this is an error as defer must disregard
	// the return types of the method
	// defer append(i[:1], 2, 3)
	//
	// if this is needed, instead you can use
	// anonymous method instead
	defer func() {
		_ = append(i[:1], 2, 3)
	}()

	return i
}

func badDefer() {
	for i := 0; i < 100; i++ {
		// assume we're doing some stuff in here and
		// we're defering something like closing a
		// io stream
		defer func() {}()

		// this defer call stack would only get executed after
		// the method ends, and the defer stack would be 99 in size!
	}

	// instead
	for i := 0; i < 100; i++ {
		// now each defer will get evaluated after each loop
		// as the function containing the defer ends
		// each loop
		func() {
			defer func() {}()
		}()
	}
}
