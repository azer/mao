package main

import . "github.com/azer/mao"

func main () {

	Desc("math.Abs", func(it It) {

		it("returns the absolute value of x", func(expect Expect) {
			expect.NotEqual(1, 2)
			expect.Equal(0.5, 0.5)
		})

		it("should return x if is positive", func(expect Expect) {
			expect.Equal(0.7, 0.5)
		})

	})

	Desc("math.Floor", func(it It) {

		it("returns the greatest integer value less than or equal to x", func(expect Expect) {
			expect.Equal(13.4, 13.0)
			expect.Equal(13.4, 13.5)
		})

		it("should return x if is int", func(expect Expect) {
			expect.Equal(13, 13.0)
		})

	})

}
