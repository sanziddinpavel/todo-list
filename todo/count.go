package todo

func (svc *service) Count() (int64, error) {
	return svc.todoRepo.Count()
}
