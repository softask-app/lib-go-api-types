package apitask

import (
	"errors"
	"time"

	"github.com/francoispqt/gojay"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apiuser"
)

const (
	JsKeyTaskDescription = "description"
	JsKeyTaskCreator     = "creator"
	JsKeyTaskAssignees   = "assignees"
	JsKeyTaskSteps       = "steps"
	JsKeyTaskCreated     = "created"
	JsKeyTaskUpdated     = "updated"
)

// TaskDetails represents a "full" task object.
type TaskDetails struct {
	TaskMeta

	Description string
	Creator     apiuser.UserMeta
	Assignees   apiuser.UserList
	Steps       StepList
	Created     time.Time
	Updated     time.Time
}

func (s *TaskDetails) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyTaskId:
		return d.Uint64(&s.Id)
	case JsKeyTaskName:
		return d.String(&s.Name)
	case JsKeyTaskDescription:
		return d.String(&s.Description)
	case JsKeyTaskCreator:
		return d.Object(&s.Creator)
	case JsKeyTaskAssignees:
		return d.Array(&s.Assignees)
	case JsKeyTaskSteps:
		return d.Array(&s.Steps)
	case JsKeyTaskCreated:
		return d.Time(&s.Created, time.RFC3339Nano)
	case JsKeyTaskUpdated:
		return d.Time(&s.Updated, time.RFC3339Nano)
	}
	return errors.New("unrecognized json key " + k)
}

func (s *TaskDetails) NKeys() int {
	return 8
}

func (s *TaskDetails) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddUint64Key(JsKeyTaskId, s.Id)
	enc.AddStringKey(JsKeyTaskName, s.Name)
	enc.AddStringKey(JsKeyTaskDescription, s.Description)
	enc.AddObjectKey(JsKeyTaskCreator, &s.Creator)
	enc.AddArrayKeyOmitEmpty(JsKeyTaskAssignees, s.Assignees)
	enc.AddArrayKeyOmitEmpty(JsKeyTaskSteps, s.Steps)
	enc.AddTimeKey(JsKeyTaskCreated, &s.Created, time.RFC3339Nano)
	enc.AddTimeKey(JsKeyTaskUpdated, &s.Updated, time.RFC3339Nano)
}

func (s TaskDetails) IsNil() bool {
	return false
}
