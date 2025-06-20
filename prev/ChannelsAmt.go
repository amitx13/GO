/*

Assignment
Run the code as-is. You should see that it doesn't do anything interesting: no ping or pong messages are printed.

Fix the bug in the pingPong function.

Remember: if a program exits before its goroutines have completed, those goroutines will be killed silently.
Which of the function calls probably shouldn't run in the background as a goroutine?

*/

//Using WaitGroups

/* package main

import (
	"fmt"
	"sync"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go ponger(pings, pongs, wg)
	go pinger(pings, numPings, wg)

	i := 0
	for range pongs {
		fmt.Println("got pong", i)
		i++
	}
	fmt.Println("pongs done")

	wg.Wait()
}

func pinger(pings chan struct{}, numPings int, wg *sync.WaitGroup) {
	defer wg.Done()
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}
*/

//Without WaitGroups

/* package main

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	done := make(chan struct{})

	go ponger(pings, pongs)
	go pinger(pings, numPings)

	go func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
		close(done) // Signal that we're done
	}()

	<-done // Block main goroutine until done
}

// don't touch below this line

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
} */

/*
Implement the processMessages function, it takes a slice of messages. It should process each message concurrently within a goroutine.
Use the process function for this to simulate the benefit of using goroutines.
Use a channel to ensure that all messages are processed and collected correctly, then return the slice of processed messages.
*/
package main

import (
	"fmt"
	"time"
)

func processMessages(messages []string) []string {
	result := make(chan string, len(messages))
	for _, msg := range messages {
		go func(msg string, res chan<- string) {
			res <- process(msg)
		}(msg, result)
	}
	processMessages := make([]string, len(messages))
	for i := range messages {
		processMessages[i] = <-result
	}
	return processMessages
}

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

func hellomain() {
	messages := []string{"messages", "is", "and"}
	processedMessages := processMessages(messages)
	fmt.Println(processedMessages)
}
