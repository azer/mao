## Māo

Pragmatic BDD Testing Framework For Go. 

```go
import "math"

Desc("math.Abs", func(it It) {
  it("returns the absolute value of x", func(expect Expect) {
    expect(math.Abs(-12)).Equal(12.0)
    expect(math.Abs(-.5)).Equal(0.5)
  })
})
```

**Why Māo?**

* More flexibility: You can locate your tests in any directory.
* Don't repeat yourself: Māo does the ceremony for you by preprocessing the test module.
* Minimalistic Output: Māo doesn't have verbose outputs. It'll output short and briefly, will show you what lines fail in case of any error.

Update: **[Development Status](#development-status)**

## Install

Install the Go library first;

```bash
$ go install github.com/azer/mao
```

And optionally, install the command-line tool to avoid the boilerplate code: *See examples/simple-with-boilerplate.go if you'd like to skip this step*

```bash
$ curl https://raw.github.com/azer/mao/master/install | sh
```

## Usage

Create a test module anywhere you want and import the code you'd like to test:

```go
import "math/rand" // or: import "github.com/you/repository"

Desc("math.Abs", func(it It) {
  it("returns the absolute value of x", func(expect Expect) {
    expect(math.Abs(-12)).Equal(12.0)
    expect(math.Abs(-.5)).Equal(0.5)
  })
})
```

And run it with `mao` command:

```bash
$ mao test.go
```

It'll either output a simple success message that will look like;

```
Ran 3 tests successfully.
```

Or some error messages with failing source code lines, like following;

![](https://i.cloudup.com/CHNocClka1.png)

## Reference

### BDD API

* Desc
* It
* BeforeEach
* AfterEach

### Assertion API

* Above
* Equal
* NotEqual
* NotExist
* ResponseBody
* Lower

## Development Status

I recently think of two major changes in the project;

* Rewriting [the preprocesser](https://github.com/azer/mao/blob/master/mao) (bash script) in Go
* Replacing "desc/it" with "test". 

So it'd be looking like;

```go
import "math"

Test("returns the absolute value of x", func(e Expect) {
  expect(math.Abs(-12)).Equal(12.0)
  expect(math.Abs(-.5)).Equal(0.5)
})
```

At this point, I wouldn't recommend to use Mao. I still have the same goal but it may cause some pain since its preprocesser is not reliable enough. 

And since other projects occupy my time nowadays, I'd appreciate pull requests to reach the goals of the project. Thanks for your interest!
