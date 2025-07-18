package assert

import (
	. "fmt"
	"reflect"
)

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}

	return v
}

// Assert will panic if the argument is zero or it's pointing to zero
func Assert(i any) {
	var v = reflect.ValueOf(i)
	if v.IsZero() {
		panic(Errorf("Zero value of type %T: %v\n", v, v))
	}

	var iv reflect.Value = reflect.Indirect(v)
	if iv.IsZero() {
		panic(Errorf("Zero value of type %T: %v\n", i, iv))
	}
}
