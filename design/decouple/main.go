package main

import (
	"fmt"
	"io"
	"time"
)

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type PullStorer interface {
	Puller
	Storer
}

// Xenia is a system we need to pull data from
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// We need struct and method when we need to implement interface
func (*Xenia) Pull(d *Data) error {
	return nil
}

type Solar struct{}

func (*Solar) Store(d *Data) error {
	return nil
}

// We should embed for behaiver
type System struct {
	Xenia
	Solar
}

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(p Storer, data *Data) error {
	return p.Store(data)
}

func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(ps, data)
		if i > 0 {
			if err := store(ps, &data[0]); err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Xenia: Xenia{},
		Solar: Solar{},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}

}
