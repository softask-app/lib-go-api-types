package apitype

import (
	"github.com/francoispqt/gojay"
)

// LoginRequest defines the body of an authentication request to the HTTP
// service using user credentials.
type LoginRequest struct {
	Email    string
	Password string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (l *LoginRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(JsKeyEmail, l.Email)
	enc.StringKey(JsKeyPassword, l.Password)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
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

// IsNil implements the gojay MarshalerJSONObject interface.
func (l *LoginRequest) IsNil() bool {
	return false
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (l *LoginRequest) NKeys() int {
	return 2
}
