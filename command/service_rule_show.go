package main

import (
	"strings"

	"github.com/mitchellh/cli"
)

type ServiceRuleShow struct {
}

func ServiceRuleShowCommand() (cli.Command, error) {
	return &ServiceRuleShow{}, nil
}

func (c *ServiceRuleShow) Help() string {
	helpText := `
	`
	return strings.TrimSpace(helpText)
}

func (c *ServiceRuleShow) Synopsis() string {
	return "Get details about an integration belonging to a service"
}

func (c *ServiceRuleShow) Run(args []string) int {
	return 0
}
