### run benchmarks

```go
go test -run none -bench . -benchtime 3s -benchmem
```

-run none make sure that there is no test function running, only benchmark, none is just a convension that no test should match regex none

```go
go test -run none -bench Sprint/sprint -benchtime 3s -benchmem
```

can also run specific sub test