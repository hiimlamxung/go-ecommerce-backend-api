package services

import "github.com/hiimlamxung/go-ecommerce-backend-api/internal/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetInfoUser(id int) string {
	return us.userRepository.GetInfoUser(id)
}
