package relations

import (
	"helmgraph/internal/parser"
	"testing"
)

func TestIdentify(t *testing.T) {
	resources := []*parser.Resource{
		{
			Kind: "Service",
			Metadata: parser.Metadata{
				Name: "my-service",
			},
			Spec: struct {
				Selector map[string]string `yaml:"selector"`
			}{
				Selector: map[string]string{"app": "my-app"},
			},
		},
		{
			Kind: "Deployment",
			Metadata: parser.Metadata{
				Name:   "my-deployment",
				Labels: map[string]string{"app": "my-app"},
			},
		},
		{
			Kind: "Deployment",
			Metadata: parser.Metadata{
				Name:   "another-deployment",
				Labels: map[string]string{"app": "another-app"},
			},
		},
	}

	relationships := Identify(resources)

	if len(relationships) != 1 {
		t.Fatalf("expected 1 relationship, but got %d", len(relationships))
	}

	if relationships[0].Type != "SELECTS" {
		t.Errorf("unexpected relationship type: %s", relationships[0].Type)
	}

	if relationships[0].Source.Metadata.Name != "my-service" {
		t.Errorf("unexpected relationship source: %s", relationships[0].Source.Metadata.Name)
	}

	if relationships[0].Target.Metadata.Name != "my-deployment" {
		t.Errorf("unexpected relationship target: %s", relationships[0].Target.Metadata.Name)
	}
}
