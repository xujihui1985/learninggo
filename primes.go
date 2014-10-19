package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

var goal int

func primetask(c chan int) {
    p := <-c

    if p > goal {
        os.Exit(0)
    }

    fmt.Println("p is ", p)

    nc := make(chan int)

    go primetask(nc)

    for {
        i := <-c
        //fmt.Println("i is ", i)

        if i%p != 0 {
            fmt.Println("nc is ", i)
            nc <- i
        }
    }
}

func main() {
    flag.Parse()

    args := flag.Args()
    if args != nil && len(args) > 0 {
        var err error
        goal, err = strconv.Atoi(args[0])
        if err != nil {
            goal = 100
        }
    } else {
        goal = 100
    }
    fmt.Println("goal = ", goal)
    c := make(chan int)
    go primetask(c)
    fmt.Println("start task")
    for i := 2 ;; i++ {
        fmt.Println("in the loop", i)
        c <- i
    }
}
