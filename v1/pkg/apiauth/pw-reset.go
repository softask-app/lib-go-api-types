package apiauth

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

const (
	JsKeyPwEmail    = "email"
	JsKeyPwUserId   = "userId"
	JsKeyPwToken    = "token"
	JsKeyPwPassword = "password"
)

type PasswordResetRequest struct {
	UserId   uint64
	Token    apitoken.Token256
	Password string
}

func (p *PasswordResetRequest) IsNil() bool {
	panic("implement me")
}

func (p *PasswordResetRequest) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	switch s {
	case JsKeyUserId:
		return d.Uint64(&p.UserId)
	case JsKeyToken:
		return p.Token.UnmarshalJSONArray(d)
	case JsKeyPwPassword:
		return d.String(&p.Password)
	}

	return nil
}

func (p *PasswordResetRequest) NKeys() int {
	return 3
}

type RequestPasswordRequest struct {
	Email string
}

func (r *RequestPasswordRequest) UnmarshalJSONObject(d *gojay.Decoder, s string) (err error) {
	if s == JsKeyEmail {
		err = d.String(&r.Email)
	}

	return nil
}

func (r *RequestPasswordRequest) NKeys() int {
	return 1
}
