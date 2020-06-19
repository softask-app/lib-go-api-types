package apiuser

import (
	"errors"
	"github.com/francoispqt/gojay"
)

const (
	JsKeyUserId          = "id"
	JsKeyUserDisplayName = "displayName"
)

type UserMeta struct {
	Id          uint64
	DisplayName string
}

func (u *UserMeta) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	switch s {
	case JsKeyUserId:
		return d.Uint64(&u.Id)
	case JsKeyUserDisplayName:
		return d.String(&u.DisplayName)
	}

	return errors.New("unrecognized json key " + s)
}

func (u *UserMeta) NKeys() int {
	return 2
}

func (u UserMeta) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyUserId, u.Id)
	e.StringKey(JsKeyUserDisplayName, u.DisplayName)
}

func (u UserMeta) IsNil() bool {
	return false
}

