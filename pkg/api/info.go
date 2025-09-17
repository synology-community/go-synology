package api

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed api_info.json
var b []byte

var ApiInfoData ApiInfo = unmarshalApiInfoData()

func unmarshalApiInfoData() ApiInfo {
	var data ApiInfo
	err := json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type (
	ApiInfo  = map[string]InfoData
	InfoData struct {
		Path          string `json:"path"`
		MinVersion    int    `json:"minVersion"`
		MaxVersion    int    `json:"maxVersion"`
		RequestFormat string `json:"requestFormat"`
	}
	UserInfoReq struct {
		Method string `json:"method" query:"method"`
	}
	UserInfo struct {
		OtpEnabled             bool   `json:"OTP_enable"`
		OtpEnforced            bool   `json:"OTP_enforced"`
		DisallowPassowrdChange bool   `json:"disallowchpasswd"`
		Editable               bool   `json:"editable"`
		Email                  string `json:"email"`
		FullName               string `json:"fullname"`
		PasswordLastChange     int    `json:"password_last_change"`
		UserName               string `json:"username"`
	}
)

func (c *Client) GetApiInfo(ctx context.Context) (*map[string]InfoData, error) {
	return List[ApiInfo](c, ctx, Api_Info)
}

// GetUserInfo queries information about the currently authenticated user; it is globally available to ensure API can be properly validated.
func (c *Client) GetUserInfo(ctx context.Context) (*UserInfo, error) {
	uir := UserInfoReq{
		Method: "get",
	}
	return Post[UserInfo](c, ctx, &uir, Core_UserInfo)
}
