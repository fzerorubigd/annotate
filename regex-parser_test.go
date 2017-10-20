package annotate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAnnotate(t *testing.T) {
	Convey("test regexp annotate", t, func() {
		c := `
		@Test {
			key1= value1
			key2 =value2
			key3 = value3
		}
		`
		g, err := LoadFromText(c)
		So(err, ShouldBeNil)
		So(len(g), ShouldEqual, 1)
		So(g[0].Name, ShouldEqual, "Test")
		So(len(g[0].Items), ShouldEqual, 3)
		So(g[0].Items["key1"], ShouldEqual, "value1")
		So(g[0].Items["key2"], ShouldEqual, "value2")
		So(g[0].Items["key3"], ShouldEqual, "value3")

	})

	Convey("test regexp annotate", t, func() {
		c := `
		@Test {}
		`
		g, err := LoadFromText(c)
		So(err, ShouldBeNil)
		So(len(g), ShouldEqual, 1)
		So(g[0].Name, ShouldEqual, "Test")
		So(len(g[0].Items), ShouldEqual, 0)
	})

	Convey("test regexp annotate", t, func() {
		c := `
		@Test {
			key
		}
		`
		_, err := LoadFromText(c)
		So(err, ShouldNotBeNil)
	})

	Convey("test regexp annotate", t, func() {
		c := `
		@Test {}

		@Test2 {
			key1= value1
			key2 =value2
			key3 = value3
		}
		`
		g, err := LoadFromText(c)
		So(err, ShouldBeNil)
		So(len(g), ShouldEqual, 2)
		So(g[0].Name, ShouldEqual, "Test")
		So(len(g[0].Items), ShouldEqual, 0)

		So(g[1].Name, ShouldEqual, "Test2")
		So(len(g[1].Items), ShouldEqual, 3)
		So(g[1].Items["key1"], ShouldEqual, "value1")
		So(g[1].Items["key2"], ShouldEqual, "value2")
		So(g[1].Items["key3"], ShouldEqual, "value3")
	})
}
