package apitype_test

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
)

func TestLoginRequest_NKeys(t *testing.T) {
	Convey("LoginRequest.NKeys", t, func() {
		tmp := apitype.LoginRequest{}
		kind := reflect.TypeOf(tmp)

		So(tmp.NKeys(), ShouldEqual, kind.NumField())
	})
}

func TestLoginRequest_IsNil(t *testing.T) {
	Convey("LoginRequest.IsNil", t, func() {
		tmp := apitype.LoginRequest{}
		So(tmp.IsNil(), ShouldBeFalse)
	})
}
