package assert

import (
	"log"
)

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}

	return v
}

func Assert(i any) {
	die := func(v any) { log.Panicf("Assertion failed for type %T, %v", i, v) }

	switch v := i.(type) {
	case nil:
		die(nil)
	case int:
		if v == 0 {
			die(v)
		}
	case uint:
		if v == 0 {
			die(v)
		}
	case float32:
		if v == 0.0 {
			die(v)
		}
	case float64:
		if v == 0.0 {
			die(v)
		}
	case bool:
		if !v {
			die(v)
		}
	case string:
		if v == "" {
			die(v)
		}
	case *int:
		if *v == 0 {
			die(*v)
		}
	case *uint:
		if *v == 0 {
			die(*v)
		}
	case *float32:
		if *v == 0.0 {
			die(*v)
		}
	case *float64:
		if *v == 0.0 {
			die(*v)
		}
	case *bool:
		if !*v {
			die(*v)
		}
	case *string:
		if *v == "" {
			die(*v)
		}
	}
}
