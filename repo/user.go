package repo

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsDone    bool   `json:"is_done"`
}
type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
}

type userRepo struct {
	users []User
}

func NewUserRepo() userRepo {
	return userRepo{}
}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}
	user.ID = len(r.users) + 1

	r.users = append(r.users, user)
	return &user, nil

}
func (r *userRepo) Find(email, password string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}
	return nil, nil
}
