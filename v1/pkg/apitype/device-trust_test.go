package apitype_test

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
)

func TestDeviceTrustResponse_NKeys(t *testing.T) {
	convey.Convey("DeviceTrustResponse.NKeys", t, func() {
		tmp := apitype.DeviceTrustResponse{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}

func TestDeviceTrustResponse_IsNil(t *testing.T) {
	convey.Convey("DeviceTrustResponse.IsNil", t, func() {
		tmp := apitype.DeviceTrustResponse{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
