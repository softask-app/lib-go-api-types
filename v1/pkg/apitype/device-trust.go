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
	e.StringKey(JsKeyToken, t.Token.String())
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (t *DeviceTrustResponse) UnmarshalJSONObject(d *gojay.Decoder, s string) error {
	if s != JsKeyToken {
		return errBadKey(s)
	}

	var tmp string
	if err := d.String(&tmp); err != nil {
		return err
	}

	return t.Token.FromString(tmp)
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (t *DeviceTrustResponse) NKeys() int {
	return 1
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (t DeviceTrustResponse) IsNil() bool {
	return false
}
