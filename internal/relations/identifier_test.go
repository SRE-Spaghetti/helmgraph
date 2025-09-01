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
				Selector parser.Selector `yaml:"selector"`
				Template struct {
					Spec struct {
						Containers []parser.Container `yaml:"containers"`
						Volumes    []parser.Volume    `yaml:"volumes"`
					} `yaml:"spec"`
				} `yaml:"template"`
				VolumeClaimTemplates []parser.PersistentVolumeClaim `yaml:"volumeClaimTemplates"`
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
			Spec: struct {
				Selector parser.Selector `yaml:"selector"`
				Template struct {
					Spec struct {
						Containers []parser.Container `yaml:"containers"`
						Volumes    []parser.Volume    `yaml:"volumes"`
					} `yaml:"spec"`
				} `yaml:"template"`
				VolumeClaimTemplates []parser.PersistentVolumeClaim `yaml:"volumeClaimTemplates"`
			}{
				Template: struct {
					Spec struct {
						Containers []parser.Container `yaml:"containers"`
						Volumes    []parser.Volume    `yaml:"volumes"`
					} `yaml:"spec"`
				}{
					Spec: struct {
						Containers []parser.Container `yaml:"containers"`
						Volumes    []parser.Volume    `yaml:"volumes"`
					}{
						Volumes: []parser.Volume{
							{
								Name: "config",
								ConfigMap: struct {
									Name string `yaml:"name"`
								}{
									Name: "my-configmap",
								},
							},
							{
								Name: "secret",
								Secret: struct {
									SecretName string `yaml:"secretName"`
								}{
									SecretName: "my-secret",
								},
							},
						},
					},
				},
			},
		},
		{
			Kind: "ConfigMap",
			Metadata: parser.Metadata{
				Name: "my-configmap",
			},
		},
		{
			Kind: "Secret",
			Metadata: parser.Metadata{
				Name: "my-secret",
			},
		},
	}

	relationships := Identify(resources)

	if len(relationships) != 3 {
		t.Fatalf("expected 3 relationships, but got %d", len(relationships))
	}

	// Add a StatefulSet and PVC for testing
	resources = append(resources, &parser.Resource{
		Kind: "StatefulSet",
		Metadata: parser.Metadata{
			Name: "my-statefulset",
		},
		Spec: struct {
			Selector parser.Selector `yaml:"selector"`
			Template struct {
				Spec struct {
					Containers []parser.Container `yaml:"containers"`
					Volumes    []parser.Volume    `yaml:"volumes"`
				} `yaml:"spec"`
			} `yaml:"template"`
			VolumeClaimTemplates []parser.PersistentVolumeClaim `yaml:"volumeClaimTemplates"`
		}{
			VolumeClaimTemplates: []parser.PersistentVolumeClaim{
				{
					Metadata: parser.Metadata{
						Name: "my-pvc",
					},
				},
			},
		},
	}, &parser.Resource{
		Kind: "PersistentVolumeClaim",
		Metadata: parser.Metadata{
			Name: "my-pvc",
		},
	})

	relationships = Identify(resources)

	if len(relationships) != 4 {
		t.Fatalf("expected 4 relationships, but got %d", len(relationships))
	}
}
