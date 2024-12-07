// This file demonstrates Go's concurrency features
// Go provides goroutines for concurrent execution and
// channels for communication between goroutines

package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple goroutine example
// This function will be run concurrently
// wg *sync.WaitGroup is used to wait for the goroutine to finish
func printNumbers(from, to int, wg *sync.WaitGroup) {
	// defer ensures wg.Done() is called when the function returns
	// This tells the WaitGroup that this goroutine is finished
	defer wg.Done()
	
	for i := from; i <= to; i++ {
		time.Sleep(100 * time.Millisecond) // Simulate work
		fmt.Printf("Number: %d\n", i)
	}
}

// Channel example
// This function demonstrates how to send data through a channel
// Channels are typed conduits through which you can send and receive values
func generateNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send number to channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Always close channels when done sending
}

func main() {
	// ==================== Goroutines Example ====================
	fmt.Println("Goroutines Example:")
	var wg sync.WaitGroup
	
	// Start two goroutines
	// WaitGroup is used to wait for all goroutines to complete
	wg.Add(2) // Tell WaitGroup to wait for 2 goroutines
	
	// go keyword starts a new goroutine
	// goroutines are lightweight threads managed by Go runtime
	go printNumbers(1, 3, &wg)
	go printNumbers(4, 6, &wg)
	
	// Wait for goroutines to finish
	// This blocks until all goroutines call wg.Done()
	wg.Wait()

	// ==================== Channels Example ====================
	fmt.Println("\nChannels Example:")
	// Create an unbuffered channel
	// Unbuffered channels block until both sender and receiver are ready
	numbers := make(chan int)

	// Start producer goroutine
	go generateNumbers(numbers)

	// Receive numbers from channel
	// range over channel continues until channel is closed
	for num := range numbers {
		fmt.Printf("Received: %d\n", num)
	}

	// ==================== Buffered Channel Example ====================
	fmt.Println("\nBuffered Channel Example:")
	// Create a buffered channel with capacity 3
	// Buffered channels block only when buffer is full (sending) or empty (receiving)
	bufferedCh := make(chan int, 3)
	
	// Send values without blocking
	// These won't block because buffer has space
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	
	// Receive values
	fmt.Printf("From buffered channel: %d\n", <-bufferedCh)
	fmt.Printf("From buffered channel: %d\n", <-bufferedCh)
	fmt.Printf("From buffered channel: %d\n", <-bufferedCh)

	// ==================== Select Example ====================
	fmt.Println("\nSelect Example:")
	// Create two channels for select demonstration
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines that send messages
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Select statement lets you wait on multiple channel operations
	// It's like a switch for channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
