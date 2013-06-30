package mao

import (
	"fmt"
	"runtime"
	"io/ioutil"
	"path"
	"strings"
)

type FailingLine struct {
	content string
	filename string
	next string
	number int
	prev string
}

var lastTitle string;

var (
	reset string = "\033[0m"
	white string = "\033[37m\033[1m"
	grey string = "\x1B[90m"
	red string = "\033[31m\033[1m"
)

func (test *Test) PrintTitle (title string) {
	fmt.Printf("\033[37m \033[1m\n    %s \n\n", title)
}

func (test *Test) PrintError (message string) {
	if lastTitle != test.Title {
		lastTitle = test.Title
		test.PrintTitle(test.Title)
	}

	failingLine, err := getFailingLine()

	if err != nil {
		return;
	}

	fmt.Printf("%s        %s %s %s %s\n", red, message, grey, path.Base(failingLine.filename), reset)
	test.PrintFailingLine(&failingLine)
}

func (test *Test) PrintFailingLine (failingLine *FailingLine) {
	fmt.Printf("%s             %d. %s\n", grey, failingLine.number - 1, failingLine.prev)
	fmt.Printf("%s             %d. %s %s\n", white, failingLine.number, failingLine.content, reset)
	fmt.Printf("%s             %d. %s\n", grey, failingLine.number + 1, failingLine.next)
	fmt.Println(reset)
}

func getFailingLine () (FailingLine, error) {
	_, filename, ln, _ := runtime.Caller(3)

	bf, err := ioutil.ReadFile(filename)

	if err != nil {
		return FailingLine{}, fmt.Errorf("Failed to open %s", filename)
	}

	lines := strings.Split(string(bf), "\n")[ln-2:ln+2]

	return FailingLine{
		softTabs(lines[1]),
		filename,
		softTabs(lines[2]),
		int(ln),
		softTabs(lines[0]),
	}, nil

}

func softTabs (text string) string {
	return strings.Replace(text, "\t", "  ", -1)
}
