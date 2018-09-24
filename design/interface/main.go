package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct {
}

func (bike) Move() {
	fmt.Println("moving the bike")
}

func (bike) Lock() {
	fmt.Println("locking the bike")
}

func (bike) Unlock() {
	fmt.Println("unlocking the bike")
}

func assignMoveLockerToMover() {
	var ml MoveLocker
	var m Mover

	ml = bike{}
	m = ml

	m.Move()
}

func assignMoverToMoverLocker() {
	// var m Mover
	// var ml MoveLocker

	// m = bike{}
	// Mover does not implement movelocker (missing lock method)
	// ml = m

	// m.Move()
}

func typeAssertion() {
	var m Mover
	var ml MoveLocker
	m = bike{}

	// if we are pretty sure the Mover interface type stored are bike,
	// we can perform a type assertion against the mover interface value to access
	// a COPY of the concrete type value of type bike that was stored inside of it.
	// then assign the COPY of the concrete type to the MoveLocker interface
	b, ok := m.(bike)
	if !ok {
		return
	}
	ml = b
	ml.Move()
	ml.Lock()
}

type cloud struct{}

func (cloud) String() string {
	return "Big Data"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// mvs := []fmt.Stringer{
	// 	cloud{},
	// }

	for i := 0; i < 10; i++ {
		rn := rand.Int()
		fmt.Println(rn)
	}

}
