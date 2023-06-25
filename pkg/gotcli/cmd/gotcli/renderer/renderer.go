package renderer

import (
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	rrenderer "dylode/gotemplate-cli/pkg/gotcli/renderer"
	"errors"

	"github.com/spf13/cobra"
)

func NewCommand(streams cmd.IOStreams) *cobra.Command {
	var jsonData string
	var yamlData string

	cmd := &cobra.Command{
		Use:   "render",
		Short: "Render go template",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(args[0], jsonData, yamlData, streams)
		},
	}

	cmd.Flags().StringVarP(&jsonData, "json", "j", "", "json data to render the template")
	cmd.Flags().StringVarP(&yamlData, "yaml", "y", "", "yaml data to render the template")
	cmd.MarkFlagsMutuallyExclusive("json", "yaml")

	return cmd
}

func run(input string, jsonData string, yamlData string, streams cmd.IOStreams) error {
	r := rrenderer.Renderer{
		Input: input,
	}

	if err := readData(&r, jsonData, yamlData); err != nil {
		return err
	}

	return r.Render(streams)
}

func readData(r *rrenderer.Renderer, jsonData string, yamlData string) error {
	if len(jsonData) != 0 {
		r.DataType = rrenderer.JSON
		r.Data = jsonData
	} else if len(yamlData) != 0 {
		r.DataType = rrenderer.YAML
		r.Data = yamlData
	} else {
		return errors.New("could not determine data type")
	}

	return nil
}
