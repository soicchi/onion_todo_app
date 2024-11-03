package priority

type CreatePriorityRequest struct {
	Level string `json:"level" validate:"required"`
}
