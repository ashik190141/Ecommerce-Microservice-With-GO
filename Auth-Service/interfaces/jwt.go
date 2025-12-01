package interfaces

type JWT struct {
	Token       string
}

type JwtService interface {
	GenerateToken(email string) (JWT, error)
	ValidateToken(token string) (bool, error)
}