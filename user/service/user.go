package service

import (
	"user-api/domain"
	"user-api/model/request"
	"user-api/model/response"
	"user-api/user_repository"
)

type userService struct {
	userRepository user_repository.UserRepository
}

type UserService interface {
	GetUser(userId string) (*response.UserResponse, error)
	SaveUser(request request.UserCreateRequest) error
}

func NewUserService(userRepository user_repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetUser(userId string) (*response.UserResponse, error) {
	userEntity, err := s.userRepository.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return &response.UserResponse{
		Name:        userEntity.Name,
		Surname:     userEntity.Surname,
		PhoneNumber: userEntity.PhoneNumber,
		Email:       userEntity.Email,
	}, nil
}

func (s *userService) SaveUser(request request.UserCreateRequest) error {
	err := s.userRepository.AddUser(&domain.UserEntity{
		Name:        request.Name,
		Surname:     request.Surname,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		UserId:      request.UserId,
	})
	if err != nil {
		return err
	}
	return nil
}
