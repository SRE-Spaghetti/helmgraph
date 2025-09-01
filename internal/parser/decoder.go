package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// Selector is a custom type to handle the different structures of Kubernetes selectors.
// It can be a simple map[string]string (for Services) or a structured
// map with matchLabels (for Deployments, StatefulSets, etc.). It normalizes both
// formats into a simple map[string]string.
type Selector map[string]string

// UnmarshalYAML implements the yaml.Unmarshaler interface to handle multiple selector formats.
func (s *Selector) UnmarshalYAML(node *yaml.Node) error {
	// Case 1: Try to decode as a simple map[string]string (e.g., for a Service).
	var simpleSelector map[string]string
	if err := node.Decode(&simpleSelector); err == nil {
		*s = simpleSelector
		return nil
	}

	// Case 2: Try to decode as a structured map with matchLabels (e.g., for a Deployment).
	var structuredSelector struct {
		MatchLabels map[string]string `yaml:"matchLabels"`
	}
	if err := node.Decode(&structuredSelector); err == nil {
		*s = structuredSelector.MatchLabels
		return nil
	}

	return fmt.Errorf("failed to unmarshal selector: expected a map[string]string or a struct with matchLabels")
}
