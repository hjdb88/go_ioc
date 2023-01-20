package main

import (
    "fmt"
    "github.com/hjdb88/go_ioc/v3/inject"
)

func main() {
	inject.BeanFactory.Set(NewOrderService())
	userService := NewUserService()
	inject.BeanFactory.Apply(userService)
	fmt.Println(userService.GetUser("1"))
	fmt.Println(userService.GetOrder("2"))
}
