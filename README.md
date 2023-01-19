# IoC

## 名词解释

### 控制正转 [Demo](v0/main.go)
程序主动创建"依赖"。

### 控制反转 [Demo](v1/main.go)
程序被动接受"依赖"。

### IOC容器
"依赖"还是需要手动创建，需要通过一个东西来统一创建，这个就是IoC容器。

## 实现

### 基于类型存储 [Demo](v2/main.go)

#### 基本设计
使用Map, key: reflect.Type, value: reflect.Value
```go
type BeanMapper map[reflect.Type]reflect.Value
```

### 处理依赖注入 [Demo](v3/main.go)
