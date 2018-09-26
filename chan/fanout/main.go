package main

import (
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	id  int
	op  string
	err error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const routines = 10
	const inserts = routines * 2

	ch := make(chan result, inserts)

	waitInserts := inserts

	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)
			ch <- updateTrans(id)
		}(i)
	}

	for waitInserts > 0 {
		r := <-ch
		fmt.Printf("N: %d ID: %d OP: %s ERR: %v\n", waitInserts, r.id, r.op, r.err)
		waitInserts--
	}
}

func insertUser(id int) result {
	return result{
		id:  id,
		op:  "insert",
		err: nil,
	}
}

func updateTrans(id int) result {
	return result{
		id:  id,
		op:  "update",
		err: nil,
	}
}
