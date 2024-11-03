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

func (td *TodoDetail) ID() uuid.UUID {
	return td.id
}

func (td *TodoDetail) PriorityID() uuid.UUID {
	return td.priorityID
}

func (td *TodoDetail) StatusID() uuid.UUID {
	return td.statusID
}

func (td *TodoDetail) Title() string {
	return td.title
}

func (td *TodoDetail) Description() string {
	return td.description
}

func (td *TodoDetail) Priority() string {
	return td.priority
}

func (td *TodoDetail) Status() string {
	return td.status
}
