package apitype

import "errors"

func errBadKey(k string) error {
	return errors.New("unrecognized JSON key" + k)
}
