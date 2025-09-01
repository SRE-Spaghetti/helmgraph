# Epic 2: Expanded Resource & Repository Support

**Goal:** This epic aims to extend HelmGraph's capabilities beyond local Helm charts and core Kubernetes resource
types. It will enable the service to fetch charts from remote repositories and parse a wider array of Kubernetes
resource kinds, thereby increasing the breadth of graph representations and providing more comprehensive insights into
complex deployments.

### Story 2.1 Remote Helm Repository Integration

As a DevOps Engineer,
I want HelmGraph to fetch Helm charts from remote repositories (URLs or OCI registries),
so that I can analyze charts without needing to download them locally first.

#### Acceptance Criteria

1.  The tool accepts a remote Helm repository URL or OCI registry path as input.
2.  The tool successfully fetches the specified Helm chart from the remote repository.
3.  The tool integrates the fetched chart into the existing manifest generation process (Story 1.1).
4.  The tool provides clear error messages if the remote repository is inaccessible or the chart cannot be fetched.

### Story 2.2 Expanded Kubernetes Resource Type Support

As a DevOps Engineer,
I want HelmGraph to parse additional Kubernetes resource types,
so that I can get a more complete graph representation of my deployments.

#### Acceptance Criteria

1.  The tool expands its parsing capabilities to include `StatefulSet`, `DaemonSet`, and `PersistentVolumeClaim`
    kinds.
2.  For each newly supported resource kind, the tool correctly extracts its `kind`, `metadata.name`, and
    `metadata.namespace`.
3.  The tool identifies and establishes relevant relationships for these new resource types (e.g., `StatefulSet`
    using `PersistentVolumeClaim`).
4.  The generated Cypher script (Story 1.5) includes nodes and relationships for these expanded resource types.
