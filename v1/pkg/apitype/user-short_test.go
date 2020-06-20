package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"testing"
)

func TestUserMeta_NKeys(t *testing.T) {
	convey.Convey("UserMeta.NKeys", t, func() {
		tmp := apitype.UserMeta{}
		convey.So(tmp.NKeys(), convey.ShouldEqual, fieldCount(tmp))
	})
}

func TestUserMeta_IsNil(t *testing.T) {
	convey.Convey("UserMeta.IsNil", t, func() {
		tmp := apitype.UserMeta{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
