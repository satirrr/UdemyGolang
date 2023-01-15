package main

import (
	"fmt"
	"github.com/satirrr/UdemyGolang/Channel/helpers"
)

const numbPool = 10

func calculateValue(intchan chan int) {
	randomNumber := helpers.RandomNumber(numbPool)
	intchan <- randomNumber
}

func main() {
	// syntax to create a channel
	//this is a channel which only hold an integer
	intChan := make(chan int)
	defer close(intChan)
	// ITS A GOOD PRACTICE TO CLOSE A CONNECTION TO ANYTHING
	// for example when you are opening a file or when you are opening a connection to a database

	// go routin run a the same time, for example if we run 5 or 10 go routin it will run at a same time
	go calculateValue(intChan)

	//num is whatever value intChan is
	num := <-intChan
	fmt.Println(num)
}
