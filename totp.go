package goauthultimate

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateOTP(ID int, username string, period int) (string, error) {

	idbyte := []byte{byte(ID)}
	// Pass a valid OTP key and get back a code
	otpStr, err := totp.GenerateCodeCustom(
		base32.StdEncoding.EncodeToString(append(idbyte, secret...)),
		time.Now(),
		totp.ValidateOpts{
			Period:    uint(period),
			Skew:      5,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA512,
		},
	)

	if err != nil {
		return "", err
	}

	return otpStr, nil
}
func ValidateTotp(code string, ID int, username string, period int) (ok bool) {

	idbyte := []byte{byte(ID)}

	ok, err := totp.ValidateCustom(
		code,
		base32.StdEncoding.EncodeToString(append(idbyte, secret...)),
		time.Now(),
		totp.ValidateOpts{
			Period:    uint(period),
			Skew:      5,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA512,
		},
	)
	fmt.Println("validate error: ", err)

	return
}
