### interface

```golang
var r reader
```

interface value are referece type in golang, interface value are two word data struct,
two word represent two different pointer, the first word point to the type of the concreat type of the struct
and the second word point to the copy of the struct

```golang

type reader interface {
  read(b []byte) (int, error)
}

type file struct {
  filename string
}

func (file) read(b []byte) (int, error) {
  s := "hello"
  copy(b, s)
  return len(s), nil
}
```