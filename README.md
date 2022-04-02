[![Build Status](https://github.com/hoani/getset/workflows/Go/badge.svg)](https://github.com/hoani/getset/actions?workflow=Go)
[![Coverage Status](https://coveralls.io/repos/github/hoani/getset/badge.svg?branch=master)](https://coveralls.io/github/hoani/getset?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoani/getset)](https://goreportcard.com/report/github.com/hoani/getset)

```sh
go get github.com/hoani/getset@latest
```

## getset

Is a lightweight library implementing sets in go. 

This package uses generics, so go 1.18 or higher is required.

## Usage

```go
strSet := getset.New("a", "b", "c")

strSet.Insert("d") // Set is now "a", "b", "c", "d"

delete(strSet, "b") // Set is now "a", "c", "d"

len(strSet) // 3

strSet.Has("c") // true
strSet.Has("e") // false

strSet.ToArray() // []string{"a", "c", "d"}
```

And that's it!

## License

MIT