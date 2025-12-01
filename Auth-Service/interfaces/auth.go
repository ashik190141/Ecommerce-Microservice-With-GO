package interfaces

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRepository interface {
	CreateUser(name string, email string, password string) (User, string, error)
	GetUserByEmail(email string) (bool, error)
	LoginUser(email string, password string) (User, error)
}