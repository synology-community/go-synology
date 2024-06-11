package api

type Credentials struct {
	SessionID string `url:"_sid"`
	Token     string `url:"SynoToken"`
}
