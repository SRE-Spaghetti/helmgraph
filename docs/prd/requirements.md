# Requirements

### Functional

*   FR1: The service shall accept a local Helm chart path as input - either tgz format or directory.
*   FR1a: The service shall accept an additional YAML values file path as input.
*   FR1b: The service shall require a release name as input.
*   FR1c: The service shall optionally accept a namespace as input.
*   FR2: The service shall execute `helm template` to generate a consolidated Kubernetes manifest.
*   FR3: The service shall parse the manifest to create Cypher nodes for `Deployment`, `Service`, `ConfigMap`,
    `Secret`, and all other kinds, not all of which might be known in advance.
*   FR4: Nodes shall be identified by their `kind`, `metadata.name`, and optionally `metadata.namespace` (only if
    specified in the manifest).
*   FR5: The service shall create Cypher relationships for common connections (e.g., `Service` selecting
    `Deployment`, `Deployment` mounting `Secret`).
*   FR6: The service shall produce a valid `.cypher` file containing all generated `CREATE` and `MERGE` commands.
*   FR7: The service shall produce CONSTRAINT statements in the `.cypher` file for the Node types it produces.
*   FR8: The service shall lint the `.cypher` file to ensure it is valid.

### Non Functional

*   NFR1: The service shall be a command-line tool.
*   NFR2: The service shall run on Linux, macOS, and Windows (via WSL).
*   NFR3: The service shall process a Helm chart with approximately 50 resources in under 30 seconds.
*   NFR4: The service shall be developed using Go.
*   NFR5: The service shall not require a database for its own operation.
*   NFR6: The service shall be distributed as a binary and a Docker image.
*   NFR7: The service shall require the `helm` CLI to be installed and available in the user's PATH or else use the
    Helm programming level API to perform the `template` action.
*   NFR8: The tool's functionality is dependent on the output of the `helm template` command.
