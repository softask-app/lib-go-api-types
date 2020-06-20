package apitype

import (
	"github.com/francoispqt/gojay"
)

type UserList []UserMeta

func (u *UserList) UnmarshalJSONArray(d *gojay.Decoder) error {
	var tmp UserMeta
	if err := d.Object(&tmp); err != nil {
		return err
	}

	*u = append(*u, tmp)

	return nil
}

func (u UserList) MarshalJSONArray(e *gojay.Encoder) {
	for i := range u {
		e.AddObject(&u[i])
	}
}

func (u UserList) IsNil() bool {
	return false
}
