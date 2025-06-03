/** -*- mode: go; -*-
  * $Id$
  *
  * Utilities of functions similar to Perl's
 */

package perl

import (
	"log"
	"os"
)

type Perl struct {
	Logger *log.Logger
}

var pl *Perl

func init() {
	pl = New()
}

func New() *Perl {
	p := new(Perl)
	p.Logger = log.New(os.Stdout, "", 0)
	return p
}

// Default returns the standard debug logger used by the package-level output functions.
func Default() *Perl {
	var p Perl = *pl
	return &p
}

func Die(msg string, args ...any) {
	pl.Die(msg, args...)
}

func Say(args ...any) {
	pl.Say(args...)
}

func (p *Perl) Die(msg string, args ...any) {
	p.Logger.Panicf("!!! "+msg+", abort", args...)
}

func (p *Perl) Say(args ...any) {
	p.Logger.Println(args...)
}
