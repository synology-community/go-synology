package core

type EventOwner = map[string]string

type EventRequest struct {
	Name               string     `url:"task_name,quoted"             json:"task_name,omitempty"`
	Owner              EventOwner `url:"owner,json"                   json:"owner,omitempty"`
	Event              string     `url:"event,quoted"                 json:"event,omitempty"`
	DependOnTask       string     `url:"depend_on_task,quoted"        json:"depend_on_task,omitempty"`
	NotifyEnabled      bool       `url:"notify_enable"                json:"notify_enable,omitempty"`
	NotifyMail         string     `url:"notify_mail,quoted"           json:"notify_mail,omitempty"`
	NotifyIfError      bool       `url:"notify_if_error"              json:"notify_if_error,omitempty"`
	OperationType      string     `url:"operation_type,quoted"        json:"operation_type,omitempty"`
	Operation          string     `url:"operation,quoted"             json:"operation,omitempty"`
	Enable             bool       `url:"enable"                       json:"enable,omitempty"`
	SynoConfirmPWToken string     `url:"SynoConfirmPWToken,omitempty" json:"SynoConfirmPWToken,omitempty"`
}

type EventResult struct{}
