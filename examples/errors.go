package main

import . "github.com/azer/mao"

func main () {

	Desc("math.Abs", func(it It) {

		it("returns the absolute value of x", func(expect Expect) {
			expect(1).Equal(2)
			expect(0.5).Equal(0.5)
		})

		it("should return x if is positive", func(expect Expect) {
			expect(0.7).Equal(0.5)
		})

	})

	Desc("math.Floor", func(it It) {

		it("returns the greatest integer value less than or equal to x", func(expect Expect) {
			expect(1).Equal(2)
			expect(0.5).Equal(0.5)
		})

		it("should return x if is int", func(expect Expect) {
			expect(13).Equal(13.0)
		})

	})

}
