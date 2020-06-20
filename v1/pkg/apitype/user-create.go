package apitype

import (
	"github.com/francoispqt/gojay"
)

type UserCreateRequest struct {
	Email       string
	DisplayName string
	Password    string
}

func (u *UserCreateRequest) MarshalJSONObject(e *gojay.Encoder) {
	e.StringKey(JsKeyEmail, u.Email)
	e.StringKey(JsKeyDisplayName, u.DisplayName)
	e.StringKey(JsKeyPassword, u.Password)
}

func (u *UserCreateRequest) IsNil() bool {
	return false
}

func (u *UserCreateRequest) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	switch s {
	case JsKeyEmail:
		return d.String(&u.Email)
	case JsKeyDisplayName:
		return d.String(&u.DisplayName)
	case JsKeyPassword:
		return d.String(&u.Password)
	}

	return errBadKey(s)
}

func (u *UserCreateRequest) NKeys() int {
	return 3
}
