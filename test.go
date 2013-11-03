package mao

type Test struct {
	Title string
	Fn func(Expect)
	BeforeEach BeforeAfterFn
	AfterEach BeforeAfterFn
}

func (test *Test) Run () {
	incTestCounter()

	if test.BeforeEach != nil {
		test.BeforeEach()
	}

	test.Fn(func (val interface{}) *Expected {
		return &Expected{ test, val }
	})

	if test.AfterEach != nil {
		test.AfterEach()
	}
}
