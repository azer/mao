package main

import "math"
import . "github.com/azer/mao"

func main () {

	Desc("math.Abs", func(it It) {

		it("returns the absolute value of x", func(expect Expect) {
			expect.Equal(math.Abs(-12), 12.0)
			expect.Equal(math.Abs(-.5), 0.5)
		})

	})

	Desc("math.Floor", func(it It) {

		it("returns the greatest integer value less than or equal to x", func(expect Expect) {
			expect.Equal(math.Floor(13.4), 13.0)
			expect.Equal(math.Floor(13.99), 13.0)
		})

		it("should return x if is int", func(expect Expect) {
			expect.Equal(math.Floor(13), 13.0)
		})

	})

}
