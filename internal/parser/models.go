package parser

// Metadata represents the metadata of a Kubernetes resource.
type Metadata struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Labels    map[string]string `yaml:"labels"`
}

// Resource represents a generic Kubernetes resource.
type Resource struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       struct {
		Selector map[string]string `yaml:"selector"`
	} `yaml:"spec"`
}
