package apitype

import (
	"github.com/francoispqt/gojay"
)

type DeviceList []DeviceDetails

func (d *DeviceList) UnmarshalJSONArray(e *gojay.Decoder) error {
	var tmp DeviceDetails
	if err := e.Object(&tmp); err != nil {
		return err
	}

	*d = append(*d, tmp)

	return nil
}

func (d DeviceList) MarshalJSONArray(e *gojay.Encoder) {
	for i := range d {
		e.AddObject(&d[i])
	}
}

func (d DeviceList) IsNil() bool {
	return false
}
