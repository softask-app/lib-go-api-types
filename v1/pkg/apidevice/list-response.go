package apidevice

import (
	"github.com/francoispqt/gojay"
)

type ListResponse []DetailsResponse

func (d ListResponse) MarshalJSONArray(e *gojay.Encoder) {
	for i := range d {
		e.AddObject(&d[i])
	}
}

func (d ListResponse) IsNil() bool {
	return false
}

