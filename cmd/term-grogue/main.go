package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tvarney/grogue/pkg/drivers/terminal"
	"github.com/tvarney/grogue/pkg/game"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run(args []string) error {
	debug := kingpin.Flag("debug", "enable debug logging").Short('D').Bool()
	seed := kingpin.Flag("seed", "world random seed").Short('s').Default(strconv.FormatInt(time.Now().Unix(), 10)).Int64()
	_ = kingpin.Parse()

	driver := terminal.New()

	if *debug {
		driver.SetLog("debug.log")
	}

	if err := driver.Init(); err != nil {
		return err
	}

	log.Printf("Starting term-grogue")
	app := game.New(*seed)

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
