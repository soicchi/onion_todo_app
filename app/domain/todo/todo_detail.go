package todo

import (
	"github.com/google/uuid"
)

type TodoDetail struct {
	id          uuid.UUID
	priorityID  uuid.UUID
	statusID    uuid.UUID
	title       string
	description string
	priority    string
	status      string
}

func ReconstructTodoDetail(id, priorityID, statusID uuid.UUID, title, description, priority, status string) *TodoDetail {
	return &TodoDetail{
		id:          id,
		priorityID:  priorityID,
		statusID:    statusID,
		title:       title,
		description: description,
		priority:    priority,
		status:      status,
	}
}
