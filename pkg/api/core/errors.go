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

type EventNotFoundError struct {
	*ListTaskResponse
}

func (e EventNotFoundError) Error() string {
	s, err := json.Marshal(&e.ListTaskResponse)
	if err == nil {
		return "Event not found: " + string(s)
	}
	return "Event not found: " + err.Error()
}
