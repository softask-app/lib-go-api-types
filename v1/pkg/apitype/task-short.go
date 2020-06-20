package apitype

import "github.com/francoispqt/gojay"

// TaskMeta represents a minimal representation of a task.
type TaskMeta struct {

	// Database ID for the task record.
	Id uint64

	// User provided name for the task.
	Name string
}

// MarshalJSONObject implements the gojay MarshalerJSONObject interface.
func (s *TaskMeta) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddUint64Key(JsKeyTaskId, s.Id)
	enc.AddStringKey(JsKeyName, s.Name)
}

// UnmarshalJSONObject implements the gojay UnmarshalerJSONObject interface.
func (s *TaskMeta) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyTaskId:
		return d.Uint64(&s.Id)
	case JsKeyName:
		return d.String(&s.Name)
	}

	return errBadKey(k)
}

// NKeys implements the gojay UnmarshalerJSONObject interface.
func (s *TaskMeta) NKeys() int {
	return 2
}

// IsNil implements the gojay MarshalerJSONObject interface.
func (s TaskMeta) IsNil() bool {
	return false
}
