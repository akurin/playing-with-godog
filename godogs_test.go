package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v10"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opts)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

type state struct {
	Godogs int
}

func (s *state) thereAreGodogs(available int) error {
	s.Godogs = available
	return nil
}

func (s *state) iEat(num int) error {
	if s.Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, s.Godogs)
	}
	s.Godogs -= num
	return nil
}

func (s *state) thereShouldBeRemaining(remaining int) error {
	if s.Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, s.Godogs)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	state := state{}

	s.BeforeScenario(func(*messages.Pickle) {
		state.Godogs = 0
	})

	s.Step(`^there are (\d+) godogs$`, state.thereAreGodogs)
	s.Step(`^I eat (\d+)$`, state.iEat)
	s.Step(`^there should be (\d+) remaining$`, state.thereShouldBeRemaining)
}
