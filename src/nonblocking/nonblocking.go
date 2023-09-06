package nonblocking

import (
	"fmt"
	"time"
)

func RunNonBlocking() {
	ch := make(chan int, 1) // Create a channel with a buffer capacity of 1

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42 // Put a value into the channel
	}()

	// Attempt to send a value to the channel
	select {
	case ch <- 10:
		fmt.Println("Sent the value 10 to the channel")
	default:
		fmt.Println("The channel is blocked. Skipping the send.")
	}

	// Wait for a value to be received from the channel
	value := <-ch
	fmt.Println("Received a value from the channel:", value)
}
