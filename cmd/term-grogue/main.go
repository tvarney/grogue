package main

import (
	"fmt"
	"os"

	"github.com/tvarney/grogue/pkg/drivers/terminal"
	"github.com/tvarney/grogue/pkg/game"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(args []string) error {
	app := game.New()
	driver := terminal.New()
	driver.SetLog("debug.log")

	if err := driver.Init(); err != nil {
		return err
	}

	driver.Draw(app)
	for app.Running {
		switch app.Update(driver.PollAction(app)) {
		case game.RenderFull:
			driver.Clear()
			driver.Draw(app)
		case game.RenderIncremental:
			driver.Draw(app)
		}
	}
	driver.Finalize()
	return nil
}
