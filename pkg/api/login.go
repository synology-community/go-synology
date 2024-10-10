package api

type LoginOptions struct {
	Username   string
	Password   string
	OTPSecret  string
	DeviceID   string
	DeviceName string
}

type LoginRequest struct {
	// query.Encoder
	Account           string `url:"account"`
	Password          string `url:"passwd"`
	OTPCode           string `url:"otp_code,omitempty"`
	Session           string `url:"session,omitempty"`
	Format            string `url:"format,omitempty"`
	EnableSynoToken   string `url:"enable_syno_token,omitempty"`
	EnableDeviceToken string `url:"enable_device_token,omitempty"`
	DeviceID          string `url:"device_id,omitempty"`
	DeviceName        string `url:"device_name,omitempty"`
	LoginType         string `url:"logintype,omitempty"`
	TimeZone          string `url:"timezone,omitempty"`
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
