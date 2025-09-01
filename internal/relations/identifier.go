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
		if r.Kind == "Deployment" {
			for _, v := range r.Spec.Template.Spec.Volumes {
				if v.ConfigMap.Name != "" {
					for _, c := range resources {
						if c.Kind == "ConfigMap" && c.Metadata.Name == v.ConfigMap.Name {
							relationships = append(relationships, &Relationship{
								Source: r,
								Target: c,
								Type:   "USES_CONFIG",
								Properties: map[string]interface{}{
									"volume": v.Name,
								},
							})
						}
					}
				}
				if v.Secret.SecretName != "" {
					for _, s := range resources {
						if s.Kind == "Secret" && s.Metadata.Name == v.Secret.SecretName {
							relationships = append(relationships, &Relationship{
								Source: r,
								Target: s,
								Type:   "USES_SECRET",
								Properties: map[string]interface{}{
									"volume": v.Name,
								},
							})
						}
					}
				}
			}
			for _, c := range r.Spec.Template.Spec.Containers {
				for _, e := range c.EnvFrom {
					if e.ConfigMapRef.Name != "" {
						for _, cm := range resources {
							if cm.Kind == "ConfigMap" && cm.Metadata.Name == e.ConfigMapRef.Name {
								relationships = append(relationships, &Relationship{
									Source: r,
									Target: cm,
									Type:   "USES_CONFIG",
									Properties: map[string]interface{}{
										"envFrom": true,
									},
								})
							}
						}
					}
					if e.SecretRef.Name != "" {
						for _, s := range resources {
							if s.Kind == "Secret" && s.Metadata.Name == e.SecretRef.Name {
								relationships = append(relationships, &Relationship{
									Source: r,
									Target: s,
									Type:   "USES_SECRET",
									Properties: map[string]interface{}{
										"envFrom": true,
									},
								})
							}
						}
					}
				}
				for _, e := range c.Env {
					if e.ValueFrom.ConfigMapKeyRef.Name != "" {
						for _, cm := range resources {
							if cm.Kind == "ConfigMap" && cm.Metadata.Name == e.ValueFrom.ConfigMapKeyRef.Name {
								relationships = append(relationships, &Relationship{
									Source: r,
									Target: cm,
									Type:   "USES_CONFIG",
									Properties: map[string]interface{}{
										"env_var_name": e.Name,
									},
								})
							}
						}
					}
					if e.ValueFrom.SecretKeyRef.Name != "" {
						for _, s := range resources {
							if s.Kind == "Secret" && s.Metadata.Name == e.ValueFrom.SecretKeyRef.Name {
								relationships = append(relationships, &Relationship{
									Source: r,
									Target: s,
									Type:   "USES_SECRET",
									Properties: map[string]interface{}{
										"env_var_name": e.Name,
									},
								})
							}
						}
					}
				}
			}
		}
	if r.Kind == "StatefulSet" {
			for _, pvc := range r.Spec.VolumeClaimTemplates {
				for _, p := range resources {
					if p.Kind == "PersistentVolumeClaim" && p.Metadata.Name == pvc.Metadata.Name {
						relationships = append(relationships, &Relationship{
							Source: r,
							Target: p,
							Type:   "USES_PVC",
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
