package main

import "fmt"

func main() {
	userService := NewUserService(NewOrderService())
	fmt.Println(userService.GetUser("1"))
	fmt.Println(userService.GetOrder("2"))
}
