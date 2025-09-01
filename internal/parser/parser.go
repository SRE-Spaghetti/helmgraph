package parser

import (
	"bytes"
	"io"

	"gopkg.in/yaml.v3"
)

// Parse takes a multi-document YAML manifest and returns a slice of Resource objects.
func Parse(manifest string) ([]*Resource, error) {
	var resources []*Resource
	decoder := yaml.NewDecoder(bytes.NewReader([]byte(manifest)))

	for {
		var resource Resource
		if err := decoder.Decode(&resource); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		resources = append(resources, &resource)
	}

	return resources, nil
}
