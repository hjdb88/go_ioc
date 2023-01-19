package inject

import "reflect"

type BeanMapper map[reflect.Type]reflect.Value

func (this BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)

	if t.Kind() != reflect.Ptr {
		panic("required ptr object")
	}

	this[t] = reflect.ValueOf(bean)
}

func (this BeanMapper) get(bean interface{}) reflect.Value {
	t := reflect.TypeOf(bean)
	if v, ok := this[t]; ok {
		return v
	}
	return reflect.Value{}
}
