package renderer_test

import (
	"bytes"
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	"dylode/gotemplate-cli/pkg/gotcli/renderer"
	"testing"
)

func TestRenderer(t *testing.T) {
	tests := []struct {
		Name           string
		Input          string
		DataType       renderer.DataType
		Data           string
		ExpectedOutput string
		ExpectedError  error
	}{
		{
			Name:           "JSON rendering - single value",
			Input:          "Name: {{.Name}}",
			DataType:       renderer.JSON,
			Data:           `{"Name": "John"}`,
			ExpectedOutput: "Name: John",
			ExpectedError:  nil,
		},
		{
			Name:           "JSON rendering - nested object",
			Input:          "Address: {{.Address.City}}, {{.Address.Country}}",
			DataType:       renderer.JSON,
			Data:           `{"Address": {"City": "New York", "Country": "USA"}}`,
			ExpectedOutput: "Address: New York, USA",
			ExpectedError:  nil,
		},
		{
			Name:           "JSON rendering - array of values",
			Input:          "Names: {{range .Names}}{{.}}, {{end}}",
			DataType:       renderer.JSON,
			Data:           `{"Names": ["John", "Jane", "Alice"]}`,
			ExpectedOutput: "Names: John, Jane, Alice, ",
			ExpectedError:  nil,
		},
		{
			Name:           "JSON rendering - missing key",
			Input:          "Age: {{.Age}}",
			DataType:       renderer.JSON,
			Data:           `{"Name": "John"}`,
			ExpectedOutput: "Age: <no value>",
			ExpectedError:  nil,
		},
		{
			Name:           "YAML rendering - single value",
			Input:          "Name: {{.Name}}",
			DataType:       renderer.YAML,
			Data:           `Name: "John"`,
			ExpectedOutput: "Name: John",
			ExpectedError:  nil,
		},
		{
			Name:     "YAML rendering - nested object",
			Input:    "Address: {{.Address.City}}, {{.Address.Country}}",
			DataType: renderer.YAML,
			Data: `Address:
  City: "New York"
  Country: "USA"`,
			ExpectedOutput: "Address: New York, USA",
			ExpectedError:  nil,
		},
		{
			Name:     "YAML rendering - array of values",
			Input:    "Names: {{range .Names}}{{.}}, {{end}}",
			DataType: renderer.YAML,
			Data: `Names:
  - "John"
  - "Jane"
  - "Alice"`,
			ExpectedOutput: "Names: John, Jane, Alice, ",
			ExpectedError:  nil,
		},
		{
			Name:           "YAML rendering - missing key",
			Input:          "Age: {{.Age}}",
			DataType:       renderer.YAML,
			Data:           `Name: "John"`,
			ExpectedOutput: "Age: <no value>",
			ExpectedError:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			renderer := renderer.Renderer{
				Input:    test.Input,
				DataType: test.DataType,
				Data:     test.Data,
			}

			outStream := &bytes.Buffer{}

			streams := cmd.IOStreams{
				Out: outStream,
			}

			err := renderer.Render(streams)
			if err != test.ExpectedError {
				t.Errorf("expected error: %v, but got: %v", test.ExpectedError, err)
			}

			if outStream.String() != test.ExpectedOutput {
				t.Errorf("expected output: %s, but got: %s", test.ExpectedOutput, outStream.String())
			}
		})
	}
}
