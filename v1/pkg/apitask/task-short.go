package apitask

import "github.com/francoispqt/gojay"

// TaskMeta represents a minimal representation of a task.
type TaskMeta struct {
	Id   uint64
	Name string
}

func (s *TaskMeta) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyTaskId:
		return d.Uint64(&s.Id)
	case JsKeyTaskName:
		return d.String(&s.Name)
	}
	return nil
}

func (s *TaskMeta) NKeys() int {
	return 2
}

func (s *TaskMeta) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddUint64Key(JsKeyTaskId, s.Id)
	enc.AddStringKey(JsKeyTaskName, s.Name)
}

func (s TaskMeta) IsNil() bool {
	return false
}
