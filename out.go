package mao

import (
	"fmt"
	"runtime"
	"io/ioutil"
	"path"
	"strings"
	"os"
	"strconv"
)

type FailingLine struct {
	content string
	filename string
	next string
	number int
	prev string
}

var lastTitle string

var (
	testCounter = 0
	failCounter = 0
	startingLine = os.Getenv("MAO_LINENO_START")
	reset string = "\033[0m"
	white string = "\033[37m\033[1m"
	grey string = "\x1B[90m"
	red string = "\033[31m\033[1m"
)

func (test *Test) PrintTitle (title string) {
	fmt.Printf("\033[37m \033[1m\n    %s \n\n", title)
}

func (test *Test) PrintError (message string) {
	failCounter += 1

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

func PrintTestSummary () {
	if failCounter > 0 {
		fmt.Printf("\n  Ran %d tests, %d assertions failed.\n\n", testCounter, failCounter)
		return
	}

	fmt.Printf("\n  Ran %d tests successfully.\n\n", testCounter)
}

func getFailingLine () (FailingLine, error) {
	_, filename, ln, _ := runtime.Caller(3)

	bf, err := ioutil.ReadFile(filename)

	if err != nil {
		return FailingLine{}, fmt.Errorf("Failed to open %s", filename)
	}

	lines := strings.Split(string(bf), "\n")[ln-2:ln+2]
	filename = strings.Replace(filename, "mao_", "", 1)
	lineno := int(ln)

	if len(startingLine) > 0 {
		lndiff, _ := strconv.Atoi(startingLine)
		lineno -= lndiff
	}

	return FailingLine{
		softTabs(lines[1]),
		filename,
		softTabs(lines[2]),
		lineno,
		softTabs(lines[0]),
	}, nil

}

func incTestCounter () {
	testCounter += 1
}

func softTabs (text string) string {
	return strings.Replace(text, "\t", "  ", -1)
}
