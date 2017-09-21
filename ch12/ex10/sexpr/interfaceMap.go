package sexpr

import "reflect"

var interfaceMap map[string]reflect.Type

func init() {
	interfaceMap = make(map[string]reflect.Type)
}
