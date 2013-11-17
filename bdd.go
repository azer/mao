package mao

import (
	"fmt"
)

type BeforeAfterFn func()
type It func(title string, fn func(Expect))

var (
	Describe func (desc string, wrapper func(It)) = Desc
	beforeEachFn BeforeAfterFn
	afterEachFn  BeforeAfterFn
)

func Desc(desc string, wrapper func(It)) {
	beforeEachFn = nil
	afterEachFn = nil

	wrapper(func(it string, fn func(Expect)) {
		test := Test{fmt.Sprintf("%s %s", desc, it), fn, beforeEachFn, afterEachFn}
		test.Run()
	})
}

func AfterEach(fn BeforeAfterFn) {
	afterEachFn = fn
}

func BeforeEach(fn BeforeAfterFn) {
	beforeEachFn = fn
}
