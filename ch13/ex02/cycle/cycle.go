// Package cycle は、循環構造を持つかどうかを報告します。
package cycle

import (
	"reflect"
	"unsafe"
)

func cycle(x reflect.Value, seen map[comparison]bool) bool {
	// cycle check
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparison{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if cycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if cycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false

	default:
		return false
	}
}

// Cycle は、循環構造を持つかどうかを報告します。
func Cycle(x interface{}) bool {
	seen := make(map[comparison]bool)
	return cycle(reflect.ValueOf(x), seen)
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}
