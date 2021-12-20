package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		text string = "Lorem Ipsum is simply dummy text of the printing and typesetting industry i"
	)

	fmt.Println(Shortest(text))

}

func Shortest(text string) string {

	temp := strings.Split(text, " ")

	var word string = temp[0]

	for _, i := range temp {
		if len(word) > len(i) {
			word = i
		}
	}

	return word

}
