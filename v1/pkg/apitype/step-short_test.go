package apitype_test

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"reflect"
	"testing"
)

func TestStepMeta_NKeys(t *testing.T) {
	convey.Convey("StepMeta.NKeys", t, func() {
		tmp := apitype.StepMeta{}
		kind := reflect.TypeOf(tmp)

		convey.So(tmp.NKeys(), convey.ShouldEqual, kind.NumField())
	})
}

func TestStepMeta_IsNil(t *testing.T) {
	convey.Convey("StepMeta.IsNil", t, func() {
		tmp := apitype.StepMeta{}
		convey.So(tmp.IsNil(), convey.ShouldBeFalse)
	})
}
