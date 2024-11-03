package status

import (
	"fmt"

	"onion_todo_app/domain/status"
	"onion_todo_app/infrastructure/postgres/repository"

	"github.com/labstack/echo/v4"
)

type CreateStatusUseCase struct {
	statusRepository status.StatusRepository
}

func NewCreateStatusUseCase() *CreateStatusUseCase {
	return &CreateStatusUseCase{
		statusRepository: repository.StatusRepository{},
	}
}

type CreateStatusInputDTO struct {
	State string
}

func (uc *CreateStatusUseCase) Execute(ctx echo.Context, dto CreateStatusInputDTO) error {
	status, err := status.NewStatus(dto.State)
	if err != nil {
		return fmt.Errorf("error creating new status: %w", err)
	}

	if err := uc.statusRepository.Create(ctx, status); err != nil {
		return fmt.Errorf("error creating status: %w", err)
	}

	return nil
}
