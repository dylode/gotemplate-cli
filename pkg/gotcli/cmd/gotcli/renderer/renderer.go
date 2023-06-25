package renderer

import (
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	rrenderer "dylode/gotemplate-cli/pkg/gotcli/renderer"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func NewCommand(streams cmd.IOStreams) *cobra.Command {
	var jsonData string
	var yamlData string
	var isFile bool

	cmd := &cobra.Command{
		Use:   "render",
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
	r := rrenderer.Renderer{}

	err := readInput(&r, input, isFile)
	if err != nil {
		return err
	}

	err = readData(&r, jsonData, yamlData)
	if err != nil {
		return err
	}

	return r.Render(streams)
}

func readInput(r *rrenderer.Renderer, input string, isFile bool) error {
	if isFile {
		fileContents, err := os.ReadFile(r.Input)
		if err != nil {
			return err
		}

		r.Input = string(fileContents)
	} else {
		r.Input = input
	}

	return nil
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
