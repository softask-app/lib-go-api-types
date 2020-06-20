package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

type PasswordResetRequest struct {
	UserId   uint64
	Token    apitoken.Token256
	Password string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (p *PasswordResetRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.Uint64Key(JsKeyUserId, p.UserId)
	enc.ArrayKey(JsKeyToken, p.Token)
	enc.StringKey(JsKeyPassword, p.Password)
}

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

func (p *PasswordResetRequest) IsNil() bool {
	return false
}

func (p *PasswordResetRequest) NKeys() int {
	return 3
}

type RequestPasswordRequest struct {
	Email string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (r *RequestPasswordRequest) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(JsKeyEmail, r.Email)
}

func (r *RequestPasswordRequest) UnmarshalJSONObject(d *gojay.Decoder, s string) (err error) {
	if s == JsKeyEmail {
		err = d.String(&r.Email)
	}

	return nil
}

func (r *RequestPasswordRequest) IsNil() bool {
	return false
}

func (r *RequestPasswordRequest) NKeys() int {
	return 1
}
