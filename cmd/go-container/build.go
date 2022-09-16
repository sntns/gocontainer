package main

import (
	"github.com/sntns/go-container/pkg/container"
	"github.com/spf13/cobra"
)

const TEMPDIR_FMT = "go-container-%s"

type CommonOptions struct {
	outdir string
}

func setCommonFlags(command *cobra.Command, opts *CommonOptions) {
	command.Flags().StringVar(&opts.outdir, "outdir", opts.outdir, "The OCI directory for container")
	command.MarkFlagRequired("outdir")
}

var buildCommand = func() *cobra.Command {
	var (
		common   CommonOptions
		binaries []string
		labels   []string
		copies   []string
	)

	command := &cobra.Command{
		Use:   "build",
		Short: "Build new container",
		RunE: func(cmd *cobra.Command, args []string) error {
			cont, err := container.New()
			if err != nil {
				return err
			}

			err = cont.SetLabels(labels...)
			if err != nil {
				return err
			}

			err = cont.Copy(copies...)
			if err != nil {
				return err
			}

			err = cont.AddBinary(binaries)
			if err != nil {
				return err
			}

			err = cont.Save(common.outdir)
			if err != nil {
				return err
			}
			return nil
		},
	}
	setCommonFlags(command, &common)

	command.Flags().StringSliceVar(&binaries, "binary", binaries, "Binary to include in container")
	command.Flags().StringSliceVar(&copies, "copy", copies, "Copy directory or file to container (format: <file|dir>[:<destination>]")
	command.Flags().StringSliceVar(&labels, "label", labels, "Add container label")

	return command
}()
