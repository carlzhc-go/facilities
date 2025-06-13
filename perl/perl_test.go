package perl

import "testing"

// import "os"
import "strings"
import "bytes"
import "log"

func TestPerlPackage(t *testing.T) {
	msg := "Test perl function"
	wb := bytes.NewBuffer(nil)

	check := func(fn, exp string) {
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

func TestPerlSubst(t *testing.T) {
	samples := [][]string{
		// subject, pattern, replacement, flags, expectation
		{"abcabc", "a", "x", "", "xbcabc"},
		{"abcabc", "a", "x", "g", "xbcxbc"},
		{"abcAbc", "A", "x", "i", "xbcAbc"},
		{"abcAbc", "A", "x", "ig", "xbcxbc"},
		{"abcAbc", "a.*b", "x", "ig", "xc"},
		{"abcAbc", "a.*b", "x", "?", "xcAbc"},
		{"abc\nAbc", "a.*b", "x", "ig", "xc\nxc"},
		{"abc\nAbc", "a.*b", "x", "igs", "xc"},
	}

	for _, s := range samples {
		got := S(s[0], s[1], s[2], s[3])
		exp := s[4]
		if got == exp {
			t.Logf(`Passed test %v => %v`, s[0:4], got)
		} else {
			t.Errorf(`Failed test %v => %v, got %v`, s[0:4], s[4], got)
		}
	}
}
