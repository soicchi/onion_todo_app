package priority

import (
	"onion_todo_app/domain/entity"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type CreatePriorityUseCase struct {
	PriorityRepository entity.PriorityRepository
}

func NewCreatePriorityUseCase() *CreatePriorityUseCase {
	return &CreatePriorityUseCase{
		PriorityRepository: repository.NewPriorityRepository(),
	}
}

type CreatePriorityInputDTO struct {
	Level string
}

func (uc *CreatePriorityUseCase) Execute(ctx echo.Context, dto CreatePriorityInputDTO) error {
	priority, err := entity.NewPriority(dto.Level)
	if err != nil {
		return err
	}

	if err := uc.PriorityRepository.Create(ctx, priority); err != nil {
		return err
	}

	return nil
}
