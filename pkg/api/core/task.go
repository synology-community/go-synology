package core

type TaskExtra struct {
	NotifyEnable  bool   `json:"notify_enable"`
	Script        string `json:"script,omitempty"`
	NotifyMail    string `json:"notify_mail"`
	NotifyIfError bool   `json:"notify_if_error"`
}

type TaskSchedule struct {
	DateType              int64    `json:"date_type"`
	Date                  string   `json:"date,omitempty"`
	WeekDay               string   `json:"week_day"`
	Hour                  int64    `json:"hour"`
	Minute                int64    `json:"minute"`
	MonthlyWeek           []string `json:"monthly_week"`
	RepeatHour            int64    `json:"repeat_hour"`
	RepeatMin             int64    `json:"repeat_min"`
	RepeatDate            int64    `json:"repeat_date"`
	RepeatMinStoreConfig  []int64  `json:"repeat_min_store_config"`
	RepeatHourStoreConfig []int64  `json:"repeat_hour_store_config"`
	LastWorkHour          *int64   `json:"last_work_hour"`
}

type TaskRequest struct {
	Name               string       `url:"name" json:"name,omitempty"`
	RealOwner          string       `url:"real_owner" json:"real_owner,omitempty"`
	Owner              string       `url:"owner" json:"owner,omitempty"`
	Schedule           TaskSchedule `url:"schedule,json,omitempty" json:"schedule,omitempty"`
	Extra              TaskExtra    `url:"extra,json,omitempty" json:"extra,omitempty"`
	Type               string       `url:"type,omitempty" json:"type,omitempty"`
	Enable             bool         `url:"enable,omitempty" json:"enable,omitempty"`
	ID                 *int64       `url:"id,omitempty" json:"id,omitempty"`
	SynoConfirmPWToken string       `url:"SynoConfirmPWToken,omitempty" json:"SynoConfirmPWToken,omitempty"`
}

type ListTaskRequest struct {
	SortBy string `form:"sort_by,omitempty" url:"sort_by,omitempty"`
}

type TaskResult struct {
	ID              *int64 `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Action          string `json:"action,omitempty"`
	CanDelete       bool   `json:"can_delete,omitempty"`
	CanEdit         bool   `json:"can_edit,omitempty"`
	CanRun          bool   `json:"can_run,omitempty"`
	Enable          bool   `json:"enable,omitempty"`
	NextTriggerTime string `json:"next_trigger_time,omitempty"`
	Owner           string `json:"owner,omitempty"`
	RealOwner       string `json:"real_owner,omitempty"`
	Type            string `json:"type,omitempty"`
}

type ListTaskResponse struct {
	Tasks  []TaskResult `json:"tasks,omitempty"`
	Total  int64        `json:"total,omitempty"`
	Offset int64        `json:"offset,omitempty"`
}

type TaskGetRequest struct {
	ID int64 `url:"id"`
}

type TaskRef struct {
	ID        int64  `json:"id,omitempty"`
	RealOwner string `json:"name,omitempty"`
}

type TaskDeleteRequest struct {
	Tasks []TaskRef `url:"tasks,json"`
}
type TaskRunRequest struct {
	Tasks []TaskRef `url:"tasks,json"`
}
