package manifest

import (
	"fmt"
	"os/exec"
)

// Generate runs the 'helm template' command and returns the output.
func Generate(chartPath, releaseName, namespace string) (string, error) {
	args := []string{"template", releaseName, chartPath}
	if namespace != "" {
		args = append(args, "--namespace", namespace)
	}

	cmd := exec.Command("helm", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run helm template: %w\n%s", err, string(output))
	}

	return string(output), nil
}
