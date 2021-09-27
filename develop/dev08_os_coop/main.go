package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	userCancel := make(chan os.Signal, 1)
	signal.Notify(userCancel, os.Interrupt)

	userInput := make(chan string)
	done := make(chan struct{})

	go readUserInput(userInput, done)

loop:
	for {
		select {
		case <-userCancel:
			close(done)
			break loop
		case command := <-userInput:
			words := strings.Split(command, " ")
			fmt.Println()
			switch words[0] {
			case "pwd":
				dir, _ := os.Getwd()
				fmt.Printf("%s\n", dir)
			}
		}
	}
}

func readUserInput(out chan<- string, done <-chan struct{}) {
	for {
		fmt.Printf("$ ")
		var command string
		fmt.Scanf("%s\n", &command)
		out <- command
	}
}
