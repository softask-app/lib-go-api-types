package apitype

import (
	"time"

	"github.com/francoispqt/gojay"
)

// TaskDetails represents a "full" task object.
type TaskDetails struct {
	TaskMeta

	Description string
	Creator     UserMeta
	Assignees   UserList
	Steps       StepList
	Created     time.Time
	Updated     time.Time
	Deleted     *time.Time
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (s *TaskDetails) MarshalJSONObject(enc *gojay.Encoder) {
	s.TaskMeta.MarshalJSONObject(enc)

	enc.AddStringKey(JsKeyDescription, s.Description)
	enc.AddObjectKey(JsKeyCreator, &s.Creator)
	enc.AddArrayKeyOmitEmpty(JsKeyAssignees, s.Assignees)
	enc.AddArrayKeyOmitEmpty(JsKeySteps, s.Steps)
	enc.AddTimeKey(JsKeyCreated, &s.Created, time.RFC3339Nano)
	enc.AddTimeKey(JsKeyUpdated, &s.Updated, time.RFC3339Nano)

	if s.Deleted != nil {
		enc.AddTimeKey(JsKeyDeleted, s.Deleted, time.RFC3339Nano)
	}
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (s *TaskDetails) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyDescription:
		return d.String(&s.Description)
	case JsKeyCreator:
		return d.Object(&s.Creator)
	case JsKeyAssignees:
		return d.Array(&s.Assignees)
	case JsKeySteps:
		return d.Array(&s.Steps)
	case JsKeyCreated:
		return d.Time(&s.Created, time.RFC3339Nano)
	case JsKeyUpdated:
		return d.Time(&s.Updated, time.RFC3339Nano)
	case JsKeyDeleted:
		return d.Time(s.Deleted, time.RFC3339Nano)
	}

	return s.TaskMeta.UnmarshalJSONObject(d, k)
}

func (s *TaskDetails) NKeys() int {
	return 7 + s.TaskMeta.NKeys()
}

func (s TaskDetails) IsNil() bool {
	return false
}
