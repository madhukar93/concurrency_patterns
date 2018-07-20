/*
In computing, the producer–consumer problem is a classic example of
a multi-process synchronization problem. The problem describes two processes,
the producer and the consumer, who share a common, fixed-size buffer
used as a queue. The producer's job is to generate data,
put it into the buffer, and start again. At the same time,
the consumer is consuming the data, one piece at a time.

The problem is to make sure that the producer won't try to add data into the buffer
if it's full and that the consumer won't try to remove data from an empty buffer

(https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem)
*/
package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int, val int) {
	for {
		ch <- val
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}

func consume(ch <-chan int) {
	for {
		fmt.Println(<-ch)
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 100)
	// done := make(chan int)
	fmt.Println("producer consumer")
	go produce(ch, 1)
	go produce(ch, 2)
	go produce(ch, 3)
	go consume(ch)
	time.Sleep(time.Duration(10) * time.Second)
	// <-done
}
