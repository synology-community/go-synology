package core

type TaskExtra struct {
	NotifyEnable  bool   `json:"notify_enable,omitempty"`
	Script        string `json:"script,omitempty"`
	NotifyMail    string `json:"notify_mail,omitempty"`
	NotifyIfError bool   `json:"notify_if_error,omitempty"`
}

type TaskSchedule struct {
	DateType              int    `json:"date_type,omitempty"`
	Date                  string `json:"date,omitempty"`
	WeekDay               string `json:"week_day,omitempty"`
	RepeatDate            int    `json:"repeat_date,omitempty"`
	MonthlyWeek           []any  `json:"monthly_week,omitempty"`
	Hour                  int    `json:"hour,omitempty"`
	Minute                int    `json:"minute,omitempty"`
	RepeatHour            int    `json:"repeat_hour,omitempty"`
	RepeatMin             int    `json:"repeat_min,omitempty"`
	LastWorkHour          int    `json:"last_work_hour,omitempty"`
	RepeatMinStoreConfig  []int  `json:"repeat_min_store_config,omitempty"`
	RepeatHourStoreConfig []int  `json:"repeat_hour_store_config,omitempty"`
}

type TaskRequest struct {
	Name               string       `url:"owner,omitempty" json:"name,omitempty"`
	RealOwner          string       `url:"owner,omitempty" json:"real_owner,omitempty"`
	Owner              string       `url:"owner,omitempty" json:"owner,omitempty"`
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
	ID              int    `json:"id,omitempty"`
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
	Total  int          `json:"total,omitempty"`
	Offset int          `json:"offset,omitempty"`
}

type TaskGetRequest struct {
	ID int `url:"id"`
}

type TaskRef struct {
	ID        int    `json:"id,omitempty"`
	RealOwner string `json:"name,omitempty"`
}

type TaskDeleteRequest struct {
	Tasks []TaskRef `url:"tasks,json"`
}
type TaskRunRequest struct {
	Tasks []TaskRef `url:"tasks,json"`
}
