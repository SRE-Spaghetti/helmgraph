# Epic 1: Core HelmGraph Functionality

**Goal:** This epic aims to establish the foundational service for HelmGraph. It will enable users to input a local
Helm chart, process its templates to generate a Kubernetes manifest, and then convert that manifest into a set of
Cypher queries. These queries will represent core Kubernetes resources as nodes and their key relationships as edges,
ultimately producing a `.cypher` script ready for import into a Neo4j graph database. This delivers the core value
proposition of visualizing Helm chart deployments.

### Story 1.1 Helm Chart Input & Manifest Generation

As a DevOps Engineer,
I want to provide a local Helm chart path as input,
so that HelmGraph can generate a consolidated Kubernetes manifest file.

#### Acceptance Criteria

1. The tool accepts a local file path to a Helm chart directory as an argument.
1. The tool accepts a release name and optionally a namespace as input.
1. The tool optionally accepts an output file name. If not present it sends the output to standard out
1. The tool successfully executes the `helm template` command using the provided chart path.
1. The tool captures the complete output of `helm template` as a single, consolidated Kubernetes manifest string.
1. The tool provides clear error messages if the provided chart path is invalid or if `helm template` execution
    fails.

### Story 1.2 Core Kubernetes Node Parsing

As a DevOps Engineer,
I want HelmGraph to parse the Kubernetes manifest,
so that it identifies and extracts core resource types as graph nodes.

#### Acceptance Criteria

1.  The tool parses the consolidated Kubernetes manifest, correctly handling multi-document YAML streams.
2.  The tool identifies and extracts resources of all `kind`s: e.g. `Deployment`, `Service`, `ConfigMap`,
    `Secret`.
3.  For each identified resource, the tool extracts its `kind`, `metadata.name`, and optionally `metadata.namespace`.
4.  The extracted data for each resource is structured appropriately for conversion into a Neo4j node.

### Story 1.3 Relationship Identification (Service & Deployment)

As a DevOps Engineer,
I want HelmGraph to identify relationships between Services and Deployments,
so that I can visualize how Services route traffic to specific application instances.

#### Acceptance Criteria

1.  The tool analyzes `Service` resources to identify their `selector` labels.
1.  The tool matches `Service` selectors to `Deployment` resources that have matching labels.
1.  For each identified match, the tool establishes a "SELECTS" relationship between the `Service` node and the
    `Deployment` node.
1.  The relationship includes relevant properties, such as the `selector_labels` used for the connection.

### Story 1.4 Relationship Identification (Deployment & ConfigMap/Secret)

As a DevOps Engineer,
I want HelmGraph to identify relationships between Deployments and ConfigMaps/Secrets,
so that I can understand which configurations and sensitive data are used by my applications.

#### Acceptance Criteria

1.  The tool analyzes `Deployment` resources to identify references to `ConfigMap` and `Secret` resources.
1.  The tool identifies references through common mechanisms like `volumeMounts`, `envFrom`, and `env` variables.
1.  For each identified reference, the tool establishes a "USES_CONFIG" relationship between the `Deployment` node
    and the `ConfigMap` node, or a "USES_SECRET" relationship between the `Deployment` node and the `Secret` node.
1.  The relationships include relevant properties, such as `mount_path` or `env_var_name`, where applicable.

### Story 1.5 Cypher Script Generation & Output

As a DevOps Engineer,
I want HelmGraph to generate a Cypher script,
so that I can import the Helm chart's graph representation into Neo4j.

#### Acceptance Criteria

1.  The tool generates a valid `.cypher` file containing Cypher `CREATE` or `MERGE` commands.
2.  The `.cypher` file includes commands for all identified nodes (`Deployment`, `Service`, `ConfigMap`, `Secret` etc.)
    with their respective properties (`kind`, `name`, `namespace`).
3.  The `.cypher` file includes commands for all identified relationships (e.g., "SELECTS", "USES_CONFIG",
    "USES_SECRET") with their respective properties.
4.  The generated `.cypher` file is saved to a user-specified or default output location.
5.  The generated Cypher is syntactically correct and can be executed against a Neo4j database.
