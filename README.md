# Grand

[![GoDoc](https://godoc.org/github.com/ernsheong/grand?status.svg)](http://godoc.org/github.com/ernsheong/grand)

Grand is a Go random string generator.

## Installation

`go get github.com/ernsheong/grand`

## Usage

1. **IMPORTANT.** Seed `rand` first to ensure you don't get the same string on every code run (initialize):

    ```go
    rand.Seed(time.Now().UTC().UnixNano())
    ```

    or

    ```go
    rand.Seed(time.Now().Unix())
    ```

1. Generate your random string, given a length parameter `n`.

    ```go
    grand.GenerateRandomString(32)
    // returns "qzrWbaoLTVpQoottZyPFfNOoMioXHRuF"
    ```

1. Generate random string from other character sets

    ```go
    gen := grand.NewGenerator(grand.CharSetBase62)
	gen.GenerateRandomString(20)
    // returns "q3rWba2LTVpQ4ottZyPv"
    ```

## Concurrency

From the `math/rand` [docs](https://golang.org/pkg/math/rand/):
> The default Source is safe for concurrent use by multiple goroutines

`grand` uses the default Source, and hence is safe to be called from multiple goroutines, at a slight performance penalty. See https://stackoverflow.com/a/31832326/1161743 for details.

## Credits

I claim no credit for the code here. It's originally from user icza in https://stackoverflow.com/a/31832326/1161743.
