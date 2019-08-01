# tenv [![Build Status][travis-ci-badge]][travis-ci] [![GoDoc][godoc-badge]][godoc]

> Set/unset temporary environment variables for testing

## Installation

```sh
go get github.com/charliekenney23/tenv
```

## Usage

```go
{
  c := tenv.New()
  defer c.Restore()

  c.Set("foo", "bar")
  c.Unset("baz")
} // environment restored to initial state
```

&copy; 2019 [Charles Kenney](https://github.com/charliekenney23)

[travis-ci-badge]: https://travis-ci.org/Charliekenney23/tenv.svg?branch=master
[travis-ci]: https://travis-ci.org/Charliekenney23/tenv
[godoc-badge]: https://godoc.org/github.com/Charliekenney23/tenv?status.svg
[godoc]: https://godoc.org/github.com/Charliekenney23/tenv
