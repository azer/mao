import "math/rand"

Desc("randInt", func(it It) {
	var (
		current int
		last int
	)

	BeforeEach(func () {
		current = rand.Int()
	})

	AfterEach(func () {
		last = current
	})

	it("generates a random number", func (expect Expect) {
		expect(current).Above(0)
	})

	it("generates a unique number on each call", func (expect Expect) {
		expect(current).Above(0)
		expect(last).Above(0)
		expect(current).NotEqual(last)
	})

})
