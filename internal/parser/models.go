package parser

// Metadata represents the metadata of a Kubernetes resource.
type Metadata struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Labels    map[string]string `yaml:"labels"`
}

// VolumeMount represents a mounting of a Volume within a container.
type VolumeMount struct {
	Name      string `yaml:"name"`
	MountPath string `yaml:"mountPath"`
}

// EnvFromSource represents the source of a set of environment variables.
type EnvFromSource struct {
	ConfigMapRef struct {
		Name string `yaml:"name"`
	} `yaml:"configMapRef"`
	SecretRef struct {
		Name string `yaml:"name"`
	} `yaml:"secretRef"`
}

// EnvVar represents an environment variable present in a Container.
type EnvVar struct {
	Name      string `yaml:"name"`
	ValueFrom struct {
		ConfigMapKeyRef struct {
			Name string `yaml:"name"`
			Key  string `yaml:"key"`
		} `yaml:"configMapKeyRef"`
		SecretKeyRef struct {
			Name string `yaml:"name"`
			Key  string `yaml:"key"`
		} `yaml:"secretKeyRef"`
	} `yaml:"valueFrom"`
}

// Container represents a single container that is expected to be run on a pod.
type Container struct {
	Name         string          `yaml:"name"`
	Env          []EnvVar        `yaml:"env"`
	EnvFrom      []EnvFromSource `yaml:"envFrom"`
	VolumeMounts []VolumeMount   `yaml:"volumeMounts"`
}

// Volume represents a named volume in a pod that is accessible to containers.
type Volume struct {
	Name   string `yaml:"name"`
	Secret struct {
		SecretName string `yaml:"secretName"`
	} `yaml:"secret"`
	ConfigMap struct {
		Name string `yaml:"name"`
	} `yaml:"configMap"`
}

// PersistentVolumeClaim represents a PersistentVolumeClaim.
type PersistentVolumeClaim struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
}

// Resource represents a generic Kubernetes resource.
type Resource struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       struct {
		Selector Selector `yaml:"selector"`
		Template struct {
			Spec struct {
				Containers []Container `yaml:"containers"`
				Volumes    []Volume    `yaml:"volumes"`
			} `yaml:"spec"`
		} `yaml:"template"`
		VolumeClaimTemplates []PersistentVolumeClaim `yaml:"volumeClaimTemplates"`
	} `yaml:"spec"`
}
