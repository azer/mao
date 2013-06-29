package mao

import "fmt"

type Expect struct {
	Scope *Test
}

func (self *Expect) Equal(a, b interface{}) {
	if a != b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%s` to equal `%s`\n", a, b))
	}
}

func (self *Expect) NotEqual(a, b interface{}) {
	if a == b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%s` to not equal `%s`\n", a, b))
	}
}
