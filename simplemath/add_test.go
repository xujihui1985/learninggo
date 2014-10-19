package simplemath

import (
    "testing"
    "fmt"
)

func TestAdd(t *testing.T) {
    r := Add(1,2)
    if r != 3 {
        t.Errorf("Add(1,2) failed, Got %d, expected 3.", r)
    }
}

func TestConst(t *testing.T) {

   fmt.Println("const is", Sunday) 
   t.Errorf("const is ", numberOfDays)
   fmt.Println("const is", Monday) 
}

func TestArray(t *testing.T) {
    arr := GetResult()
    for i := 0; i < len(arr); i++ {
        fmt.Println("element is: ", i)
    }

    for i, v := range arr {
        fmt.Printf("element [%d] is: %d \n",i, v)
    }
}

func TestArrayToSlice(t *testing.T) {
    arr := [10]int {1,9,3,4,5,6,8}
    slice := createSliceFromArray(arr)
    for i, v := range slice {
        fmt.Printf("element [%d] is: %d \n",i, v)
    }
}

func TestMap(t *testing.T) {
    personMap := createMap()

    for k, v := range personMap {
        fmt.Printf("key is %s personName is %s  address is %s \n",k, v.Name, v.Address)
    }
}

func TestMapContain(t *testing.T) {
    personMap := createMap()
    value, ok := personMap["sean"]
    if ok {
        fmt.Printf("sean exists in the map %s\n", value.Name)
    }
    value, ok = personMap["jack"]
    if ok {
        fmt.Printf("jack exists in the map %s\n", value.Name)
    }

}

func TestMultiReturn(t *testing.T) {
    ret, err := multiReturn(1,-1)
    if(err != nil) {
        fmt.Printf("error throwed %s\n", err.Error())
    } else {
        fmt.Printf("result is %d\n", ret)
    }
}

func TestReturnValue(t *testing.T) {
    fmt.Println("testing...")
    resultfunc := funcAsReturnValue()
    result := resultfunc(100)

    fmt.Println("the result is ",result)

    if result == 100 {
        fmt.Println("success")
    }
}

func TestClosure(t *testing.T) {
    closure()
}
