package main

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (this OrderService) GetOrder(id string) string {
	return "Order-" + id
}
