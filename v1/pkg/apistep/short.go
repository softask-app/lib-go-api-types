package apistep

import (
	"errors"
	"github.com/francoispqt/gojay"
)

const (
	JsKeyStepId          = "id"
	JsKeyStepTaskId      = "taskId"
	JsKeyStepDescription = "description"
	JsKeyStepPosition    = "position"
)

type StepMeta struct {
	Id          uint64
	TaskId      uint64
	Description string
	Position    uint16
}

func (s *StepMeta) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyStepId:
		return d.Uint64(&s.Id)
	case JsKeyStepTaskId:
		return d.Uint64(&s.TaskId)
	case JsKeyStepDescription:
		return d.String(&s.Description)
	case JsKeyStepPosition:
		return d.Uint16(&s.Position)
	}

	return errors.New("unrecognized json key " + k)
}

func (s *StepMeta) NKeys() int {
	return 4
}

func (s StepMeta) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyStepId, s.Id)
	e.Uint64Key(JsKeyStepTaskId, s.TaskId)
	e.StringKey(JsKeyStepDescription, s.Description)
	e.Uint16Key(JsKeyStepPosition, s.Position)
}

func (s StepMeta) IsNil() bool {
	return false
}
