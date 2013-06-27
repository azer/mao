package test

import "fmt"

func PrintTitle (title string) {
	fmt.Printf("\033[37m \033[1m\n    %s \n\n", title)
}

func PrintError (message string) {
	fmt.Printf("\033[31m \033[1m        %s\033[0m", message)
}
