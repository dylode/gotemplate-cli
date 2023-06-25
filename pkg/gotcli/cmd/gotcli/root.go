package gotcli

import (
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	"dylode/gotemplate-cli/pkg/gotcli/cmd/gotcli/renderer"

	"github.com/spf13/cobra"
)

func NewCommand(streams cmd.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.ExactArgs(1),
		Use:   "gotcli",
		Short: "gotcli renders a template using go template syntax",
		Long:  "gotcli renders a template using go template syntax",
	}

	cmd.SetOut(streams.Out)
	cmd.SetErr(streams.ErrOut)

	cmd.AddCommand(renderer.NewCommand(streams))

	return cmd
}
