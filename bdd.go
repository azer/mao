package mao

import (
	"fmt"
)

type It func(title string, fn func(Expect))

func Desc (desc string, wrapper func(It) ) {

	wrapper(func (it string, fn func(Expect)) {
		test := Test{ fmt.Sprintf("%s %s", desc, it), fn }
		test.Run()
	})

}
