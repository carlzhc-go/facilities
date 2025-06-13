/** -*- mode: go; -*-
 * $Id$
 *
 * Utilities of functions similar to Perl's
 */

package perl

import (
	pcre2 "github.com/Jemmic/go-pcre2"
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

// String subsituation, flags in [msixpodualngcer] (ref: https://perldoc.perl.org/perlop#s/PATTERN/REPLACEMENT/msixpodualngcer)
func S(subject, pattern, replacement, flags string) string {
	var reflags uint32
	var subflags uint32
	var replace_all bool
	for _, flag := range flags {
		switch flag {
		// ^ and $ match start and end of lines
		case 'm':
			reflags |= pcre2.MULTILINE

		// . matches any character, including newlines.
		case 's':
			reflags |= pcre2.DOTALL

		// Case-insensitive matching
		case 'i':
			reflags |= pcre2.CASELESS

		// Allows whitespace and comments in the pattern.
		case 'x':
			reflags |= pcre2.EXTENDED

		// Makes quantifiers lazy (ungreedy).
		case '?':
			reflags |= pcre2.UNGREEDY

		// Enables UTF-8 handling for patterns and strings
		case 'u':
			reflags |= pcre2.UTF

		case 'g':
			replace_all = true
		case 'r': // ignored, always return a new string
		case 'p', 'o', 'd', 'a', 'l', 'n', 'c', 'e':
			Die("Not implemented '%c' flag", flag)
		default:
			Die("Unknow flag '%c'", flag)
		}
	}

	var r *pcre2.Regexp = pcre2.MustCompile(pattern, reflags)
	if replace_all {
		return r.ReplaceAllString(subject, replacement, subflags)
	} else {
		return replaceString(r, subject, replacement, subflags)
	}

}

func replace(re *pcre2.Regexp, bytes, repl []byte, flags uint32) []byte {
	m := re.Matcher(bytes, flags)
	defer m.Free()
	r := []byte{}
	if m.Matches() {
		loc := m.Index()
		r = append(r, bytes[0:loc[0]]...)
		r = append(r, repl...)
		r = append(r, bytes[loc[1]:]...)
	} else {
		r = append(r, bytes...)
	}
	return r
}

// ReplaceAllString is equivalent to ReplaceAll with string return type.
func replaceString(re *pcre2.Regexp, in, repl string, flags uint32) string {
	return string(replace(re, []byte(in), []byte(repl), flags))
}
