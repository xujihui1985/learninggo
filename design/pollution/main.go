package main

type Server interface {
	Start() error
	Stop() error
	Wait() error
}

//
type server struct {
}

func (*server) Start() error {
	return nil
}
func (*server) Stop() error {
	return nil
}
func (*server) Wait() error {
	return nil
}

func NewServer() Server {
	return &server{}
}

func main() {
	s := NewServer()
	s.Start()
	s.Wait()
	s.Stop()
}

// Smells:
// The package declares an interface that matches the entire API of its own concrete type
// The interface is exported but the concrete type is unexported
// the factor function returns the interface value with the unexported concreate type value
// the interface can be removed and nothing changes for the user of the API
// the interface is not decoupling the API from change

// NOTES:
// * Use an interface:
// When users of the API need to provide an implementation detail
// When API's have multiple implementations that need to be maintained
// When parts of the API that can change have been identified and required decoupling

// Question an interface:
// when its only purpose is for writing testable api
// when it's not providing support for the api to decouple from change
// when it's not clear how the interface makes the code better
