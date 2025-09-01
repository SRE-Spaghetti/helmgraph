# Coding Standards

These standards are **MANDATORY** for AI agents (and human developers). We need to define only the critical rules needed
to prevent bad code. It's important to understand that:

1.  This section directly controls AI developer behavior.
2.  We should keep it minimal – assume AI knows general best practices.
3.  Focus on project-specific conventions and "gotchas."
4.  Overly detailed standards bloat context and slow development.
5.  These standards will be extracted to a separate file for the dev agent's direct use.

For each standard, I will seek your explicit confirmation that it's necessary.

---

**Core Standards**

-   **Languages & Runtimes:** Go 1.24.4
    -   **Detail:** We are standardizing on Go version 1.24.4. This version provides the most up-to-date language
        features, performance improvements, and security patches. It aligns with the PRD's explicit requirement for Go
        (NFR4) and ensures we leverage the latest advancements in the Go ecosystem for building a robust and efficient
        CLI tool. Pinning to a specific major.minor version (e.g., 1.24.4) helps ensure build reproducibility across
        different development environments and CI/CD pipelines.

-   **Style & Linting:**
    -   `gofmt`: **Mandatory** for code formatting.
        -   **Detail:** `gofmt` is the official Go formatting tool. It automatically formats Go source code according to
            the standard Go style. Its use is non-negotiable to ensure absolute consistency in code style across the
            entire project, regardless of who (or what AI agent) writes the code. This eliminates style debates and
            makes code reviews more efficient.
    -   `golint`: For stylistic checks.
        -   **Detail:** `golint` checks for stylistic errors in Go source code, such as naming conventions (e.g.,
            exported names should have comments), unexported struct fields, and other common Go idioms. While some of
            its checks might be superseded by `staticcheck`, it still provides valuable feedback on code readability
            and adherence to Go's idiomatic style.
    -   `staticcheck`: Comprehensive static analysis for Go code.
        -   **Detail:** `staticcheck` is a powerful static analysis tool that detects various kinds of bugs and
            suspicious constructs in Go code. It includes checks for unused code, inefficient operations, potential
            panics, and more. This tool is critical for catching subtle bugs early in the development cycle and
            ensuring high code quality and performance.
    -   `go vet`: Standard Go tool for suspicious constructs.
        -   **Detail:** `go vet` is a built-in Go tool that examines Go source code and reports suspicious constructs,
            suchs as unreachable code, incorrect format string usage, and common concurrency issues. It acts as a first
            line of defense for identifying potential runtime problems before they manifest.

-   **Test Organization:**
    -   Unit tests: `_test.go` files in the same package as the code they test.
        -   **Detail:** This is the standard Go convention. For example, `internal/parser/parser.go` would have its unit
            tests in `internal/parser/parser_test.go`. This co-location makes it easy to find tests for a given piece
            of code and ensures that tests can access unexported functions/types within the same package.
    -   Integration tests: `_test.go` files in a separate `test/integration` directory, or within the package but
        clearly separated (e.g., `integration_test.go`).
        -   **Detail:** Integration tests verify the interaction between multiple components or external dependencies
            (like the `helm` CLI). If they are tightly coupled to a specific package, they can reside in that package
            (e.g., `internal/manifest/generator_integration_test.go`). Otherwise, for broader integration scenarios,
            they will be placed in `test/integration/`. This separation helps distinguish between isolated unit tests
            and tests that require more setup or external resources.
    -   E2E tests: Located in `test/e2e`.
        -   **Detail:** End-to-end tests validate the entire `helmgraph` CLI tool from user invocation to final output.
            These tests will simulate real user scenarios and verify the complete workflow. Placing them in `test/e2e/`
            provides a clear, high-level testing suite that confirms the overall functionality of the built binary or
            Docker image.

---

**Rationale for Core Standards (Expanded):**
This expanded rationale provides a deeper understanding of *why* each standard is chosen and *how* it contributes to the
overall quality and maintainability of the `helmgraph` codebase. By detailing the purpose of each linter and the
rationale behind test organization, we ensure that both AI agents and human developers have a clear and consistent
understanding of the project's coding expectations. These standards are designed to enforce Go idioms, catch common
errors, and streamline the development and review process.

### Naming Conventions

| Element | Convention | Example |
|---|---|---|

### Critical Rules

### Language-Specific Guidelines

This section is intended for **highly specific Go rules** that are critical for preventing AI (or human) developers from
making common or project-specific mistakes that are not covered by general Go best practices or the linters mentioned
above.

**Rationale:**
Most Go projects do not need this section. Go's strong conventions and the effectiveness of tools like `gofmt`,
`golint`, and `staticcheck` cover the vast majority of coding standards. Adding rules here should be done **only if
absolutely critical** for `helmgraph` to prevent specific, recurring issues or to enforce a very particular project-level
idiom that deviates from common Go patterns. Overly detailed rules here can add unnecessary context and slow down
development.

If we do add rules here, they should be concise and actionable.

Here's the template for this section:

#### {{language_name}} Specifics

-   **{{rule_topic}}:** {{rule_detail}}

**Example (if we were to add one, but generally avoid unless critical):**
-   **Error Handling:** Always wrap errors with `fmt.Errorf("...: %w", err)` when propagating, to preserve the original
    error chain for debugging.

---

**Rationale for Language-Specific Guidelines (Expanded):**
This section is a last resort for very specific, critical Go-related rules. The emphasis is on avoiding redundancy with
existing Go tooling and general best practices. Its purpose is to address unique `helmgraph` project requirements or to
explicitly prevent known pitfalls that the AI might otherwise introduce. We should only populate this if there's a clear,
demonstrated need.
