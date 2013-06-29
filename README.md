## Māo

BDD Testing Framework For Go

## Usage

```go
package main

import (
  . "github.com/azer/test.go"
  "math"
)

func main () {

  It("returns the absolute value of x", func(expect Expect) {
    expect(math.Abs(-5), 5.0)
    expect(math.Abs(-15), 123)
  })

  It("returns the greatest integer value less than or equal to x", func(expect Expect) {
    expect(math.Floor(10.3), 10.0)
    expect(math.Floor(10.9), 10.0)
  })

}
```

## Projects Using It

* [ansi-codes](http://github.com/azer/go-ansi-codes)

![](https://dl.dropboxusercontent.com/s/77k6n4vxjhgbauf/npmel_36.jpg)
