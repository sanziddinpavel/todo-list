package user

import (
	"Todo-list/domain"
	userHandler "Todo-list/rest/handler/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}
