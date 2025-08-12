package user_usecase

import (
	"context"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/entity/user_entity"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/internal_error"
)

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

type UserOutputDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCaseInterface interface {
	FindUserById(
		ctx context.Context, id string) (*UserOutputDto, *internal_error.InternalError)
}

func (u *UserUseCase) FindUserById(
	ctx context.Context, id string) (*UserOutputDto, *internal_error.InternalError) {

	userEntity, err := u.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDto{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
