package apitype

import (
	"time"

	"github.com/francoispqt/gojay"
)

type StepDetails struct {
	Id          uint64
	Task        TaskDetails
	Description string
	Position    uint16
	Creator     UserMeta
	Created     time.Time
	Updated     time.Time
	Deleted     *time.Time
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (s StepDetails) MarshalJSONObject(e *gojay.Encoder) {
	e.Uint64Key(JsKeyId, s.Id)
	e.ObjectKey(JsKeyTask, &s.Task)
	e.StringKey(JsKeyDescription, s.Description)
	e.Uint16Key(JsKeyPosition, s.Position)
	e.ObjectKey(JsKeyCreator, &s.Creator)
	e.TimeKey(JsKeyCreated, &s.Created, time.RFC3339Nano)
	e.TimeKey(JsKeyUpdated, &s.Updated, time.RFC3339Nano)

	if s.Deleted != nil {
		e.TimeKey(JsKeyDeleted, s.Deleted, time.RFC3339Nano)
	}
}

func (s *StepDetails) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyId:
		return d.Uint64(&s.Id)
	case JsKeyTask:
		return d.Object(&s.Task)
	case JsKeyDescription:
		return d.String(&s.Description)
	case JsKeyPosition:
		return d.Uint16(&s.Position)
	case JsKeyCreator:
		return d.Object(&s.Creator)
	case JsKeyCreated:
		return d.Time(&s.Created, time.RFC3339Nano)
	case JsKeyUpdated:
		return d.Time(&s.Updated, time.RFC3339Nano)
	case JsKeyDeleted:
		return d.Time(s.Deleted, time.RFC3339Nano)
	}

	return errBadKey(k)
}

func (*StepDetails) NKeys() int {
	return 8
}

func (s StepDetails) IsNil() bool {
	return false
}
