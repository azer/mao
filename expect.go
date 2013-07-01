package mao

import (
	"fmt"
	"reflect"
)

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

	msg := fmt.Sprintf("Expected `%v` to not exist.", self.Value)

	if self.Value == nil {
		return
	}

	if self.Value != nil {
		self.Scope.PrintError(msg)
		return
	}

	v := reflect.ValueOf(self.Value)

	fmt.Println("value:", self.Value, "value == nil", self.Value == nil, " v.isNil?", v.IsNil())

	if ! v.IsNil() {
		self.Scope.PrintError(msg)
	}
}
