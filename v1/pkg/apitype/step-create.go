package apitype

import (
	"github.com/francoispqt/gojay"
)

type StepCreateRequest struct {
	Description string
	Position    uint16
}

func (s *StepCreateRequest) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyDescription:
		return d.String(&s.Description)
	case JsKeyPosition:
		return d.Uint16(&s.Position)
	}

	return errBadKey(k)
}

func (s *StepCreateRequest) NKeys() int {
	return 2
}

func (s *StepCreateRequest) MarshalJSONObject(e *gojay.Encoder) {
	e.StringKey(JsKeyDescription, s.Description)
	e.Uint16Key(JsKeyPosition, s.Position)
}

func (s *StepCreateRequest) IsNil() bool {
	return false
}
