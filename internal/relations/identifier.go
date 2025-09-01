package relations

import "helmgraph/internal/parser"

// Relationship represents a relationship between two Kubernetes resources.
type Relationship struct {
	Source    *parser.Resource
	Target    *parser.Resource
	Type      string
	Properties map[string]interface{}
}

// Identify identifies relationships between Kubernetes resources.
func Identify(resources []*parser.Resource) []*Relationship {
	var relationships []*Relationship

	for _, r := range resources {
		if r.Kind == "Service" {
			for _, d := range resources {
				if d.Kind == "Deployment" {
					if selectorsMatch(r.Spec.Selector, d.Metadata.Labels) {
						relationships = append(relationships, &Relationship{
							Source: r,
							Target: d,
							Type:   "SELECTS",
							Properties: map[string]interface{}{
								"selector_labels": r.Spec.Selector,
							},
						})
					}
				}
			}
		}
	}

	return relationships
}

func selectorsMatch(serviceSelector, deploymentLabels map[string]string) bool {
	if len(serviceSelector) == 0 {
		return false
	}

	for k, v := range serviceSelector {
		if labelValue, ok := deploymentLabels[k]; !ok || labelValue != v {
			return false
		}
	}

	return true
}
