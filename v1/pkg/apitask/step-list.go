package apitask

import (
	"github.com/francoispqt/gojay"
)

type StepList []StepMeta

func (s *StepList) UnmarshalJSONArray(d *gojay.Decoder) error {
	var tmp StepMeta
	if err := d.Object(&tmp); err != nil {
		return err
	}

	*s = append(*s, tmp)

	return nil
}

func (s StepList) MarshalJSONArray(e *gojay.Encoder) {
	for i := range s {
		e.AddObject(&s[i])
	}
}

func (s StepList) IsNil() bool {
	return false
}


