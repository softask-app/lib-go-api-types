package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"testing"
)

func TestTaskDetails_NKeys(t *testing.T) {
	convey.Convey("TaskDetails.NKeys", t, func() {
		tmp := apitype.TaskDetails{}
		convey.So(tmp.NKeys(), convey.ShouldEqual, fieldCount(tmp))
	})
}
