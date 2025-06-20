package api

import "fmt"

type Credentials struct {
	SessionID string `url:"_sid"                json:"sid"`
	DeviceID  string `url:"-"                   json:"did"`
	Token     string `url:"SynoToken,omitempty" json:"token"`
}

func (c Credentials) GetCookie() string {
	return fmt.Sprintf("did=%s; id=%s", c.DeviceID, c.SessionID)
}
