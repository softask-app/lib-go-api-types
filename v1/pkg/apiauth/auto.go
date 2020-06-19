package apiauth

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

const (
	JsKeyUserId   = "userId"
	JsKeyDeviceId = "deviceId"
	JsKeyToken    = "token"
)

// AutoAuth defines the request body for automatic/device authentication
// requests.
type AutoAuth struct {
	UserId   uint64
	DeviceId apitoken.Token128
	Token    apitoken.Token256
}

func (a *AutoAuth) UnmarshalJSONObject(dec *gojay.Decoder, s string) (err error) {
	switch s {
	case JsKeyUserId:
		err = dec.Uint64(&a.UserId)
	case JsKeyDeviceId:
		err = a.DeviceId.UnmarshalJSONArray(dec)
	case JsKeyToken:
		err = a.Token.UnmarshalJSONArray(dec)
	}

	return
}

func (a AutoAuth) NKeys() int {
	return 3
}
