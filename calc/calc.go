package main

import "os"
import "fmt"
import "strconv"

var Usage = func() {
    fmt.Println("Usage: calc command [arguments] ...")
}

func main() {
    args := os.Args[1:]
    if args == nil || len(args) < 2 {
        Usage()
        return
    }
    switch args[0] {
        case "add":
            if len(args) != 3 {
                fmt.Println("Usage: calc add <integer1><integer2>")
                return
            }
            v1, err1 := strconv.Atoi(args[1])
            v2, err2 := strconv.Atoi(args[2])
            if(err1 != nil || err2 != nil) {
                fmt.Println("Usage: calc add <integer1><integer2>")
            }
            ret := v1 + v2
            fmt.Println("Result: ", ret)
        default:
            Usage()
    }
}
