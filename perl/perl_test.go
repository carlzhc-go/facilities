package perl

import "testing"
// import "os"
import "strings"
import "bytes"
import "log"

func TestPerlPackage(t *testing.T) {
	msg := "Test perl function"
	wb := bytes.NewBuffer(nil)

	check := func (fn, exp string) {
		got := strings.TrimSpace(wb.String())
		if got == exp {
			t.Logf(`Passed testing %s, [%d]"%v"`, fn, len(exp), exp)
		} else {
			t.Errorf(`Failed testing %s, expected [%d]"%v", got [%d]"%v"`,
				fn, len(exp), exp, len(got), got)
		}
		wb.Reset() // Empty the buf
	}

	p := New()
	p.Logger = log.New(wb, "", 0)

	p.Say(msg)
	check("Say", msg)

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf(`No panic happened`)
			} else {
				t.Logf(`Panic happened`)
			}
			check("Die", "!!! "+msg+", abort")
		}()
		p.Die(msg)
	}()
}
