package mao

import "fmt"

var lastTitle string;

func (test *Test) PrintTitle (title string) {
	fmt.Printf("\033[37m \033[1m\n    %s \n\n", title)
}

func (test *Test) PrintError (message string) {

	if lastTitle != test.Title {
		lastTitle = test.Title
		test.PrintTitle(test.Title)
	}

	fmt.Printf("\033[31m \033[1m        %s\033[0m", message)
}
