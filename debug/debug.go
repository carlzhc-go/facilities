/** -*- mode: go; -*-
  * $Id$
  *
  * Utilities of functions for debugging
 */

package debug

import (
	"io"
	"log"
	"os"
	"runtime"
)

type Debug struct {
	Enabled bool
	WithCaller bool
	Logger *log.Logger
}

var debugger *Debug

// package functions
func Println(args ...any) {
	debugger.Println(args...)
}

func Print(args ...any) {
	debugger.Print(args...)
}

func Printf(format string, v ...any) {
	debugger.Printf(format, v...)
}

func SetLogger(l *log.Logger) {
	debugger.Logger = l
}

// New creates a new debug logger. The out variable sets the destination to which log
// data will be written. The prefix appears at the beginning of each generated
// log line, or after the log header if the Lmsgprefix flag is provided. The
// flag argument defines the logging properties.
func New(out io.Writer, prefix string, flag int) *Debug {
	var dbg = new(Debug)
	dbg.Logger = log.New(out, prefix, flag)
	return dbg
}

// Default returns the standard debug logger used by the package-level output functions.
func Default() *Debug {
	var dbg Debug = *debugger
	return &dbg
}

func (d *Debug) SetLogger(l *log.Logger) {
	d.Logger = l
}

func (d *Debug) Print(args ...any) {
	if !d.Enabled { return }

	if d.WithCaller {
		// Get caller's Name
		var caller = "UNKNOWN"

		pc, _, _, ok := runtime.Caller(1)
		if ok {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				caller = fn.Name()
			}
		}

		caller = "["+caller+"] "
		args = append([]any{caller}, args...)
	}

	d.Logger.Print(args...)
}

func (d *Debug) Println(args ...any) {
	if !d.Enabled { return }

	if d.WithCaller {
		// Get caller's Name
		var caller string = "UNKNOWN"
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				caller = fn.Name()
			}
		}

		caller = "["+caller+"] "
		args = append([]any{caller}, args...)
	}
	d.Logger.Println(args...)
}

func (d *Debug) Printf(format string, args ...any) {
	if !d.Enabled { return }

	if d.WithCaller {
		// Get caller's Name
		var caller string = "UNKNOWN"
		pc, _, _, ok := runtime.Caller(1)
		if ok {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				caller = fn.Name()
			}
		}

		format = "["+caller+"] "+format
	}

	d.Logger.Printf(format, args...)
}

func init() {
	debugger = New(os.Stderr, "> ", 0)
}
