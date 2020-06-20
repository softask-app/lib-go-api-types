package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"testing"
)

func TestUserDetail_NKeys(t *testing.T) {
	convey.Convey("UserDetail.NKeys", t, func() {
		tmp := apitype.UserDetail{}
		convey.So(tmp.NKeys(), convey.ShouldEqual, fieldCount(tmp))
	})
}
