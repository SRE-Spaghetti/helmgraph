package manifest

import (
	"fmt"
	"os/exec"
)

// Generate runs the 'helm template' command and returns the output.
func Generate(chartPath, releaseName, namespace, repo string) (string, error) {
	if repo != "" {
		// Fetch the chart from the repository
		pullArgs := []string{"pull", chartPath, "--repo", repo, "--untar", "--destination", ".//.helm-charts"}
		pullCmd := exec.Command("helm", pullArgs...)
		if output, err := pullCmd.CombinedOutput(); err != nil {
			return "", fmt.Errorf("failed to pull helm chart: %w\n%s", err, string(output))
		}
		chartPath = fmt.Sprintf(".//.helm-charts/%s", chartPath)
	}

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
