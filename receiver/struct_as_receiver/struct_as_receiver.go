package receiver

import "fmt"

type foo struct {
	name string
}

func (f foo) receiverAsValue() string {
	return fmt.Sprintf("%p\n", &f.name)
}

func (f *foo) receiverAsPointer() string {
	return fmt.Sprintf("%p\n", &f.name)
}
