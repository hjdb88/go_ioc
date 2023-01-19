package inject

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
