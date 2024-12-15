package core

type PasswordConfirmRequest struct {
	Password string `url:"password,quoted"`
}

type PasswordConfirmResponse struct {
	SynoConfirmPWToken string `json:"SynoConfirmPWToken,omitempty"`
}

type User struct {
	ID   string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserListRequest struct {
	Additional []string `url:"additional,omitempty"`
}

type UserListResponse struct {
	Users  []User `json:"users,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Total  int    `json:"total,omitempty"`
}
