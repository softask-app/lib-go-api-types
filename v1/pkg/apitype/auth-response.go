package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

const (
	JsApiKey = "apiKey"
)

// AuthResponse defines the response body for a successful authentication
// request.
type AuthResponse struct {
	ApiKey apitoken.Token128
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (a AuthResponse) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey(JsApiKey, a.ApiKey.String())
}

func (a *AuthResponse) UnmarshalJSONObject(dec *gojay.Decoder, s string) error {
	if s == JsApiKey {
		return a.ApiKey.UnmarshalJSONArray(dec)
	}

	return nil
}

func (a AuthResponse) IsNil() bool {
	return false
}

func (a *AuthResponse) NKeys() int {
	return 1
}
