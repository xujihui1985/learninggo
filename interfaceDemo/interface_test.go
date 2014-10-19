package interfaceDemo

import (
    "testing"
    "fmt"
)

func TestErrorTest(t *testing.T) {
    _, err := ErrorTest()
    if err != nil {
        t.Errorf(err.Error())
    } 
}

func TestInteger(t *testing.T) {
    var b Integer = 2
    var a *Integer = &b
    fmt.Printf("a = %d\n", *a)
    a.Add(2)
    fmt.Printf("a = %d\n", *a)
}

func TestFoo(t *testing.T) {
    var f Foo = Foo{}
    f.Base = &Base { Name:"sean" }
    f.string = "adfsdsaf"

    f.Log()
    fmt.Printf("name is %s address is %s\n", f.Base.Name, f.string)
}

