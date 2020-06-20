package apitype

import (
	"github.com/francoispqt/gojay"
	"time"
)

type UserDetail struct {
	UserMeta

	Email     string
	Providers UserList
	Supports  UserList
	Created   time.Time
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (u UserDetail) MarshalJSONObject(e *gojay.Encoder) {
	u.UserMeta.MarshalJSONObject(e)
	e.StringKey(JsKeyEmail, u.Email)
	e.ArrayKeyOmitEmpty(JsKeyProviders, u.Providers)
	e.ArrayKeyOmitEmpty(JsKeySupports, u.Supports)
	e.TimeKey(JsKeyCreated, &u.Created, time.RFC3339Nano)
}

func (u *UserDetail) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	switch s {
	case JsKeyEmail:
		return d.String(&u.Email)
	case JsKeyCreated:
		return d.Time(&u.Created, time.RFC3339Nano)
	case JsKeyProviders:
		return d.Array(&u.Providers)
	case JsKeySupports:
		return d.Array(&u.Supports)
	}

	return u.UserMeta.UnmarshalJSONObject(d, s)
}

func (u *UserDetail) NKeys() int {
	return 4 + u.UserMeta.NKeys()
}

func (u UserDetail) IsNil() bool {
	return false
}
