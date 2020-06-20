package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"reflect"
	"testing"
)

func TestTaskCreateRequest_NKeys(t *testing.T) {
	convey.Convey("TaskCreateRequest.NKeys", t, func() {
		tmp := apitype.TaskCreateRequest{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}
