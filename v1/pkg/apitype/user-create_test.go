package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"reflect"
	"testing"
)

func TestUserCreateRequest_NKeys(t *testing.T) {
	convey.Convey("UserCreateRequest.NKeys", t, func() {
		tmp := apitype.UserCreateRequest{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}
