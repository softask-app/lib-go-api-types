package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

// AuthResponse defines the response body for a successful authentication
// request.
type AuthResponse struct {
	ApiKey apitoken.Token128
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (a AuthResponse) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(JsKeyApiKey, a.ApiKey.String())
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (a *AuthResponse) UnmarshalJSONObject(dec *gojay.Decoder, s string) error {
	if s != JsKeyApiKey {
		return errBadKey(s)
	}

	var tmp string
	if err := dec.String(&tmp); err != nil {
		return err
	}

	return a.ApiKey.FromString(tmp)
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (a AuthResponse) IsNil() bool {
	return false
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (a *AuthResponse) NKeys() int {
	return 1
}
