package apitype

import (
	"time"

	"github.com/francoispqt/gojay"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

// DeviceDetails defines the available information about a user device.
//
// A device may be from either the web app or the phone app.
type DeviceDetails struct {

	// Id is the sequential datastore ID assigned to this device record.
	Id uint64

	// DeviceId is a token generated by a phone app to identify that phone and
	// enable automatic authentication for that device.
	DeviceId apitoken.Token128

	// DeviceName is the reported name of the device/browser.
	DeviceName string

	// DisplayName is an optional, custom, user provided name for the device.
	DisplayName string

	// LastLogin contains the last login time for a device.
	LastLogin time.Time

	// Trusted contains details about the "trust" status for a device.
	Trusted *DeviceTrust
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (d *DeviceDetails) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyId, d.Id)
	e.StringKey(JsKeyDeviceId, d.DeviceId.String())
	e.StringKey(JsKeyDeviceName, d.DeviceName)
	e.StringKeyOmitEmpty(JsKeyDisplayName, d.DisplayName)
	e.TimeKey(JsKeyLastLogin, &d.LastLogin, time.RFC3339Nano)
	if d.Trusted != nil {
		e.ObjectKeyOmitEmpty(JsKeyTrusted, d.Trusted)
	}
}

func (d *DeviceDetails) UnmarshalJSONObject(e *gojay.Decoder, k string) error {
	switch k {
	case JsKeyId:
		return e.Uint64(&d.Id)
	case JsKeyDeviceId:
		var tmp string
		if err := e.String(&tmp); err != nil {
			return err
		}

		return d.DeviceId.FromString(tmp)
	case JsKeyDeviceName:
		return e.String(&d.DeviceName)
	case JsKeyDisplayName:
		return e.String(&d.DisplayName)
	case JsKeyLastLogin:
		return e.Time(&d.LastLogin, time.RFC3339Nano)
	case JsKeyTrusted:
		return e.Object(d.Trusted)
	}

	return errBadKey(k)
}

func (d *DeviceDetails) NKeys() int {
	return 6
}

func (d *DeviceDetails) IsNil() bool {
	return false
}

type DeviceTrust struct {
	Since time.Time
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (d *DeviceTrust) MarshalJSONObject(e *gojay.Encoder) {
	e.TimeKey(JsKeySince, &d.Since, time.RFC3339Nano)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (d *DeviceTrust) UnmarshalJSONObject(e *gojay.Decoder, s string) error {
	if s != JsKeySince {
		return errBadKey(s)
	}

	return e.Time(&d.Since, time.RFC3339Nano)
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (d *DeviceTrust) NKeys() int {
	return 1
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (d *DeviceTrust) IsNil() bool {
	return false
}
