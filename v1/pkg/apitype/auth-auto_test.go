package apitype_test

import (
	"reflect"
	"testing"

	"github.com/francoispqt/gojay"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/softask-app/lib-go-api-types/v1/pkg/apitype"
	"github.com/softask-app/lib-go-token/v1/pkg/apitoken"
)

func TestAutoAuth_MarshalJSONObject(t *testing.T) {
	Convey("AutoAuth.MarshalJSONObject", t, func() {
		test := apitype.AutoAuth{
			UserId:   123,
			DeviceId: apitoken.NewToken128(),
			Token:    apitoken.NewToken256(),
		}

		bytes, err := gojay.Marshal(&test)
		So(err, ShouldBeNil)
		So(string(bytes), ShouldEqual,
			`{"userId":123,"deviceId":"`+test.DeviceId.String()+
				`","token":"`+test.Token.String()+`"}`)
	})
}

func TestAutoAuth_UnmarshalJSONObject(t *testing.T) {
	Convey("AutoAuth.UnmarshalJSONObject", t, func() {
		k1 := apitoken.NewToken128()
		k2 := apitoken.NewToken256()
		js := `{"userId":1234,"deviceId":"` + k1.String() + `","token":"` + k2.String() + `"}`

		var test apitype.AutoAuth
		So(gojay.Unmarshal([]byte(js), &test), ShouldBeNil)
		So(test.UserId, ShouldEqual, 1234)
		So(test.DeviceId, ShouldResemble, k1)
		So(test.Token, ShouldResemble, k2)
	})
}

func TestAutoAuth_NKeys(t *testing.T) {
	Convey("AutoAuth.NKeys", t, func() {
		tmp := apitype.AutoAuth{}
		kind := reflect.TypeOf(tmp)

		So(tmp.NKeys(), ShouldEqual, kind.NumField())
	})
}
