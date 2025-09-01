# Core Workflows

This section illustrates the key system workflows using sequence diagrams, focusing on the critical user journey of
converting a Helm chart to a Cypher script.

```mermaid
sequenceDiagram
    participant User
    participant HelmGraphCLI as HelmGraph CLI Core
    participant ManifestGen as Manifest Generator
    participant HelmCLI as Helm CLI (External)
    participant Parser as Kubernetes Manifest Parser
    participant RelIdentifier as Relationship Identifier
    participant CypherGen as Cypher Generator
    participant CypherLinter as Cypher Linter
    participant FileSystem as File System

    User->>HelmGraphCLI: 1. Invokes with chart path, release name, etc.
    HelmGraphCLI->>ManifestGen: 2. Request manifest generation
    ManifestGen->>HelmCLI: 3. Execute 'helm template' command
    HelmCLI-->>ManifestGen: 4. Returns consolidated Kubernetes manifest
    ManifestGen-->>HelmGraphCLI: 5. Returns manifest string
    HelmGraphCLI->>Parser: 6. Provide manifest for parsing
    Parser-->>HelmGraphCLI: 7. Returns parsed Kubernetes resource objects
    HelmGraphCLI->>RelIdentifier: 8. Provide parsed objects for relationship identification
    RelIdentifier-->>HelmGraphCLI: 9. Returns identified relationships
    HelmGraphCLI->>CypherGen: 10. Provide parsed objects and relationships for Cypher generation
    CypherGen-->>HelmGraphCLI: 11. Returns raw Cypher script
    HelmGraphCLI->>CypherLinter: 12. Provide Cypher script for linting
    CypherLinter-->>HelmGraphCLI: 13. Returns linting result (success/errors)
    alt Linting Successful
        HelmGraphCLI->>FileSystem: 14. Write .cypher file
        FileSystem-->>HelmGraphCLI: 15. File write confirmation
        HelmGraphCLI-->>User: 16. Operation complete, .cypher file generated
    else Linting Failed
        HelmGraphCLI-->>User: 14. Error: Cypher linting failed, show errors
    end
```

---

**Rationale for Core Workflows:**
This sequence diagram visualizes the end-to-end process of `helmgraph`, from user invocation to Cypher file output. It
clearly shows the interaction flow between the main CLI component and its internal sub-components, as well as the
interaction with the external `helm` CLI. The inclusion of the Cypher Linter step and its conditional success/failure
path highlights the importance of FR8 (linting the Cypher file) in the overall workflow. This diagram clarifies the
operational sequence and dependencies within the monolithic application.
