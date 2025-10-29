package core

type PasswordConfirmRequest struct {
	Password string `url:"password,quoted"`
}

type PasswordConfirmResponse struct {
	SynoConfirmPWToken string `json:"SynoConfirmPWToken,omitempty"`
}

type User struct {
	ID          string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email,omitempty"`
	Expire      string `json:"expire,omitempty"`
}

type UserListRequest struct {
	Additional []string `url:"additional,omitempty"`
}

type UserListResponse struct {
	Users  []User `json:"users,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Total  int    `json:"total,omitempty"`
}

// UserCreateRequest for creating a user.
type UserCreateRequest struct {
	Name              string   `url:"name"`
	Password          string   `url:"password"`
	Description       string   `url:"description,omitempty"`
	Email             string   `url:"email,omitempty"`
	Expire            string   `url:"expire,omitempty"`
	CannotChangePass  bool     `url:"cannot_chg_passwd,omitempty"`
	PasswdNeverExpire bool     `url:"passwd_never_expire,omitempty"`
	NotifyByEmail     bool     `url:"notify_by_email,omitempty"`
	SendPassword      bool     `url:"send_password,omitempty"`
	Groups            []string `url:"groups,omitempty"`
}

// UserCreateResponse for creating a user.
type UserCreateResponse struct {
	User User `json:"user,omitempty"`
}

// UserDeleteRequest for deleting a user.
type UserDeleteRequest struct {
	Name string `url:"name"`
}

// UserDeleteResponse for deleting a user.
type UserDeleteResponse struct {
	Success bool `json:"success"`
}

// UserModifyRequest for modifying a user.
type UserModifyRequest struct {
	Name              string   `url:"name"`
	NewName           string   `url:"new_name,omitempty"`
	Password          string   `url:"password,omitempty"`
	Description       string   `url:"description,omitempty"`
	Email             string   `url:"email,omitempty"`
	Expire            string   `url:"expire,omitempty"`
	CannotChangePass  bool     `url:"cannot_chg_passwd,omitempty"`
	PasswdNeverExpire bool     `url:"passwd_never_expire,omitempty"`
	NotifyByEmail     bool     `url:"notify_by_email,omitempty"`
	SendPassword      bool     `url:"send_password,omitempty"`
	Groups            []string `url:"groups,omitempty"`
}

// UserModifyResponse for modifying a user.
type UserModifyResponse struct {
	User User `json:"user,omitempty"`
}
