Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

  Scenario: Eat 1 out of 10
    Given there are 10 godogs
    When I eat 1
    Then there should be 9 remaining