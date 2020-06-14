package main

import (
	"fmt"

	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages-go/v10"
)

var Godogs int

func thereAreGodogs(available int) error {
	Godogs = available
	return nil
}

func iEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

// godog v0.9.0 (latest) and earlier
func FeatureContext(s *godog.Suite) {
	s.BeforeSuite(func() { Godogs = 0 })

	s.BeforeScenario(func(*messages.Pickle) {
		Godogs = 0 // clean the state before every scenario
	})

	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}
