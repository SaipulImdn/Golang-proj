package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data

	}

	go sayHelloTo("syaiful imanudin")
	go sayHelloTo("rizky")
	go sayHelloTo("ahmad")
	go sayHelloTo("abdul")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	var message4 = <-messages
	fmt.Println(message4)

	{
		runtime.GOMAXPROCS(2)

		var messages = make(chan string)

		for _, each := range []string{"syaiful", "abdul", "ahmad"} {
			go func(who string) {
				var data = fmt.Sprintf("hello %s", who)
				messages <- data
			}(each)
		}

		for i := 0; i < 3; i++ {
			printMessage((messages))
		}
	}
}
