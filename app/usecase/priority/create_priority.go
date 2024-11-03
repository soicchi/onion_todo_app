package priority

import (
	"onion_todo_app/domain/priority"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type CreatePriorityUseCase struct {
	PriorityRepository priority.PriorityRepository
}

func NewCreatePriorityUseCase() *CreatePriorityUseCase {
	return &CreatePriorityUseCase{
		PriorityRepository: repository.PriorityRepository{},
	}
}

type CreatePriorityInputDTO struct {
	Level string
}

func (uc *CreatePriorityUseCase) Execute(ctx echo.Context, dto CreatePriorityInputDTO) error {
	priority, err := priority.NewPriority(dto.Level)
	if err != nil {
		return err
	}

	if err := uc.PriorityRepository.Create(ctx, priority); err != nil {
		return err
	}

	return nil
}
