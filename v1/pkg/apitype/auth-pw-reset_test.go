package apitype_test

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
)

func TestPasswordResetRequest_NKeys(t *testing.T) {
	convey.Convey("PasswordResetRequest.NKeys", t, func() {
		tmp := apitype.PasswordResetRequest{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}

func TestRequestPasswordResetRequest_NKeys(t *testing.T) {
	convey.Convey("RequestPasswordResetRequest.NKeys", t, func() {
		tmp := apitype.RequestPasswordResetRequest{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}

func TestPasswordResetRequest_IsNil(t *testing.T) {
	convey.Convey("PasswordResetRequest.IsNil", t, func() {
		tmp := apitype.PasswordResetRequest{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}

func TestRequestPasswordResetRequest_IsNil(t *testing.T) {
	convey.Convey("RequestPasswordResetRequest.IsNil", t, func() {
		tmp := apitype.RequestPasswordResetRequest{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
