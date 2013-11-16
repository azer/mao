Desc("Expect", func (it It) {
	it("has an above method to assert target is greater than value", func (expect Expect) {
		expect(5).Above(3)
	})

	it("has an equal method to assert target is greater than value", func (expect Expect) {
		expect(5).Equal(5)
	})
})
