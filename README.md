# Go Nanoid

[![CI](https://github.com/enjoyZhou/go-nanoid/workflows/CI/badge.svg)](https://github.com/enjoyZhou/go-nanoid/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/enjoyZhou/go-nanoid/v6.svg)](https://pkg.go.dev/github.com/enjoyZhou/go-nanoid/v6)
[![Go Report Card](https://goreportcard.com/badge/github.com/enjoyZhou/go-nanoid/v6)](https://goreportcard.com/report/github.com/enjoyZhou/go-nanoid/v6)
[![GitHub issues](https://img.shields.io/github/issues/enjoyZhou/go-nanoid.svg)](https://github.com/enjoyZhou/go-nanoid/issues)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/enjoyZhou/go-nanoid/blob/main/LICENSE)

This package is a Go implementation of [ai's](https://github.com/ai)
[Nano ID](https://github.com/ai/nanoid). Its release number follows the
upstream Nano ID compatibility baseline; this release targets 6.0.1.

**Safe.** It uses cryptographically strong random generator.

**Compact.** It uses more symbols than UUID (`A-Za-z0-9_-`) and has a UUID v4-like collision probability in 21 symbols instead of 36.

**Fast.** Nanoid is as fast as UUID but can be used in URLs.

The Nano ID 6 alignment also removed repeated alphabet conversions from the
default generator. On an Apple M1 Pro with Go 1.24.6, the recorded `New()`
microbenchmark improved from 3264.7 ns/op and 24 allocations to 402.6 ns/op
and 3 allocations. Results are environment-specific; see the reproducible
[upgrade performance report](./PERFORMANCE.md) for raw runs, commands, and
scope limitations.

> [!NOTE]  
> There's little to no development on this repo, intentionally. It does what it needs to do. Bug reports are welcomed, features _might_ be implemented.
>
> If you are considering more heavy weight solution that integrates with UUIDs (supported by many databases) I would suggest you take a look at [typeid](https://github.com/sumup/typeid) or other equivalents.

## Install

Via go get tool

``` bash
$ go get github.com/enjoyZhou/go-nanoid/v6@v6.0.1
```

The repository publishes two tags for each synchronized release:

- `6.0.1`, matching the upstream Nano ID tag exactly;
- `v6.0.1`, required by Go modules for the `/v6` semantic import path.

## Usage

Generate ID

``` go
id, err := gonanoid.New()
```

Generate ID with a custom alphabet and length

``` go
id, err := gonanoid.Generate("abcde", 54)
```

The default size and URL-safe alphabet are aligned with Nano ID 6:

```go
fmt.Println(gonanoid.DefaultSize)  // 21
fmt.Println(gonanoid.URLAlphabet) // A-Za-z0-9_- in Nano ID's compression-optimized order
```

Custom alphabets may contain 1 to 256 Unicode characters. A size of `0`
returns an empty ID, matching Nano ID 6. Negative sizes and invalid alphabets
return an error.

## Notice

If you use Go Nanoid in your project, please let me know!

If you have any issues, just feel free and open it in this repository, thanks!

## Credits

- [ai](https://github.com/ai) - [nanoid](https://github.com/ai/nanoid)
- icza - his tutorial on [random strings in Go](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang)

## License

The MIT License (MIT). Please see [License File](./LICENSE) for more information.
