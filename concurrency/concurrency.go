package main

import "fmt"

func main() {
	c := make(chan int) //Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go coockingGopher(i, c) // start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func coockingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id
}
