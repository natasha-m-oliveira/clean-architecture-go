package utils

import "reflect"

func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Copy(dst, src any) {
	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	dstType := dstVal.Type()

	for i := 0; i < dstType.NumField(); i++ {
		field := dstType.Field(i)
		srcField := srcVal.FieldByName(field.Name)
		dstField := dstVal.FieldByName(field.Name)

		if srcField.IsValid() && dstField.IsValid() && srcField.CanSet() {
			if !reflect.DeepEqual(srcField.Interface(), dstField.Interface()) {
				dstField.Set(srcField)
			}
		}
	}
}
