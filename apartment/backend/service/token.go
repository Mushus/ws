package service

func NewAuth() AuthService {
	return &authService{}
}

type AuthService interface {
	CreateToken() string
}

type authService struct{}

func (a *authService) CreateToken() string {
	return "hoge"
}
