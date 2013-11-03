package main

import "math"
import . "github.com/azer/mao"

func main() {

	Desc("math.Abs", func(it It) {

		it("returns the absolute value of x", func(expect Expect) {
			expect(math.Abs(-12)).Equal(12.0)
			expect(math.Abs(-.5)).Equal(0.5)
		})

	})

	Desc("math.Floor", func(it It) {

		it("returns the greatest integer value less than or equal to x", func(expect Expect) {
			expect(math.Floor(13.4)).Equal(13.0)
			expect(math.Floor(13.99)).Equal(13.0)
		})

		it("should return x if is int", func(expect Expect) {
			expect(math.Floor(13)).Equal(13.0)
		})

	})

}
