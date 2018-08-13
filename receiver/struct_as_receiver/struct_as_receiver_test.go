package receiver

import (
	"fmt"
	"testing"
)

func Test_foo_receiverAsValue(t *testing.T) {
	f := &foo{
		name: "bar",
	}

	nameptr := fmt.Sprintf("%p\n", &f.name)

	valueNameptr := f.receiverAsValue()

	pointNamerptr := f.receiverAsPointer()

	if nameptr == valueNameptr {
		t.Errorf("nameptr %s and valueNameptr %s are expected to be different", nameptr, valueNameptr)
	}

	if nameptr != pointNamerptr {
		t.Errorf("nameptr %s and valueNameptr %s are expected to be same", nameptr, pointNamerptr)
	}
}
