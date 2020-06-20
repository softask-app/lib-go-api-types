package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

type DeviceTrustResponse struct {
	Token apitoken.Token256
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (t *DeviceTrustResponse) MarshalJSONObject(e *gojay.Encoder) {
	e.AddArrayKey(JsKeyToken, t.Token)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (t *DeviceTrustResponse) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	if s != JsKeyToken {
		return errBadKey(s)
	}

	return d.Array(t.Token)
}

func (t *DeviceTrustResponse) NKeys() int {
	return 1
}

func (t DeviceTrustResponse) IsNil() bool {
	return false
}
