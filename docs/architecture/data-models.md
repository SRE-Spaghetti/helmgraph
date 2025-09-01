# Data Models

This section defines the core data models and entities that `helmgraph` will extract from Kubernetes manifests and
represent as nodes and relationships in the Neo4j graph. These models are conceptual and will guide the in-memory
representation before conversion to Cypher.

We will start with a generic Kubernetes Resource model, and then define specific examples based on the PRD.

---

**Purpose:** Represents any Kubernetes resource identified in the manifest. This serves as the base for all specific
resource kinds.

**Key Attributes:**
- `kind`: String - The Kubernetes API `kind` of the resource (e.g., "Deployment", "Service").
- `name`: String - The `metadata.name` of the resource.
- `namespace`: String (Optional) - The `metadata.namespace` of the resource, if specified.

**Relationships:**
- Can have relationships like `SELECTS`, `USES_CONFIG`, `USES_SECRET`, etc., with other Kubernetes resources.

---

**Purpose:** Represents a Kubernetes `Deployment` resource.

**Key Attributes:**
- `kind`: String - "Deployment"
- `name`: String - The name of the Deployment.
- `namespace`: String (Optional) - The namespace of the Deployment.
- `labels`: Map (Optional) - Key-value pairs from `metadata.labels`, used for selectors.

**Relationships:**
- Can `USES_CONFIG` a `ConfigMap`.
- Can `USES_SECRET` a `Secret`.
- Can be `SELECTED_BY` a `Service`.

---

**Purpose:** Represents a Kubernetes `Service` resource.

**Key Attributes:**
- `kind`: String - "Service"
- `name`: String - The name of the Service.
- `namespace`: String (Optional) - The namespace of the Service.
- `selector`: Map (Optional) - Key-value pairs from `spec.selector`, used to select pods/deployments.

**Relationships:**
- Can `SELECTS` a `Deployment` (or other workload).

---

**Purpose:** Represents a Kubernetes `ConfigMap` resource.

**Key Attributes:**
- `kind`: String - "ConfigMap"
- `name`: String - The name of the ConfigMap.
- `namespace`: String (Optional) - The namespace of the ConfigMap.

**Relationships:**
- Can be `USED_BY` a `Deployment` (or other workload).

---

**Purpose:** Represents a Kubernetes `Secret` resource.

**Key Attributes:**
- `kind`: String - "Secret"
- `name`: String - The name of the Secret.
- `namespace`: String (Optional) - The namespace of the Secret.

**Relationships:**
- Can be `USED_BY` a `Deployment` (or other workload).

---

**Rationale for Data Models:**
These conceptual models are derived directly from FR3, FR4, FR5, Story 1.3, and Story 1.4 of the PRD. They define the
essential attributes for each node type (`kind`, `name`, `namespace`) and the primary relationships that `helmgraph`
needs to identify. The generic "Kubernetes Resource" allows for extensibility to other `kind`s as required by FR3. The
attributes are chosen to directly map to properties in Neo4j nodes and relationships.
