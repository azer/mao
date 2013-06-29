package mao

type Test struct {
	Title string
	Fn func(Expect)
}

func (test *Test) Run () {
	test.Fn(Expect { test })
}
