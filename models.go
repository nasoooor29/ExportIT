package ExportIT

import "reflect"

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

func isError(k reflect.Type) bool {
	return k.Implements(errorInterface)
}
