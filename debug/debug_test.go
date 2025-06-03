package debug

import "testing"
// import "os"
import "strings"
import "bytes"

func TestDebugPackage(t *testing.T) {
	msg := "Test debug message"
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

	dbg := New(wb, "", 0)
	dbg.Enabled = true

	dbg.Print(msg)
	check("Print", msg)

	dbg.Println(msg)
	check("Println", msg)

	dbg.Printf("%s", msg)
	check("Printf", msg)

	// Test with prefix
	dbg = New(wb, "> ", 0)
	dbg.Enabled = true

	dbg.Print(msg)
	check("Print with prefix", "> "+msg)
}
