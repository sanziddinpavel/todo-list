package todo

import "Todo-list/domain"

func (svc *service) Get(todoID int) (*domain.Todos, error) {
	return svc.todoRepo.Get(todoID)

}
