package miro

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParseQueryTags(t *testing.T) {
	Convey("Given a struct with query tags", t, func() {
		type testStruct struct {
			Foo string `query:"foo,omitempty"`
			Bar int    `query:"bar"`
			Baz string `query:"baz,omitempty"`
			Qux string `query:"qux"`
		}
		input := testStruct{
			Foo: "hello",
			Bar: 42,
			Baz: "",
			Qux: "",
		}
		expected := []Parameter{
			{"foo": "hello"},
			{"bar": "42"},
			{"qux": ""},
		}

		Convey("When parseQueryTags is called", func() {
			result := parseQueryTags(input)

			Convey("The result should match the expected output", func() {
				So(result, ShouldResemble, expected)
			})
		})
	})

	Convey("Given a struct with no query tags", t, func() {
		type testStruct struct {
			Foo string
			Bar int
			Baz string
		}
		input := testStruct{
			Foo: "hello",
			Bar: 42,
			Baz: "",
		}
		expected := []Parameter{}

		Convey("When parseQueryTags is called", func() {
			result := parseQueryTags(input)

			Convey("The result should be an empty slice", func() {
				So(result, ShouldResemble, expected)
			})
		})
	})
}
