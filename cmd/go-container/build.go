package stack

import (
	"github.com/spf13/cobra"
)

var buildCommand = newBuildCommand()

func newBuildCommand() *cobra.Command {
	var (
		service string = ""
		port           = uint16(0)
		listen         = "0.0.0.0:0"
		proto          = "TCP"
	)

	command := &cobra.Command{
		Use:   "build",
		Short: "Build container",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	command.Flags().StringVar(&service, "service", service, "The target service name")
	command.MarkFlagRequired("service")

	command.Flags().Uint16Var(&port, "port", port, "The target service port")
	command.MarkFlagRequired("port")

	command.Flags().StringVar(&proto, "proto", proto, "The target service protocol (TCP or UDP)")

	command.Flags().StringVar(&listen, "listen", listen, "The listen host:port")

	return command
}
