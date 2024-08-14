package builder

import (
	"fmt"
	"reflect"
)

func Fill(src any, opts ...fillerOpt) (any, error) {
	f := &filler{
		stringMax: 10,
	}
	for _, opt := range opts {
		opt(f)
	}
	v := reflect.ValueOf(src)
	return f.fill(v)
}

type filler struct {
	arrayCount   int
	sliceCount   int
	stringMin    int
	stringMax    int
	respectEmpty bool
}

func (f filler) fill(src reflect.Value) (reflect.Value, error) {

	v := src.Elem()
	t := v.Type()

	switch t.Kind() {
	case reflect.Ptr:
		return f.fill(v)
	case reflect.Struct:
		return fillStruct(v)
	case reflect.Slice:
		return fillSlice(v)
	case reflect.Array:
		return fillArray(v)
	default:
		return reflect.Value{
			Kind: reflect.Invalid,
			
		}, fmt.Errorf("unsupported type %s", t.Kind())
	}
}

// Función para establecer valores en los campos del struct
func (f filler) fillStruct(v reflect.Value) (reflect.Value, error) {
	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)
		if fieldValue.CanSet() {
			v = f.fill(fieldValue.Addr().Interface())
		}
	}
	return v, nil
}
