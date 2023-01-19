package main

import "fmt"

func main() {
	user := NewUserService().GetUser("1")
	fmt.Println(user)
	order := NewOrderService().GetOrder("2")
	fmt.Println(order)
}
