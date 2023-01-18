package main

import "log"

var word = "now"

func main() {
	log.Println(word)
	// try to pass the word and change it
	changeWordFake(word)
	log.Println(word)

	// change it with pointer
	changeWord(&word)
	log.Println(word)

}

func changeWordFake(word string) {
	log.Println("masuk change fake, word : ", word)
	newFakeWord := "fake"
	word = newFakeWord
}

func changeWord(word *string) {
	log.Println("word in func changeWord is ", word)
	newWord := "berubah"
	*word = newWord
}
