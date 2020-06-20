package apitype

import (
	"github.com/francoispqt/gojay"
)

type StepMeta struct {
	Id          uint64
	TaskId      uint64
	Description string
	Position    uint16
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (s StepMeta) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyId, s.Id)
	e.Uint64Key(JsKeyTaskId, s.TaskId)
	e.StringKey(JsKeyDescription, s.Description)
	e.Uint16Key(JsKeyPosition, s.Position)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (s *StepMeta) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyId:
		return d.Uint64(&s.Id)
	case JsKeyTaskId:
		return d.Uint64(&s.TaskId)
	case JsKeyDescription:
		return d.String(&s.Description)
	case JsKeyPosition:
		return d.Uint16(&s.Position)
	}

	return errBadKey(k)
}

func (s *StepMeta) NKeys() int {
	return 4
}

func (s StepMeta) IsNil() bool {
	return false
}
