package manifest

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Create a dummy chart directory for testing
	err := os.MkdirAll("testdata/mychart/templates", 0755)
	if err != nil {
		t.Fatalf("failed to create test chart directory: %v", err)
	}
	defer os.RemoveAll("testdata")

	// Create a dummy Chart.yaml file
	chart := `apiVersion: v2
name: mychart
description: A Helm chart for Kubernetes
type: application
version: 0.1.0
appVersion: "1.16.0"`

	err = os.WriteFile("testdata/mychart/Chart.yaml", []byte(chart), 0644)
	if err != nil {
		t.Fatalf("failed to write test template file: %v", err)
	}

	// Create a dummy template file
	template := `
apiVersion: v1
kind: Service
metadata:
  name: {{ .Release.Name }}-nginx
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
      name: http
  selector:
    app: nginx
`
	err = os.WriteFile("testdata/mychart/templates/service.yaml", []byte(template), 0644)
	if err != nil {
		t.Fatalf("failed to write test template file: %v", err)
	}

	// Test case 1: Successful generation
	output, err := Generate("testdata/mychart", "my-release", "default")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedOutput := `---
# Source: mychart/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-release-nginx
spec:
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
      name: http
  selector:
    app: nginx
`
	if output != expectedOutput {
		t.Errorf("unexpected output.\nGot:\n%s\nExpected:\n%s", output, expectedOutput)
	}

	// Test case 2: Chart not found
	_, err = Generate("testdata/nonexistent-chart", "my-release", "default")
	if err == nil {
		t.Error("expected an error, but got nil")
	}
}
