package core

type EventOwner = map[string]string

type EventRequest struct {
	Name               string     `url:"task_name" json:"task_name,omitempty"`
	Owner              EventOwner `url:"owner" json:"owner,omitempty"`
	Event              string     `url:"event" json:"event,omitempty"`
	DependOnTask       string     `url:"depend_on_task" json:"depend_on_task,omitempty"`
	NotifyEnabled      bool       `url:"notify_enable" json:"notify_enable,omitempty"`
	NotifyMail         string     `url:"notify_mail" json:"notify_mail,omitempty"`
	NotifyIfError      bool       `url:"notify_if_error" json:"notify_if_error,omitempty"`
	OperationType      string     `url:"operation_type" json:"operation_type,omitempty"`
	Operation          string     `url:"operation" json:"operation,omitempty"`
	Enable             bool       `url:"enable,omitempty" json:"enable,omitempty"`
	SynoConfirmPWToken string     `url:"SynoConfirmPWToken,omitempty" json:"SynoConfirmPWToken,omitempty"`
}

type EventResult struct {
}
