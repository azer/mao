package mao

import "fmt"

type Expect func(val interface{}) *Expected

type Expected struct {
	Scope *Test
	Value interface{}
}

func (self *Expected) Equal(b interface{}) {
	if self.Value != b {
		self.Scope.PrintError(fmt.Sprintf("Expecteded `%s` to equal `%s`\n", self.Value, b))
	}
}

func (self *Expected) NotEqual(b interface{}) {
	if self.Value == b {
		self.Scope.PrintError(fmt.Sprintf("Expecteded `%s` to not equal `%s`\n", self.Value, b))
	}
}
