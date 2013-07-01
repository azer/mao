## Māo

BDD Testing Framework For Go.

```go
Desc("math.Abs", func(it It) {
  it("returns the absolute value of x", func(expect Expect) {
    expect(math.Abs(-12)).Equal(12.0)
    expect(math.Abs(-.5)).Equal(0.5)
  })
})
```

See `examples/` for more info.

## How does it output?

![](https://pbs.twimg.com/media/BN_4mnrCUAAR24G.png:large)

<!--
![](https://dl.dropboxusercontent.com/s/77k6n4vxjhgbauf/npmel_36.jpg)
-->
