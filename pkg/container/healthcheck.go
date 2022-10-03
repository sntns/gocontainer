package container

import (
	"errors"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type (
	HealthcheckConfig struct {
		Test        []string       `json:"Test"`
		Interval    *time.Duration `json:"Interval,omitempty"`
		Timeout     *time.Duration `json:"Timeout,omitempty"`
		StartPeriod *time.Duration `json:"StartPeriod,omitempty"`
		Retries     int            `json:"Retries,omitempty"`
	}
)

func (c *Container) SetHealthcheck(cmd []string, interval, timeout, start *time.Duration, retries int) error {
	c.Healthcheck = &HealthcheckConfig{
		Test:        append([]string{"CMD"}, cmd...),
		Interval:    interval,
		Timeout:     timeout,
		StartPeriod: start,
		Retries:     retries,
	}
	return nil
}
func (c *Container) SetHealthcheckFromArg(arg string) error {
	health, err := parseHealthcheckArg(arg)
	if err != nil {
		return err
	}
	c.Healthcheck = health
	return nil
}

func parseHealthcheckArg(arg string) (*HealthcheckConfig, error) {
	var interval, timeout, start time.Duration
	var retries int
	var health *HealthcheckConfig
	cmd := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			list := []string{}
			if len(args) < 3 {
				return errors.New("invalid healthcheck value")
			}
			if args[0] == "CMD" && args[1] == "[" && args[len(args)-1] == "]" {
				list = append(list, "CMD")
				for _, item := range args[2 : len(args)-1] {
					list = append(list, strings.Trim(item, "\""))
				}
			} else {
				return errors.New("invalid healthcheck value")
			}

			health = &HealthcheckConfig{
				Test:    list,
				Retries: retries,
			}
			if interval != time.Duration(0) {
				health.Interval = &interval
			}
			if timeout != time.Duration(0) {
				health.Timeout = &timeout
			}
			if start != time.Duration(0) {
				health.StartPeriod = &start
			}
			return nil
		},
	}
	cmd.Flags().DurationVar(&interval, "interval", interval, "")
	cmd.Flags().DurationVar(&timeout, "timeout", timeout, "")
	cmd.Flags().DurationVar(&start, "start-period", start, "")
	cmd.Flags().IntVar(&retries, "retries", retries, "")

	arg = strings.Replace(arg, "CMD", "-- CMD", 1)
	cmd.SetArgs(strings.Split(arg, " "))
	err := cmd.Execute()
	if err != nil {
		return nil, err
	}
	return health, nil
}
