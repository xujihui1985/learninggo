package simplemath

import (
    "errors"
    "fmt"
)

const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    numberOfDays
)

type PersonInfo struct {
    id int
    Name string
    Address string 
}

func Add(a, b int) int {
    return a + b 
}

func GetResult() []int {
    arr := []int {1,2,3,4,5}
    return arr
}

func createSliceFromArray(arr [10]int) []int {
    return arr[:5]
}

func createMap() map[string] PersonInfo {
    var personMap map[string] PersonInfo
    personMap = make(map[string] PersonInfo)

    personMap["sean"] = PersonInfo {
        id: 0,
        Name: "sean",
        Address: "shanghai",
    }
    personMap["anna"] = PersonInfo {
        id: 0,
        Name: "anna",
        Address: "shanghai",
    }

    return personMap
}

func multiReturn(a, b int) (ret int, err error) {
    if a < 0 || b < 0 {
       err = errors.New("a and b can not be negitive value")
       return
    }
    ret = a + b
    return
}

/*return func as return type*/
func funcAsReturnValue() (func(int) int) {
    return func(a int) int {
        return a
    }
}

/*clousure*/
func closure() {
    var j int = 5 /* same as j := 5 */
    a := func()(func()) {
       var i int = 10
       return func() {
           j = j * 2
           fmt.Printf("i is %d j is %d \n", i, j)
       }
    }()
    a()
    a()
    a()
}
