package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These will be set by goreleaser
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var Command = func() *cobra.Command {
	command := &cobra.Command{
		Use:   "go-container",
		Short: "Tool to generate Docker/OCI container for GO",
		Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
	}
	command.AddCommand(
		buildCommand,
	)
	return command
}()

func main() {
	/*
		f, err := os.Create("cpu.prof")
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	*/
	if err := Command.Execute(); err != nil {
		panic(err)
	}

	/*
		f, err = os.Create("mem.prof")
		if err != nil {
			panic(err)
		}
		pprof.WriteHeapProfile(f)
	*/
}
