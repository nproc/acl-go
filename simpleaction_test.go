package acl_test

import (
	"testing"

	"github.com/nproc/acl-go"
	"github.com/nproc/acl-go/driver/memory"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleAction(t *testing.T) {
	Convey("SimpleAction", t, func() {
		driver := memory.NewDriver()
		actor, err := driver.GetActor("my actor")
		So(err, ShouldBeNil)
		action, err := driver.GetAction("my action")
		So(err, ShouldBeNil)

		Convey("it should return the correct id", func() {
			So(action.String(), ShouldEqual, "my action")
		})

		Convey("it should return the correct value for .IsAllowed", func() {
			err := driver.SetDefaultPolicy(acl.Deny)
			So(err, ShouldBeNil)

			can, err := action.Allows(actor)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)

			err = driver.Set(actor, action, acl.Allow)

			can, err = action.Allows(actor)
			So(err, ShouldBeNil)
			So(can, ShouldBeTrue)

			err = driver.Set(actor, action, acl.Deny)

			can, err = action.Allows(actor)
			So(err, ShouldBeNil)
			So(can, ShouldBeFalse)
		})
	})
}