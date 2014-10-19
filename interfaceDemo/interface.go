package interfaceDemo

import (
    "errors"
    "fmt"
)

type PathError struct {
    Op string
    Path string
    Err error
}

type Integer int

type Base struct {
    Name string
}
type Base2 struct {
    Address string
}

func (base *Base) Log() {
    fmt.Printf("log from base %s\n", base.Name);
}

type Foo struct {
    *Base
    Base2
    string
}


func (a *Integer) Add(b Integer) {
    *a += b
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + " " + e.Err.Error()
}

func ErrorTest() (ret string, err error) {
    if 1 == 1 {
        err = &PathError{
            Op: "add",
            Path: "Path",
            Err: errors.New("path error"),
        }
        ret = ""
        return
    } 
    return "success", nil
}

