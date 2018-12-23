package main

import "fmt"

func main() {

	//Style #0  infinte loop
	//for {
	//	fmt.Printf("Infinitle Loop \n")
	//}

	//Style #1
	var counter int
	counter = 0
	for counter < 10 { //  while loop
		fmt.Printf("#1 Hello World! counter: %d\n", counter)
		counter++
	}

	//Style #2
	for counter2 := 0; counter2 < 10; counter2++ {
		fmt.Printf("#2 Hello World! counter: %d\n \n", counter2)
	}

	//Style #3  multiple initiation variables
	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		fmt.Printf("#3  two vaiables :  i = %d\t j = %d\n", i, j)
	}

	// Style $4  while loop boolean
	counter = 0
	var stop bool
	for !stop {
		fmt.Printf("#4 for/while loop count : %d\n", counter)
		counter++
		if counter == 10 {
			stop = true
		}
	}

}
