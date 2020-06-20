package apitype

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

// AutoAuth defines the request body for automatic/device authentication
// requests.
type AutoAuth struct {

	// UserId is the database Id for the user record to which the device that is
	// attempting auth belongs.
	UserId uint64

	// DeviceId is a 128bit token generated by the device.
	//
	// This value is not used for the web app as auto-auth is not allowed for that
	// context.
	//
	// For phone apps, this will be generated on first launch per user/per device.
	DeviceId apitoken.Token128

	// Token is a 256bit token provided to the device by the HTTP API.
	Token apitoken.Token256
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (a *AutoAuth) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyUserId, a.UserId)
	e.StringKey(JsKeyDeviceId, a.DeviceId.String())
	e.StringKey(JsKeyToken, a.Token.String())
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (a *AutoAuth) UnmarshalJSONObject(dec *gojay.Decoder, s string) error {
	switch s {
	case JsKeyUserId:
		return dec.Uint64(&a.UserId)
	case JsKeyDeviceId:
		var tmp string
		if err := dec.String(&tmp); err != nil {
			return err
		}

		return a.DeviceId.FromString(tmp)
	case JsKeyToken:
		var tmp string
		if err := dec.String(&tmp); err != nil {
			return err
		}

		return a.Token.FromString(tmp)
	}

	return errBadKey(s)
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (a *AutoAuth) IsNil() bool {
	return false
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (a AutoAuth) NKeys() int {
	return 3
}
