package api

import "fmt"

type Credentials struct {
	SessionID string `url:"_sid"`
	DeviceID  string `url:"device_id,omitempty"`
	Token     string `url:"SynoToken,omitempty"`
}

func (c Credentials) GetCookie() string {
	return fmt.Sprintf("did=%s; id=%s", c.DeviceID, c.SessionID)
}
