package functional

import (
	"fmt"
	"time"

	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/result"
)

// Allocate a channel.
var channel1 = make(chan int)
var channel2 = make(chan int)
var ints = []int{0, 1, 2, 3}
var items = iter.Collect[int](
	iter.Map[int](
		iter.Lift(ints),
		double,
	),
)

var results = result.Ok[int](100)

func double(a int) int { return a * 2 }

func routine(
	collection []int,
	channel chan int,
	duration time.Duration,
	signal int,
) {
	time.Sleep(duration * time.Second)
	fmt.Println(collection)
	channel <- signal
}

func RunFun() {
	fmt.Println(ints)

	go routine(items, channel1, 1, 1)
	go routine(items, channel2, 2, 2)

	var signal1 = <-channel1
	var signal2 = <-channel2

	fmt.Println(results)
	fmt.Println(signal1)
	fmt.Println(signal2)
}
