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
			t.Errorf("Failed test on (%T)%v", i, i)
		} else {
			t.Logf("Passed panic test on (%T)%v", i, i)
		}
	}()

	switch v := i.(type) {
	default:
		Assert(v)
	}
}

func expectAssertPass(t *tt.T, i any) {
	defer func() {
		if r := recover(); r == nil {
			t.Logf("Passed test on (%T)%v", i, i)
		} else {
			t.Errorf("Failed test on (%T)%v", i, i)
		}
	}()

	switch v := i.(type) {
	default:
		Assert(v)
	}
}

func TestAssert(t *tt.T) {
	i := 0
	expectAssertPass(t, 1)
	expectAssertPanic(t, i)
	expectAssertPanic(t, &i)
	expectAssertPass(t, 0.1)
	expectAssertPanic(t, 0.0)
	expectAssertPass(t, "hello")
	expectAssertPanic(t, "")
	expectAssertPass(t, true)
	expectAssertPanic(t, false)
	expectAssertPanic(t, nil)

	b := true
	expectAssertPass(t, &b)
	b = !b
	expectAssertPanic(t, &b)

	var o struct{}
	expectAssertPanic(t, o)
	expectAssertPanic(t, &o)

	var s []byte
	expectAssertPanic(t, s)
	expectAssertPanic(t, &s)

	var m map[any]any
	expectAssertPanic(t, m)
	expectAssertPanic(t, &m)

	m = make(map[any]any)
	expectAssertPass(t, m)
	expectAssertPass(t, &m)

	var c chan int
	expectAssertPanic(t, c)
	expectAssertPanic(t, &c)

	c = make(chan int)
	expectAssertPass(t, c)
	expectAssertPass(t, &c)

}
