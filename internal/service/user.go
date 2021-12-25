package service

import "context"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Create(ctx context.Context) {

}
