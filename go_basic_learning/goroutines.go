package main

import (
	"fmt"
	"time"
)

func digits(number int, dchnl chan int) {
	for number != 0 {

		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		fmt.Println("calcSquares")
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		fmt.Println("calcCubes")
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func runWhileTrue() {
	for {
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("loggggggg:")
	}
}

func main() {
	go runWhileTrue()
	//go runWhileTrue()
	//go runWhileTrue()
	//go runWhileTrue()
	//go runWhileTrue()

	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("main:")
	}
	// number := 589
	// sqrch := make(chan int)
	// cubech := make(chan int)
	// go calcSquares(number, sqrch)
	// go calcCubes(number, cubech)
	// squares, cubes := <-sqrch, <-cubech
	// fmt.Println("Final output", squares+cubes)
	// fmt.Println("a")
}
