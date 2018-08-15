package deepequal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeepEqual(t *testing.T) {
	f1 := foo{
		value: 22,
		name:  "sean",
		arr:   []string{"a", "b"},
		mp: map[string]bar{
			"hello": bar{
				value: "world",
			},
		},
	}

	f2 := foo{
		name:  "sean",
		value: 22,
		arr:   []string{"a", "b"},
		mp: map[string]bar{
			"hello": bar{
				value: "world",
			},
		},
	}

	res := reflect.DeepEqual(f1, f2)
	assert.Equal(t, res, true, "f1 should eq to f2")
}
