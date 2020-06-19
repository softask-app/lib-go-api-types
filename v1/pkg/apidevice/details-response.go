package apidevice

import (
	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
	"time"
)

const (
	JsKeyDetailsId          = "id"
	JsKeyDetailsDeviceId    = "deviceId"
	JsKeyDetailsDeviceName  = "deviceName"
	JsKeyDetailsDisplayName = "displayName"
	JsKeyDetailsLastLogin   = "lastLogin"
	JsKeyDetailsTrusted     = "trusted"
	JsKeyTrustedSince       = "since"
)

type DetailsResponse struct {
	Id          uint64
	DeviceId    apitoken.Token128
	DeviceName  string
	DisplayName string
	LastLogin   time.Time
	Trusted     *DeviceTrust
}

func (d *DetailsResponse) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyDetailsId, d.Id)
	e.StringKey(JsKeyDetailsDeviceId, d.DeviceId.String())
	e.StringKey(JsKeyDetailsDeviceName, d.DeviceName)
	if len(d.DisplayName) > 0 {
		e.StringKey(JsKeyDetailsDisplayName, d.DisplayName)
	}
	e.TimeKey(JsKeyDetailsLastLogin, &d.LastLogin, time.RFC3339Nano)
	if d.Trusted != nil {
		e.ObjectKey(JsKeyDetailsTrusted, d.Trusted)
	}
}

func (d *DetailsResponse) IsNil() bool {
	return false
}

type DeviceTrust struct {
	Since time.Time
}

func (d *DeviceTrust) MarshalJSONObject(e *gojay.Encoder) {
	e.TimeKey(JsKeyTrustedSince, &d.Since, time.RFC3339Nano)
}

func (d *DeviceTrust) IsNil() bool {
	return false
}

