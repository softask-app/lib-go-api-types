package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-password/v1/pkg/passwd"
)

// LoginRequest defines the body of an authentication request to the HTTP
// service using user credentials.
type LoginRequest struct {
	Email    string
	Password passwd.Password
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (l *LoginRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(JsKeyEmail, l.Email)
	enc.StringKey(JsKeyPassword, string(l.Password))
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (l *LoginRequest) UnmarshalJSONObject(dec *gojay.Decoder, s string) error {
	if s == JsKeyEmail {
		if err := dec.String(&l.Email); err != nil {
			return err
		}
	} else if s == JsKeyPassword {
		var tmp string

		if err := dec.String(&tmp); err != nil {
			return err
		}

		l.Password = passwd.Password(tmp)
	}

	return nil
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (l *LoginRequest) IsNil() bool {
	return false
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (l *LoginRequest) NKeys() int {
	return 2
}
