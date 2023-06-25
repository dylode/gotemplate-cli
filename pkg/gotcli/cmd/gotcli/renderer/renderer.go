package renderer

import (
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	rrenderer "dylode/gotemplate-cli/pkg/gotcli/renderer"

	"github.com/spf13/cobra"
)

func NewCommand(streams cmd.IOStreams) *cobra.Command {
	var jsonData string
	var yamlData string
	var isFile bool

	cmd := &cobra.Command{
		Use:   "renderer",
		Short: "Render go template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(args[0], jsonData, yamlData, isFile, streams)
		},
	}

	cmd.Flags().StringVarP(&jsonData, "json", "j", "", "json data to render the template")
	cmd.Flags().StringVarP(&yamlData, "yaml", "y", "", "yaml data to render the template")
	cmd.MarkFlagsMutuallyExclusive("json", "yaml")

	cmd.Flags().BoolVarP(&isFile, "file", "f", false, "input is a file path")

	return cmd
}

func run(input string, jsonData string, yamlData string, isFile bool, streams cmd.IOStreams) error {
	r := rrenderer.Renderer{
		Input:    input,
		DataType: rrenderer.JSON,
		Data:     jsonData,
	}

	return r.Render(streams)
}
