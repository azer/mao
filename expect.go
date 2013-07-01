package mao

import "fmt"

type Expect func(val interface{}) *Expected

type Expected struct {
	Scope *Test
	Value interface{}
}

func (self *Expected) Equal (b interface{}) {
	if self.Value != b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to equal `%v`", self.Value, b))
	}
}

func (self *Expected) NotEqual (b interface{}) {
	if self.Value == b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to not equal `%v`", self.Value, b))
	}
}

func (self *Expected) NotExist () {
	if self.Value != nil {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to not exist.", self.Value))
	}
}
