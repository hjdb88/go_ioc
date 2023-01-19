package main

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this UserService) GetUser(id string) string {
	return "User-" + id
}
