package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"testing"
)

func TestUserDetail_NKeys(t *testing.T) {
	convey.Convey("UserDetails.NKeys", t, func() {
		tmp := apitype.UserDetails{}
		convey.So(tmp.NKeys(), convey.ShouldEqual, fieldCount(tmp))
	})
}

func TestUserDetail_IsNil(t *testing.T) {
	convey.Convey("UserDetails.IsNil", t, func() {
		tmp := apitype.UserDetails{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
