package main

import (
	"fmt"

	"github.com/hjdb88/go_ioc/v2/inject"
)

func main() {
	inject.BeanFactory.Set(NewOrderService())
	orderService := inject.BeanFactory.Get((*OrderService)(nil))
	fmt.Println(orderService)
	inject.BeanFactory.Set(NewUserService(orderService.(*OrderService)))
	userServiceInterface := inject.BeanFactory.Get((*UserService)(nil))
	fmt.Println(userServiceInterface)
	userService := userServiceInterface.(*UserService)
	fmt.Println(userService.GetUser("1"))
	fmt.Println(userService.GetOrder("2"))
}
