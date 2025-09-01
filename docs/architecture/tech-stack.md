# Tech Stack

This section is the **DEFINITIVE** technology selection for the `helmgraph` project. It's crucial that we make specific
choices here, as these will guide all subsequent development.

I will present viable options for each category, make a recommendation based on the PRD and common best practices for Go
CLI tools, and then seek your explicit approval for each selection. We will document exact versions to ensure
consistency. This table will serve as the single source of truth for our technology choices.

Before we dive into the table, please be aware of the importance of these choices. We should look for any gaps or
disagreements with anything I suggest, and I'm here to clarify anything that's unclear. If you're unsure about any
category, I can provide suggestions with rationale.

Here are the key decisions we need to finalize:

*   **Languages and Runtimes:** The PRD specifies Go.
*   **Frameworks and Libraries / Packages:** We'll need to consider Go libraries for CLI parsing, YAML processing, and
    potentially a Helm SDK if we don't rely solely on the external `helm` CLI.
*   **Cloud provider and key services choices:** As a CLI tool, direct cloud services for its own operation are minimal,
    but we should consider if any external cloud services are implicitly needed (e.g., for fetching remote charts,
    though the PRD's explicitly state this for MVP).
*   **Database and storage solutions:** The PRD states the tool does not require a database for its own operation, but
    outputs to Neo4j.
*   **Development tools:** Standard Go development tools.

---

**Cloud Infrastructure**

-   **Provider:** N/A (CLI Tool)
-   **Key Services:** N/A (CLI Tool)
-   **Deployment Regions:** N/A (CLI Tool)

**Rationale:** As `helmgraph` is designed as a standalone command-line tool that does not require its own backend
infrastructure or persistent storage for its operation (NFR5), direct cloud infrastructure is not applicable for the
tool itself. Its output is consumed by an external Neo4j database.

---

**Technology Stack Table**

| Category | Technology | Version | Purpose | Rationale |
| :------- | :--------- | :------ | :------ | :-------- |
| **Language** | Go | 1.24.4 | Primary development language | NFR4 explicitly requires Go. It's well-suited for CLI tools, cross-platform compilation, and has a strong ecosystem for parsing and file operations. |
| **CLI Framework** | Cobra | v1.8.0 | Building robust command-line interfaces | Cobra is a popular and powerful library for creating modern CLI applications in Go, providing structure, argument parsing, and subcommands. |
| **YAML Parsing** | `gopkg.in/yaml.v3` | v3.0.1 | Parsing Kubernetes manifests (multi-document YAML) | This library is widely used and robust for handling complex YAML structures, including multi-document streams, which is essential for parsing `helm template` output (FR3). |
| **Helm Interaction** | `os/exec` (for `helm` CLI) | N/A | Executing external `helm` CLI commands | NFR7 states the tool requires the `helm` CLI. Using `os/exec` allows direct invocation of the external `helm` command for manifest generation (FR2). |
| **Graph Data Model** | Custom Go structs | N/A | Representing Kubernetes resources and relationships in-memory | A custom, lightweight Go struct model will be used to represent nodes and relationships before converting them to Cypher, ensuring flexibility and direct mapping to Cypher concepts. |
| **Cypher Generation** | Custom Go functions | N/A | Generating Cypher `CREATE`, `MERGE`, and `CONSTRAINT` statements | Given the specific nature of Cypher output (FR6, FR7), custom functions will provide precise control over syntax and property mapping, avoiding unnecessary dependencies. |
| **Linting** | `github.com/neilotoole/sq/libsq/ast/sqlparser` (or similar) | Latest stable | Linting generated Cypher output | FR8 requires linting the Cypher file. A Go-native SQL/Cypher parser/linter library will be investigated to ensure validity without external dependencies. |
| **Distribution** | Go build system | N/A | Creating standalone binaries and Docker images | Go's native build capabilities allow for easy cross-compilation to Linux, macOS, and Windows, and creating small Docker images, fulfilling NFR2 and NFR6. |

---

**Rationale for Tech Stack:**
The choices here are driven directly by the PRD's requirements (Go language, CLI tool, no internal database) and best
practices for building performant and maintainable Go applications. Cobra provides a solid foundation for the CLI,
`gopkg.in/yaml.v3` handles the complex YAML parsing, and `os/exec` integrates with the required external `helm` CLI.
Custom Go structs and functions are chosen for graph data modeling and Cypher generation to maintain full control and
minimize external dependencies. A Go-native linter will be sought for Cypher validation.
