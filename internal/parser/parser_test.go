package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	manifest := `
apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: default
spec:
  selector:
    app: my-app
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  namespace: default
spec:
  selector:
    matchLabels:
      app: my-app
`
	resources, err := Parse(manifest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resources) != 2 {
		t.Fatalf("expected 2 resources, but got %d", len(resources))
	}

	if resources[0].Kind != "Service" || resources[0].Metadata.Name != "my-service" {
		t.Errorf("unexpected resource[0]: %+v", resources[0])
	}

	if resources[1].Kind != "Deployment" || resources[1].Metadata.Name != "my-deployment" {
		t.Errorf("unexpected resource[1]: %+v", resources[1])
	}
}
