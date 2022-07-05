package structBytes

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func (o *object) Write(obj interface{}) bool {
	o.Buffer = &bytes.Buffer{}
	return o.WriteValue(obj, 0)
}

func (o *object) WriteValue(obj interface{}, depth int) (ok bool) {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Interface:
	case reflect.Ptr:
		if !o.WriteValue(v.Elem().Interface(), depth+1) {
			return
		}
	case reflect.Struct:
		l := v.NumField()
		for i := 0; i < l; i++ {
			if !o.WriteValue(v.Field(i).Interface(), depth+1) {
				return
			}
		}

	case reflect.Slice, reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			if !o.WriteValue(v.Index(i).Interface(), depth+1) {
				return
			}
		}

	case reflect.Int:
		i := int32(obj.(int))
		if int(i) != obj.(int) {
			return c.Error("Int does not fit into int32")
		}
		if !c.Error(binary.Write(o.Buffer, binary.BigEndian, i)) {
			return
		}

	case reflect.Bool:
		b := uint8(0)
		if v.Bool() {
			b = 1
		}
		if !c.Error(binary.Write(o.Buffer, binary.BigEndian, b)) {
			return
		}

	default:
		return c.Error(binary.Write(o.Buffer, binary.BigEndian, obj))
	}
	return true
}
