package main

type UserService struct {
	orderService *OrderService
}

func NewUserService(orderService *OrderService) *UserService {
	return &UserService{orderService: orderService}
}

func (this UserService) GetUser(id string) string {
	return "User-" + id
}

func (this UserService) GetOrder(id string) string {
	return this.orderService.GetOrder(id)
}
