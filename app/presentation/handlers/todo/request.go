package todo

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PriorityID  string `json:"priority_id" validate:"required"`
	StatusID    string `json:"status_id" validate:"required"`
}
