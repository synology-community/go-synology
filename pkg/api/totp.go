package api

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func generateTotp(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}
