package apitype

import (
	"github.com/francoispqt/gojay"
)

type TaskCreateRequest struct {
	Name        string
	Description string
}

func (t *TaskCreateRequest) MarshalJSONObject(e *gojay.Encoder) {
	e.StringKey(JsKeyName, t.Name)
	e.StringKeyOmitEmpty(JsKeyDescription, t.Description)
}

func (t *TaskCreateRequest) UnmarshalJSONObject(d *gojay.Decoder, k string) error {
	switch k {
	case JsKeyName:
		return d.String(&t.Name)
	case JsKeyDescription:
		return d.String(&t.Description)
	}

	return errBadKey(k)
}

func (t *TaskCreateRequest) IsNil() bool {
	return false
}

func (t *TaskCreateRequest) NKeys() int {
	return 2
}
