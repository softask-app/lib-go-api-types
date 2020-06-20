package apitype_test

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
)

func TestDeviceDetails_NKeys(t *testing.T) {
	convey.Convey("DeviceDetails.NKeys", t, func() {
		tmp := apitype.DeviceDetails{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}

func TestDeviceDetails_IsNil(t *testing.T) {
	convey.Convey("DeviceDetails.IsNil", t, func() {
		tmp := apitype.DeviceDetails{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
