package main

import "fmt"

func main() {


	//Style #1
	var counter int
	counter = 0
	for counter < 10 {          //  while loop
		fmt.Printf("Hello World!\n")
		counter ++
	}

	//Style #2
	for counter2 := 0 ; counter2 < 10 ; counter2++ {
		fmt.Printf("Hello World! version 2 \n")
	}



	//Style #3  multiple initiation variables
	for i, j := 0, 1; i < 10  ; i, j = i+1, j*2 {
		fmt.Printf("i = %d\t j = %d\n", i, j)
	}

	// Style $4  while loop boolean
	counter = 0
	var stop bool
	for !stop {
		fmt.Printf("for/while loop count : %d\n", counter)
		counter++
		if counter == 10 {
			stop = true
		}
	}

}
