# tenv

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
