package apitype

import (
	"time"

	"github.com/francoispqt/gojay"
)

// TaskDetails represents a "full" task object.
type TaskDetails struct {
	TaskMeta

	// User provided description of the task.
	Description string

	// User metadata for the user that created the task.
	Creator UserMeta

	// User metadata for the users to whom this task is assigned.
	//
	// This value will only be populated with users that the current user is
	// marked as a provider for.
	//
	// If the current user is not a provider, the task has no assignees, or the
	// task is self assigned, this field will be omitted by the server.
	Assignees UserList

	// An optional list of steps that make up the current task.
	Steps StepList

	// Timestamp of when this task was created.
	Created time.Time

	// Timestamp of when this task was last updated.
	Updated time.Time

	// Optional timestamp of when this task was deleted (only present if the task
	// actually has been deleted).
	Deleted *time.Time
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

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (s *TaskDetails) NKeys() int {
	return 7 + s.TaskMeta.NKeys()
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (s TaskDetails) IsNil() bool {
	return false
}
