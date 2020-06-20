package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"reflect"
	"testing"
)

func TestTaskMeta_NKeys(t *testing.T) {
	convey.Convey("TaskMeta.NKeys", t, func() {
		tmp := apitype.TaskMeta{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}

func TestTaskMeta_IsNil(t *testing.T) {
	convey.Convey("TaskMeta.IsNil", t, func() {
		tmp := apitype.TaskMeta{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
