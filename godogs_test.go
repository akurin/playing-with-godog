package main

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

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
