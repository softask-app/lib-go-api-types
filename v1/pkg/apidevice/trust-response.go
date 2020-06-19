package apidevice

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

const (
	JsKeyTrustToken = "token"
)

type TrustResponse struct {
	Token apitoken.Token256
}

func (t *TrustResponse) MarshalJSONObject(e *gojay.Encoder) {
	e.AddArrayKey(JsKeyTrustToken, t.Token)
}

func (t TrustResponse) IsNil() bool {
	return false
}
