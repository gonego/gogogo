#COMPONENT OF TEST SUITE
- Test
- Benchmark
- Testable Example

# TEST
```
TestXxx(t *testing.T)
```

#BENCHMARK
```
BenchmarkXxx(b *testing.B)
```

#TESTABLE EXAMPLE
```
ExampleXxx()
```

# Commands

```
godoc -http=:6060   // note: 6060 is the port number

go test             // use go test -v for verbose output
go test -bench .
go test -cover
go test -coverprofile c.out

go tool cover -html=c.out
```








































