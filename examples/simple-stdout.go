import (
	"math"
	"fmt"
)

Desc("math.Abs", func(it It) {
	it("returns the absolute value of x", func(expect Expect) {
		fmt.Println("before")
		expect(math.Abs(-12)).Equal(12.0)
		expect(math.Abs(-.5)).Equal(0.5)
		fmt.Println("after")
	})
})
