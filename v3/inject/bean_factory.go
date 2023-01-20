package inject

import (
    "reflect"
)

var BeanFactory *beanFactoryImpl

func init() {
	BeanFactory = &beanFactoryImpl{beanMapper: make(BeanMapper)}
}

type beanFactoryImpl struct {
	beanMapper BeanMapper
}

func (this beanFactoryImpl) Set(beanList ...interface{}) {
	if beanList == nil || len(beanList) == 0 {
		return
	}

	for _, bean := range beanList {
		this.beanMapper.add(bean)
	}
}

func (this beanFactoryImpl) Get(bean interface{}) interface{} {
	if bean == nil {
		return nil
	}
	v := this.beanMapper.get(bean)
	if v.IsValid() {
		return v.Interface()
	}
	return nil
}

func (this beanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}

	v := reflect.ValueOf(bean)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			if getV := this.Get(field.Type); getV != nil {
				v.Field(i).Set(reflect.ValueOf(getV))
			}
		}
	}

	return
}
