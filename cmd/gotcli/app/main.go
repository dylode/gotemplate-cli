package app

import (
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	"dylode/gotemplate-cli/pkg/gotcli/cmd/gotcli"
	"os"
)

func Main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

func run(args []string) error {
	cmd := gotcli.NewCommand(cmd.StandardIOStreams())
	cmd.SetArgs(args)
	return cmd.Execute()
}
