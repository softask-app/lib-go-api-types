package apitype

import (
	"github.com/francoispqt/gojay"
)

const (
	JsKeyEmail    = "email"
	JsKeyPassword = "password"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (l *LoginRequest) UnmarshalJSONObject(dec *gojay.Decoder, s string) error {
	if s == JsKeyEmail {
		if err := dec.String(&l.Email); err != nil {
			return err
		}
	} else if s == JsKeyPassword {
		if err := dec.String(&l.Password); err != nil {
			return err
		}
	}

	return nil
}

func (l *LoginRequest) NKeys() int {
	return 2
}
