package dynamic

import (
	"encoding/json"
	"reflect"
)

func Parse(data []byte, ptr interface{}) (err error) {
	if err = json.Unmarshal(data, ptr); err != nil {
		return err
	}
	return parseDynamic(reflect.ValueOf(ptr), nil, "")
}

func parseDynamic(v reflect.Value, dynFielder DynamicFielder, dynFieldName string) (err error) {
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return parseDynamic(v.Elem(), dynFielder, dynFieldName)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err = parseDynamic(v.Index(i), dynFielder, dynFieldName); err != nil {
				return err
			}
		}
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			if err = parseDynamic(iter.Value(), dynFielder, dynFieldName); err != nil {
				return err
			}
		}
	case reflect.Struct:
		if v.CanAddr() {
			return nil
		}

		if dynFielder != nil && v.Type() == DynamicType {
			ptr := dynFielder.NewDynamicField(dynFieldName)
			if ptr != nil {
				dynValue := v.Addr().Interface().(*Type)
				if err = dynValue.Unmarshal(ptr); err != nil {
					return err
				}
				dynValue.Value = ptr
				return nil
			}
		}

		dynFielder, ok := v.Addr().Interface().(DynamicFielder)
		if !ok {
			return nil
		}

		typ := v.Type()
		for i := 0; i < v.NumField(); i++ {
			sf := typ.Field(i)
			field := v.Field(i)

			if !field.CanSet() {
				continue
			}

			if err = parseDynamic(field, dynFielder, sf.Name); err != nil {
				return err
			}
		}
	}
	return nil
}
