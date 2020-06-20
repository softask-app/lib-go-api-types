package apitype

import (
	"github.com/francoispqt/gojay"
)

// UserMeta defines a base representation of a user record.
type UserMeta struct {

	// Database ID for the user record.
	Id uint64

	// User display name / nickname.
	DisplayName string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (u UserMeta) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyId, u.Id)
	e.StringKey(JsKeyDisplayName, u.DisplayName)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (u *UserMeta) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	switch s {
	case JsKeyId:
		return d.Uint64(&u.Id)
	case JsKeyDisplayName:
		return d.String(&u.DisplayName)
	}

	return errBadKey(s)
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (u *UserMeta) NKeys() int {
	return 2
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (u UserMeta) IsNil() bool {
	return false
}
