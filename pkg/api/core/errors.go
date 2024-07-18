package core

import "encoding/json"

type TaskNotFoundError struct {
	TaskResult
}

func (e TaskNotFoundError) Error() string {
	s, err := json.Marshal(&e.TaskResult)
	if err == nil {
		return "Task not found: " + string(s)
	}
	return "Task not found: " + err.Error()
}
