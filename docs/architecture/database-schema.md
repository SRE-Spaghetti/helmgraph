# Database Schema

This section defines how the conceptual data models will be transformed into concrete database schemas. Given that
`helmgraph` is a command-line tool that outputs Cypher queries for Neo4j (NFR5, NFR8), it does not maintain its own
internal database. Therefore, this section will focus on the **structure of the Cypher output** that is intended for a
Neo4j graph database.

The Cypher output will define nodes and relationships based on the Kubernetes resources and their connections identified
in the manifest.

**Node Definitions (Cypher Labels and Properties):**

*   **Generic Kubernetes Resource Node:**
    ```cypher
    (:KubernetesResource {kind: "...", name: "...", namespace: "..."})
    ```
    *   **Rationale:** This generic label allows for flexible representation of all Kubernetes `kind`s as required by
        FR3, even those not explicitly known in advance.

*   **Specific Resource Nodes (Examples):**
    *   **Deployment:**
        ```cypher
        (:Deployment:KubernetesResource {kind: "Deployment", name: "my-app-deployment", namespace: "default", labels: {app: "my-app"}})
        ```
    *   **Service:**
        ```cypher
        (:Service:KubernetesResource {kind: "Service", name: "my-app-service", namespace: "default", selector: {app: "my-app"}})
        ```
    *   **ConfigMap:**
        ```cypher
        (:ConfigMap:KubernetesResource {kind: "ConfigMap", name: "my-app-config", namespace: "default"})
        ```
    *   **Secret:**
        ```cypher
        (:Secret:KubernetesResource {kind: "Secret", name: "my-app-secret", namespace: "default"})
        ```
    *   **Rationale:** Specific labels (e.g., `:Deployment`, `:Service`) provide more granular typing in Neo4j,
        allowing for more precise queries and visualizations. All specific resource nodes will also carry the
        `:KubernetesResource` label for broader queries.

**Relationship Definitions (Cypher Types and Properties):**

*   **Service to Deployment (SELECTS):**
    ```cypher
    (:Service)-[:SELECTS {selector_labels: {app: "my-app"}}]->(:Deployment)
    ```
    *   **Rationale:** Directly implements Story 1.3, showing how a Service routes traffic to a Deployment based on
        selectors.

*   **Deployment to ConfigMap (USES_CONFIG):**
    ```cypher
    (:Deployment)-[:USES_CONFIG {mount_path: "/etc/config", env_var_name: "CONFIG_FILE"}]->(:ConfigMap)
    ```
    *   **Rationale:** Directly implements Story 1.4, showing a Deployment's dependency on a ConfigMap. Properties like
        `mount_path` or `env_var_name` provide context.

*   **Deployment to Secret (USES_SECRET):**
    ```cypher
    (:Deployment)-[:USES_SECRET {mount_path: "/etc/secrets", env_var_name: "DB_PASSWORD"}]->(:Secret)
    ```
    *   **Rationale:** Directly implements Story 1.4, showing a Deployment's dependency on a Secret. Properties like
        `mount_path` or `env_var_name` provide context.

**Constraint Statements (FR7):**

The generated `.cypher` file will include `CREATE CONSTRAINT` statements to ensure data integrity and optimize query
performance in Neo4j. Examples:

```cypher
CREATE CONSTRAINT IF NOT EXISTS FOR (n:KubernetesResource) REQUIRE (n.kind, n.name, n.namespace) IS UNIQUE;
CREATE CONSTRAINT IF NOT EXISTS FOR (n:Deployment) REQUIRE (n.name, n.namespace) IS UNIQUE;
CREATE CONSTRAINT IF NOT EXISTS FOR (n:Service) REQUIRE (n.name, n.namespace) IS UNIQUE;
CREATE CONSTRAINT IF NOT EXISTS FOR (n:ConfigMap) REQUIRE (n.name, n.namespace) IS UNIQUE;
CREATE CONSTRAINT IF NOT EXISTS FOR (n:Secret) REQUIRE (n.name, n.namespace) IS UNIQUE;
```
*   **Rationale:** These constraints enforce uniqueness for key node properties, which is crucial for `MERGE` operations
    and efficient graph traversal in Neo4j.

---

**Rationale for Database Schema:**
This section clarifies the structure of the Cypher output, which serves as the "database schema" for the external Neo4j
instance. It directly maps the conceptual data models to concrete Cypher syntax for nodes, relationships, and
constraints, fulfilling FR6, FR7, Story 1.5, and the overall goal of generating a usable graph representation. The
examples provided illustrate how different Kubernetes resource types and their connections will be translated into graph
elements.
