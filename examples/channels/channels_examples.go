package main

import (
	"fmt"
	"time"
)

func main() {
	channelStatesExample()
}

func channelStatesExample() {
	c1 := make(chan int)

	// you cannot send on closed channel
	// go func() {
	// 	c1 <- "value"
	// 	fmt.Println("added value to the channel")
	// }()
	// close(c1)

	c1 = make(chan int)

	// you can read from closed channel
	go func() {
		value := <-c1
		fmt.Println("Got value %d", value)

		value1, isClosed := <-c1
		if !isClosed {
			fmt.Println("Got value %d", value1)
		} else {
			fmt.Println("Channel closed")
		}

	}()
	close(c1)

	time.Sleep(2 * time.Second)
}

func timersExample() {
	timer1 := time.NewTimer(2 * time.Second)

	// lets wait 2 seconds
	a := <-timer1.C
	fmt.Println("Timer one fired at %d", a)
	timer1.Reset(3 * time.Second)

	a = <-timer1.C
	fmt.Println("Timer one fired second time at %d", a)

	// we can stop timer befor it fired
	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		fmt.Println("Trying to reach channel")
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()

	var timerStop = timer2.Stop()
	if timerStop {
		fmt.Println("Timer2 was stopped")
	} else {
		<-timer2.C
	}

	// if we want to wait just use
	time.Sleep(time.Second * 2)
}
func example5_routine(c1 chan string) {
	time.Sleep(10 * time.Second)
	for i := 0; i < 20; i++ {

		//c1 <- fmt.Sprintf("num%d", i)
		//fmt.Println("Added %d to queue", i)
	}
	close(c1)
}

func example5() {
	c1 := make(chan string, 5)
	go example5_routine(c1)

	for element := range c1 {
		fmt.Println(element)
	}
}

func example4() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "hey1"
	}()

	select {
	case res := <-c1:
		println(res)
	case <-time.After(1 * time.Second):
		println("timeout!")
	}

	time.Sleep(2 * time.Second)

	// we can use for to read multiple messages from the channel
	for {
		select {
		case res := <-c1:
			println(res)
			return
		default:
			println("waiting")
			time.Sleep(1 * time.Second)

		}
	}

}

func example3_select() {
	myChan := make(chan string)

	go example3_produce_a(myChan)
	go example3_produce_b(myChan)

	for i := 0; i < 20; i++ {
		select {
		case receivedValue := <-myChan:
			fmt.Println(receivedValue)
		}
	}
}

func example3_produce_a(pipe chan<- string) {
	for i := 0; i < 10; i++ {
		pipe <- fmt.Sprintf("%dA", i)
		time.Sleep(time.Second)
	}
}

func example3_produce_b(pipe chan<- string) {
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		pipe <- fmt.Sprintf("%dB", i)
		time.Sleep(time.Second)
	}
}

func example2() {
	done := make(chan bool, 1)
	go example2_worker(done)

	<-done
}

func example2_worker(done chan<- bool) { // one way write only channel
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	//myMsg := <-done
}

func example1() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		println("this is after ping") // if the buffer is full then client is blocked
		messages <- "ping1"
		println("this is after ping1")
		messages <- "ping2"
		println("this is after ping2")
		messages <- "ping3"
		println("this is after ping3")
		messages <- "ping4"
		println("this is after ping4")
	}()

	println("ready to read from the channel")
	msg := <-messages
	fmt.Println(msg)

	println("reading again...")
	msg = <-messages
	fmt.Println(msg)

	time.Sleep(2 * time.Second)
	msg = <-messages

	time.Sleep(2 * time.Second)
	msg = <-messages

	time.Sleep(2 * time.Second)
	msg = <-messages

	// msg = <-messages // if there are no gorouties that are attached to chanell it will fail.
	println("Exit")

}
