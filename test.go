package test

import "fmt"

type Expect func(a, b interface{})

func It (title string, fn func(Expect) ) {
	fn(newExpect(title))
}

func newExpect(title string) Expect {
	var errorCount int = 0

	return func(a, b interface{}) {

		if a != b {
			errorCount++

			if errorCount == 1 {
				PrintTitle(title)
			}

			PrintError(fmt.Sprintf("Expected `%s` to equal `%s`\n", a, b))
		}

	}

}
