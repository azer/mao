package mao

import (
	"fmt"
	"reflect"
	"net/http"
	"io/ioutil"
)

type Expect func(val interface{}) *Expected

type Expected struct {
	Scope *Test
	Value interface{}
}

func (self *Expected) Above(b interface{}) {
	value, ok := self.Value.(int)
	otherValue, otherOk := b.(int)

	if !ok || !otherOk {
		self.Scope.PrintError(fmt.Sprintf("Unable to compare `%v` with `%v`", self.Value, b))
		return
	}

	if value <= otherValue {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to be above than `%v`", self.Value, b))
	}
}

func (self *Expected) Equal(b interface{}) {
	if self.Value != b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to equal `%v`", self.Value, b))
	}
}

func (self *Expected) Lower(b interface{}) {
	value, ok := self.Value.(int)
	otherValue, otherOk := b.(int)

	if !ok || !otherOk {
		self.Scope.PrintError(fmt.Sprintf("Unable to compare `%v` with `%v`", self.Value, b))
		return
	}

	if value >= otherValue {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to be lower than `%v`", self.Value, b))
	}
}

func (self *Expected) NotEqual(b interface{}) {
	if self.Value == b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to not equal `%v`", self.Value, b))
	}
}

func (self *Expected) NotExist() {

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

	if !v.IsNil() {
		self.Scope.PrintError(msg)
	}
}

func (self *Expected) ResponseBody(b interface{}) {
	response, err := http.Get(self.Value.(string))

	if err != nil {
		self.Scope.PrintError(fmt.Sprintf("Unable to get %s", self.Value))
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		self.Scope.PrintError(fmt.Sprintf("Unable to read `%v`", self.Value))
		return
	}

	if string(body) != b {
		self.Scope.PrintError(fmt.Sprintf("Expected `%v` to equal `%v`", string(body), b))
	}
}
