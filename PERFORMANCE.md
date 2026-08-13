# Nano ID 6 Upgrade Performance Report

Date: 2026-08-13

This report records the functional and performance impact of aligning the Go
implementation with Nano ID 6.0.1.

The public Go module for this compatibility baseline is
`github.com/enjoyZhou/go-nanoid/v6@v6.0.1`. The repository also publishes the
matching upstream-style tag `6.0.1`.

## Compared revisions

| Version | Revision | Description |
|---|---|---|
| Before | [`2ab893b`](https://github.com/enjoyZhou/go-nanoid/commit/2ab893b) | Implementation before the Nano ID 6 alignment |
| After | [`5a1612d`](https://github.com/enjoyZhou/go-nanoid/commit/5a1612dc821b18c3055180967930a2928c64ec12) | Nano ID 6 alignment merged into `main` by [PR #1](https://github.com/enjoyZhou/go-nanoid/pull/1) |

The upgrade:

- aligns the default size and URL alphabet order with Nano ID 6.0.1;
- accepts custom alphabets containing up to 256 Unicode characters;
- handles a requested size of `0` consistently with Nano ID 6;
- updates the rejection-sampling prefetch calculation;
- removes third-party test dependencies;
- avoids repeatedly converting the default alphabet to `[]rune` while
  generating every ID character.

## Benchmark scope

The benchmark measures the default, secure, 21-character `New()` path:

```go
func BenchmarkNanoid(b *testing.B) {
    for n := 0; n < b.N; n++ {
        _, _ = New()
    }
}
```

It was run serially for the before and after revisions so the two test batches
did not compete for CPU resources.

Environment:

```text
Go:      go1.24.6
OS/Arch: darwin/arm64
CPU:     Apple M1 Pro
```

Command:

```bash
go test -run '^$' -bench '^BenchmarkNanoid$' \
  -benchmem -benchtime=1s -count=10
```

## Ten-run comparison

| Metric | Before | After | Change |
|---|---:|---:|---:|
| Mean time | 3264.7 ns/op | 402.6 ns/op | 8.1x faster; 87.7% lower |
| Observed time range | 3074-3705 ns/op | 332.9-536.8 ns/op | Consistent separation across all runs |
| Allocated bytes | 5520 B/op | 72 B/op | 98.7% lower |
| Allocations | 24 allocs/op | 3 allocs/op | 87.5% lower |

Raw results:

| Run | Before ns/op | After ns/op |
|---:|---:|---:|
| 1 | 3335 | 385.6 |
| 2 | 3074 | 378.8 |
| 3 | 3167 | 361.2 |
| 4 | 3231 | 431.1 |
| 5 | 3090 | 447.5 |
| 6 | 3705 | 385.2 |
| 7 | 3264 | 536.8 |
| 8 | 3295 | 416.0 |
| 9 | 3210 | 350.7 |
| 10 | 3276 | 332.9 |

Every before run reported `5520 B/op` and `24 allocs/op`. Every after run
reported `72 B/op` and `3 allocs/op`.

## Verification against the merged `main`

The merged revision `5a1612d` was checked out independently and tested again:

```bash
go test ./...
go test -race ./...
go test -run '^$' -bench '^BenchmarkNanoid$' \
  -benchmem -benchtime=1s -count=5
```

Both correctness commands passed. The five benchmark results from the merged
revision were:

```text
435.6 ns/op    72 B/op    3 allocs/op
404.6 ns/op    72 B/op    3 allocs/op
414.0 ns/op    72 B/op    3 allocs/op
381.7 ns/op    72 B/op    3 allocs/op
367.8 ns/op    72 B/op    3 allocs/op
```

Their mean was 400.7 ns/op, consistent with the ten-run after-upgrade mean of
402.6 ns/op.

## Why the default path is faster

The previous implementation converted the 64-character default alphabet to a
new `[]rune` value inside the loop that emits each character. A default ID has
21 characters, so this created repeated short-lived allocations.

The upgraded implementation keeps the Nano ID 6 URL alphabet as an ASCII
constant, builds the result in a byte slice, and directly indexes the alphabet:

```go
id := make([]byte, size)
for i := 0; i < size; i++ {
    id[i] = URLAlphabet[randomBytes[i]&63]
}
```

It still uses `crypto/rand`; the performance improvement does not replace the
secure random source with a non-secure generator or introduce shared mutable
random state.

## Interpretation and limits

- The measured improvement applies to the default `New()` path, not to an
  entire application using this package.
- `Generate(customAlphabet, size)` was covered by correctness and distribution
  tests but was not included in this before/after microbenchmark.
- Timing varies with hardware, operating system load, Go version, and entropy
  source performance. Allocation counts are the more deterministic signal.
- The batches were executed serially and were not analyzed with `benchstat`.
  The observed gap is nevertheless much larger than the run-to-run variance.
- Performance does not provide uniqueness by itself. Applications must still
  enforce appropriate uniqueness constraints and retry on an actual collision.
