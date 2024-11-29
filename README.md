# GO strings.Repeat Benchmark
```console
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/sumomo-99/go-strings.Repeat-benchmark
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkRepeat-8               11474935               103.3 ns/op
BenchmarkStringsRepeat-8        31269271                35.19 ns/op
PASS
ok      github.com/sumomo-99/go-strings.Repeat-benchmark 2.822s
```

## Usage
```console
$ go test -bench=.
```
## Requirements
- go

## License
MIT
