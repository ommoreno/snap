package control

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/intelsdilabs/pulse/control/plugin"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Mock Executor used to test
type MockPluginExecutor struct {
	Killed          bool
	Response        string
	WaitTime        time.Duration
	WaitError       error
	WaitForResponse func(time.Duration) (*plugin.Response, error)
}

var (
	PluginName = "pulse-collector-dummy"
	PulsePath  = os.Getenv("PULSE_PATH")
	PluginPath = path.Join(PulsePath, "plugin", "collector", "pulse-collector-dummy")
)

// Uses the dummy collector plugin to simulate Loading
func TestLoad(t *testing.T) {
	// These tests only work if PULSE_PATH is known.
	// It is the responsibility of the testing framework to
	// build the plugins first into the build dir.
	if PulsePath != "" {
		Convey("pluginControl.Load", t, func() {

			Convey("loads successfully", func() {
				c := Control()
				c.Start()
				loadedPlugin, err := c.Load(PluginPath)

				So(loadedPlugin, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})

			Convey("returns error if not started", func() {
				c := Control()
				loadedPlugin, err := c.Load(PluginPath)

				So(loadedPlugin, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

			Convey("adds to pluginControl.LoadedPlugins on successful load",
				func() {
					c := Control()
					c.Start()
					loadedPlugin, err := c.Load(PluginPath)

					So(loadedPlugin, ShouldNotBeNil)
					So(err, ShouldBeNil)
					So(len(c.LoadedPlugins), ShouldBeGreaterThan, 0)
				})

		})

	} else {
		fmt.Printf("PULSE_PATH not set. Cannot test %s plugin.\n", PluginName)
	}
}

func TestStop(t *testing.T) {
	Convey("pluginControl.Stop", t, func() {
		c := Control()
		c.Start()
		c.Stop()

		Convey("stops", func() {
			So(c.Started, ShouldBeFalse)
		})

	})

}