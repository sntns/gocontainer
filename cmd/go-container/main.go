package main

import (
	"github.com/spf13/cobra"
)

var Command = newCommand()

func newCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "go-container",
		Short: "Tool to generate Docker/OCI container for GO",
	}
	command.AddCommand()
	return command
}

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
