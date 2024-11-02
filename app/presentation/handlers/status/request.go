package status

type CreateStatusRequest struct {
	State string `json:"state" validate:"required"`
}
