package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse takes a multi-document YAML manifest and returns a slice of Resource objects.
func Parse(manifest string) ([]*Resource, error) {
	var resources []*Resource
	decoder := yaml.NewDecoder(bytes.NewReader([]byte(manifest)))
	resourceNum := 0

	for {
		var resource Resource
		if err := decoder.Decode(&resource); err != nil {
			if err == io.EOF {
				break
			}

			lines := strings.Split(manifest, "\n")
			builder := strings.Builder{}
			builder.WriteString("\n")
			startNum := 0
			endNum := len(lines)
			if strings.Contains(err.Error(), "unmarshal errors:") {
				resourceStart := 0
				resourceCount := 0
				for i, line := range lines {
					if line == "---" {
						if resourceCount < resourceNum {
							resourceCount++
							continue
						}
						resourceStart = i
						break
					}
				}

				re := regexp.MustCompile(`line (\d+):`)
				matches := re.FindStringSubmatch(err.Error())
				if len(matches) > 1 {
					if num, err := strconv.Atoi(matches[1]); err == nil {
						if resourceStart+num-3 > 0 {
							startNum = resourceStart + num - 3
						}
						if resourceStart+num+3 < len(lines) {
							endNum = resourceStart + num + 3
						}
					}
				}
				for i := startNum; i < endNum; i++ {
					builder.WriteString(fmt.Sprintf("%d. %s\n", i+1-resourceStart, lines[i]))
				}
			}

			return nil, fmt.Errorf("error decoding YAML resource %d: %w\n\nManifest:\n %s", resourceNum, err, builder.String())
		}
		resources = append(resources, &resource)
		resourceNum++
	}

	return resources, nil
}
