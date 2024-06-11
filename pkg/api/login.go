package api

type LoginRequest struct {
	// query.Encoder
	Account         string `url:"account"`
	Password        string `url:"passwd"`
	Session         string `url:"session,omitempty"`
	Format          string `url:"format,omitempty"`
	EnableSynoToken string `url:"enable_syno_token,omitempty"`
}

type LoginResponse struct {
	DeviceID     string `json:"did,omitempty"`
	SessionID    string `json:"sid"`
	Token        string `json:"synotoken"`
	IsPortalPort bool   `json:"is_portal_port,omitempty"`
}

// func (l LoginRequest) EncodeValues(key string, v *url.Values) error {
// 	if key == "enable_syno_token" || key == "EnableSynoToken" {
// 		if l.EnableSynoToken {
// 			v.Add(key, "yes")
// 		} else {
// 			v.Add(key, "no")
// 			(*v).Set(key, "no")
// 		}
// 	}

// 	return nil
// }
