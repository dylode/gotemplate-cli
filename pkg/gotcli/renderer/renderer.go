package renderer

import (
	"dylode/gotemplate-cli/pkg/gotcli/cmd"
	"encoding/json"
	"text/template"

	"gopkg.in/yaml.v3"
)

type DataType int

const (
	JSON DataType = iota
	YAML
)

type Renderer struct {
	Input    string
	DataType DataType
	Data     string
}

func (r Renderer) Render(streams cmd.IOStreams) error {
	data, err := r.parseData()
	if err != nil {
		return err
	}

	tmpl := template.New("gotcli")
	_, err = tmpl.Parse(r.Input)
	if err != nil {
		return err
	}

	return tmpl.Execute(streams.Out, data)
}

func (r Renderer) parseData() (any, error) {
	var data any

	switch r.DataType {
	case JSON:
		if err := json.Unmarshal([]byte(r.Data), &data); err != nil {
			return nil, err
		}
	case YAML:
		if err := yaml.Unmarshal([]byte(r.Data), &data); err != nil {
			return nil, err
		}
	}

	return data, nil
}
