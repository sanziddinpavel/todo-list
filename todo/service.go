package todo

type service struct {
	todoRepo TodoRepo
}

func NewService(todoRepo TodoRepo) Service {
	return &service{
		todoRepo: todoRepo,
	}

}
