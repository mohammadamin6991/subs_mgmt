package service

import (
	"context"

	"gitlab.amin.run/general/project/subs-mgmt/authentication/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (userService *UserService) Create(ctx context.Context, u *repository.User) (*repository.User, error) {
	user, err := userService.UserRepository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *UserService) GetByEmail(ctx context.Context, email string) (*repository.User, error) {
	user, err := userService.UserRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *UserService) List(ctx context.Context) ([]repository.User, error) {
	users, err := userService.UserRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService *UserService) Update(ctx context.Context, u *repository.User) (*repository.User, error) {
	user, err := userService.UserRepository.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService *UserService) Delete(ctx context.Context, id int64) error {
	return userService.UserRepository.Delete(ctx,id)
}



type SessionService struct {
	SessionRepository *repository.SessionRepository
}

func NewSessionService(sessionRepository *repository.SessionRepository) *SessionService {
	return &SessionService{
		SessionRepository: sessionRepository,
	}
}


func (s *SessionService) CreateSession(ctx context.Context, se *repository.Session) (*repository.Session, error) {
	return s.SessionRepository.CreateSession(ctx, se)
}

func (s *SessionService) GetSession(ctx context.Context, id string) (*repository.Session, error) {
	return s.SessionRepository.GetSession(ctx, id)
}

func (s *SessionService) RevokeSession(ctx context.Context, id string) error {
	return s.SessionRepository.RevokeSession(ctx, id)
}

func (s *SessionService) DeleteSession(ctx context.Context, id string) error {
	return s.SessionRepository.DeleteSession(ctx, id)
}
