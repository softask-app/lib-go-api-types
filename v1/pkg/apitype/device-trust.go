package apitype

import (
	"errors"
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

type DeviceTrustResponse struct {
	Token apitoken.Token256
}

func (t *DeviceTrustResponse) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	if s != JsKeyToken {
		return errBadKey(s)
	}

	return d.Array(t.Token)
}

func (t *DeviceTrustResponse) NKeys() int {
	return 1
}

func (t *DeviceTrustResponse) MarshalJSONObject(e *gojay.Encoder) {
	e.AddArrayKey(JsKeyToken, t.Token)
}

func (t DeviceTrustResponse) IsNil() bool {
	return false
}
