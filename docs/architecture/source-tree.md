# Source Tree

This section defines the proposed project folder structure for `helmgraph`. This structure reflects the chosen monolithic
Go application architecture within a monorepo, adhering to Go's standard project layout conventions and ensuring clear
separation of concerns.

```plaintext
helmgraph/
├── cmd/
│   └── helmgraph/            # Main application entry point
│       └── main.go           # CLI initialization and command handling
├── internal/
│   ├── manifest/             # Logic for generating Kubernetes manifests
│   │   └── generator.go      # Handles 'helm template' execution
│   ├── parser/               # Logic for parsing Kubernetes manifests
│   │   ├── parser.go         # Core parsing logic
│   │   └── models.go         # Go structs for Kubernetes resource representation
│   ├── relations/            # Logic for identifying relationships between resources
│   │   └── identifier.go     # Contains relationship identification rules
│   ├── cypher/               # Logic for generating Cypher statements
│   │   ├── generator.go      # Converts Go structs to Cypher strings
│   │   └── linter.go         # Cypher syntax validation
│   └── app/                  # Core application logic, orchestrating components
│       └── app.go            # Main application orchestrator
├── pkg/                      # Reusable packages (if any, currently none planned for MVP)
├── test/
│   ├── unit/                 # Unit tests for internal packages
│   ├── integration/          # Integration tests for component interactions
│   └── e2e/                  # End-to-end tests for the CLI tool
├── docs/                     # Project documentation (PRD, Architecture, etc.)
│   ├── prd.md
│   └── architecture.md
├── scripts/                  # Build, test, and utility scripts
├── .github/
│   └── workflows/            # GitHub Actions CI/CD workflows
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
├── LICENSE
└── README.md
```

---

**Rationale for Source Tree:**
This source tree structure follows standard Go project layout recommendations, particularly the use of `cmd/` for main
executables and `internal/` for application-specific private packages. This promotes modularity and prevents accidental
import of internal packages by external projects. The separation of `manifest`, `parser`, `relations`, and `cypher`
within `internal/` directly maps to the logical components defined earlier, ensuring a clear separation of concerns. The
`test/` directory is structured to accommodate unit, integration, and end-to-end tests, aligning with the PRD's testing
requirements. This layout is designed to be maintainable, scalable for future features, and idiomatic for Go
development.
