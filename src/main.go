package main

import (
	"fmt"
	"gofun/src/nonblocking"
)

func main() {
	fmt.Println("This is the main function")

	nonblocking.RunNonBlocking()

	// functional.RunFun()
}
