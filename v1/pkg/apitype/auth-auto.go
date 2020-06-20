package apitype

import (
	"errors"
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
	e.ArrayKey(JsKeyDeviceId, a.DeviceId)
	e.ArrayKey(JsKeyToken, a.Token)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (a *AutoAuth) UnmarshalJSONObject(dec *gojay.Decoder, s string) (err error) {
	switch s {
	case JsKeyUserId:
		err = dec.Uint64(&a.UserId)
	case JsKeyDeviceId:
		err = a.DeviceId.UnmarshalJSONArray(dec)
	case JsKeyToken:
		err = a.Token.UnmarshalJSONArray(dec)
	}

	return errors.New("unrecognized ")
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (a *AutoAuth) IsNil() bool {
	return false
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (a AutoAuth) NKeys() int {
	return 3
}
