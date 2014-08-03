package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func analyze(cResponse chan string) {
	for {
		sOneResponse := <-cResponse
		// fmt.Println(sOneResponse)
		fmt.Println("length of the page : " + strconv.Itoa(utf8.RuneCountInString(sOneResponse)))
	}

}
