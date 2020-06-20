package apitype

import (
	"time"

	"github.com/francoispqt/gojay"
)

// UserDetails defines the info available for a specific user record.
type UserDetails struct {
	UserMeta

	Email     string
	Providers UserList
	Supports  UserList
	Created   time.Time
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (u UserDetails) MarshalJSONObject(e *gojay.Encoder) {
	u.UserMeta.MarshalJSONObject(e)
	e.StringKey(JsKeyEmail, u.Email)
	e.ArrayKeyOmitEmpty(JsKeyProviders, u.Providers)
	e.ArrayKeyOmitEmpty(JsKeySupports, u.Supports)
	e.TimeKey(JsKeyCreated, &u.Created, time.RFC3339Nano)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (u *UserDetails) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
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

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (u *UserDetails) NKeys() int {
	return 4 + u.UserMeta.NKeys()
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (u UserDetails) IsNil() bool {
	return false
}
