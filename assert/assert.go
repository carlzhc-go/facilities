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

func Assert(i any) {
	var v reflect.Value = reflect.ValueOf(i)
	if v.Kind() == reflect.Struct {
		return // Ignore stuct type
		// panic(Errorf("Cannot Assert on struct type %T: %v\n", i, i))
	}

	// Indirect returns the value that v points to. If v is a nil pointer,
	// Indirect returns a zero Value, otherwise, Indirect returns v.
	var iv reflect.Value = reflect.Indirect(v)
	if iv.IsZero() {
		panic(Errorf("Zero value of type %T: %v\n", i, iv))
	}
}
