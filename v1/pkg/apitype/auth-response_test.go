package apitype_test

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
)

func TestAuthResponse_NKeys(t *testing.T) {
	convey.Convey("AuthResponse.NKeys", t, func() {
		tmp := apitype.AuthResponse{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}
