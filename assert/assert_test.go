package assert

import (
	"errors"
	tt "testing"
)

func TestMust(t *tt.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Failed test on errors")
			}
		}()
		Must(1, errors.New("Failed Must 1 test"))
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Failed test on no errors")
			}
		}()
		if Must(1, nil) != 1 {
			t.Errorf("Failed test on return value")
		}
	}()

}

func expectAssertPanic(t *tt.T, i any) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Failed test on %T %v", i, i)
		}
	}()

	switch v := i.(type) {
	default:
		Assert(v)
	}
}

func TestAssert(t *tt.T) {
	i := 0
	Assert(1)
	expectAssertPanic(t, i)
	expectAssertPanic(t, &i)
	Assert(0.1)
	expectAssertPanic(t, 0.0)
	Assert("hello")
	expectAssertPanic(t, "")
	Assert(true)
	expectAssertPanic(t, false)
	expectAssertPanic(t, nil)

	b := true
	Assert(&b)
	b = !b
	expectAssertPanic(t, &b)

	var o struct{}
	Assert(o)

	Assert(&o)
}
