package main

type UserService struct {
	OrderService *OrderService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this UserService) GetUser(id string) string {
	return "User-" + id
}

func (this UserService) GetOrder(id string) string {
	return this.OrderService.GetOrder(id)
}
