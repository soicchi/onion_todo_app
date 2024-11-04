package entity

type TodoDetail struct {
	todo     *Todo
	priority *Priority
	status   *Status
}

func ReconstructTodoDetail(todo *Todo, priority *Priority, status *Status) *TodoDetail {
	return &TodoDetail{
		todo:     todo,
		priority: priority,
		status:   status,
	}
}

func (td *TodoDetail) Todo() *Todo {
	return td.todo
}

func (td *TodoDetail) Priority() *Priority {
	return td.priority
}

func (td *TodoDetail) Status() *Status {
	return td.status
}
