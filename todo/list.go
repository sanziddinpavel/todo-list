package todo

import "Todo-list/domain"

func (svc *service) List(page, limit int64) ([]*domain.Todos, error) {
	return svc.todoRepo.List(page, limit)
}
