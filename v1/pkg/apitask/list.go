package apitask

import "github.com/francoispqt/gojay"

const (
	JsKeyTaskId   = "id"
	JsKeyTaskName = "name"
)

type TaskList []TaskMeta

func (l TaskList) MarshalJSONArray(e *gojay.Encoder) {
	for i := range l {
		e.AddObject(&l[i])
	}
}

func (l TaskList) IsNil() bool {
	return false
}

func (l *TaskList) UnmarshalJSONArray(d *gojay.Decoder) error {
	var tmp TaskMeta
	if err := d.DecodeObject(&tmp); err != nil {
		return err
	}

	*l = append(*l, tmp)
	return nil
}
