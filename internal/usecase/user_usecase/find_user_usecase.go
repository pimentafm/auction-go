package user_usecase

import (
	"context"

	"github.com/pimentafm/auction-go/internal/entity/user_entity"
	"github.com/pimentafm/auction-go/internal/internal_error"
)

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) UserUserCaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUserCaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError)
}

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
