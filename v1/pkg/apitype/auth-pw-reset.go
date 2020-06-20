package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

// PasswordResetRequest defines the body of a request to the HTTP service to
// update a user's password to the given value.
type PasswordResetRequest struct {

	// Database ID of the user record to update.
	UserId uint64

	// Password reset token (will have been sent to the user's email)
	Token apitoken.Token256

	// New password value
	Password string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (p *PasswordResetRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Uint64Key(JsKeyUserId, p.UserId)
	enc.ArrayKey(JsKeyToken, p.Token)
	enc.StringKey(JsKeyPassword, p.Password)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (p *PasswordResetRequest) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	switch s {
	case JsKeyUserId:
		return d.Uint64(&p.UserId)
	case JsKeyToken:
		return p.Token.UnmarshalJSONArray(d)
	case JsKeyPassword:
		return d.String(&p.Password)
	}

	return nil
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (p *PasswordResetRequest) IsNil() bool {
	return false
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (p *PasswordResetRequest) NKeys() int {
	return 3
}

// RequestPasswordResetRequest defines the body of a request to the HTTP server
// to send a password reset to the given email (provided that email is attached
// to a known user).
type RequestPasswordResetRequest struct {
	Email string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (r *RequestPasswordResetRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(JsKeyEmail, r.Email)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (r *RequestPasswordResetRequest) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	if s == JsKeyEmail {
		return d.String(&r.Email)
	}

	return nil
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (r *RequestPasswordResetRequest) NKeys() int {
	return 1
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (r *RequestPasswordResetRequest) IsNil() bool {
	return false
}
