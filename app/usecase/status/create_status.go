package status

import (
	"fmt"

	"onion_todo_app/domain/entity"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type CreateStatusUseCase struct {
	statusRepository entity.StatusRepository
}

func NewCreateStatusUseCase() *CreateStatusUseCase {
	return &CreateStatusUseCase{
		statusRepository: repository.NewStatusRepository(),
	}
}

type CreateStatusInputDTO struct {
	State string
}

func (uc *CreateStatusUseCase) Execute(ctx echo.Context, dto CreateStatusInputDTO) error {
	// convert DTO to domain object
	status, err := entity.NewStatus(dto.State)
	if err != nil {
		return fmt.Errorf("error creating new status: %w", err)
	}

	if err := uc.statusRepository.Create(ctx, status); err != nil {
		return fmt.Errorf("error creating status: %w", err)
	}

	return nil
}
