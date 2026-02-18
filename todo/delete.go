package todo

func (svc *service) Delete(todoID int) error {
	return svc.todoRepo.Delete(todoID)

}
