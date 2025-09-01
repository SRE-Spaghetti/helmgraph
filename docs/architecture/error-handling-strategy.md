# Error Handling Strategy

This section defines the comprehensive error handling approach for `helmgraph`. Given that `helmgraph` is a command-line
tool, the strategy will focus on providing clear, actionable error messages to the user, robust internal error
management, and effective logging for debugging.

---

**General Approach**

-   **Error Model:** Go's idiomatic error handling will be used, returning `error` as the last return value from
    functions. Custom error types will be defined for specific, recoverable, or categorizable errors.
-   **Exception Hierarchy:** Go does not have exceptions in the traditional sense. Errors will be propagated up the call
    stack using `return err` until they can be handled appropriately (e.g., logged, presented to the user).
-   **Error Propagation:** Errors will be propagated with context using `fmt.Errorf` and `errors.Wrap` (if using a
    wrapping library like `pkg/errors` or Go 1.13+ `errors.Is`/`errors.As` for error chaining) to provide a clear trail
    of where the error originated and what caused it.

**Rationale:** Adhering to Go's native error handling patterns ensures consistency and readability. Propagating errors
with context is crucial for debugging a CLI tool, allowing developers to understand the full chain of events leading to
an issue.

---

**Logging Standards**

-   **Library:** `log` (Go standard library) or `logrus` / `zap` (for more advanced features)
-   **Format:** Structured logging (e.g., JSON) for machine readability, plain text for console output.
-   **Levels:** `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`.
-   **Required Context:**
    -   Correlation ID: N/A (single-process CLI tool, no distributed tracing needed for MVP)
    -   Service Context: Component/module name (e.g., `[parser]`, `[cypher-generator]`)
    -   User Context: Input parameters (e.g., chart path, release name - sanitized to avoid sensitive data)

**Rationale:** While `helmgraph` is a CLI tool, structured logging is beneficial for debugging, especially if logs are
ever redirected to a file or a log aggregation system. Clear logging levels help in filtering output based on
verbosity. Including component context helps pinpoint the source of issues.

---

**Error Handling Patterns**

#### External CLI Errors (e.g., `helm` CLI)

-   **Retry Policy:** No automatic retries for `helm` CLI invocation. Fail fast and report the error.
-   **Circuit Breaker:** Not applicable for a single, synchronous CLI invocation.
-   **Timeout Configuration:** A reasonable timeout will be applied to the `helm template` command execution to prevent
    indefinite hangs (e.g., 5 minutes).
-   **Error Translation:** `helm` CLI errors will be captured (stderr) and translated into user-friendly messages,
    indicating issues like invalid chart paths, missing `helm` executable, or template rendering failures.

**Rationale:** For external CLI dependencies, direct error reporting is preferred. Timeouts prevent unresponsive
behavior. Translating raw CLI errors into understandable messages is critical for a good user experience.

#### Business Logic Errors (e.g., Parsing, Relationship Identification)

-   **Custom Exceptions:** Custom error types will be defined for specific business logic failures, such as
    `ErrInvalidManifest`, `ErrResourceNotFound`, `ErrRelationshipNotIdentified`.
-   **User-Facing Errors:** Errors presented to the user will be concise, clear, and actionable, guiding them on how to
    resolve the issue (e.g., "Error: Invalid Helm chart path. Please check the path and try again.").
-   **Error Codes:** Simple, internal error codes or enumerated error types can be used for programmatic identification
    of error categories, but not exposed directly to the user.

**Rationale:** Custom error types allow for more precise error handling and testing. User-facing errors must be helpful,
not just technical dumps.

#### Data Consistency (Internal)

-   **Transaction Strategy:** Not applicable, as `helmgraph` does not manage its own persistent data or transactions.
    All operations are in-memory and atomic for a single execution.
-   **Compensation Logic:** Not applicable. If an error occurs during processing, the tool will stop and report the
    error; no partial state needs to be rolled back.
-   **Idempotency:** The `helmgraph` execution itself is idempotent in terms of its output for a given input. Running it
    multiple times with the same input will produce the same `.cypher` file.

**Rationale:** This section clarifies that traditional database-centric data consistency patterns are not relevant for
`helmgraph` due to its stateless, CLI nature.

---

**Rationale for Error Handling Strategy:**
This strategy is tailored for a Go-based command-line tool. It emphasizes Go's idiomatic error handling, clear user
feedback, and robust internal logging. Specific patterns are defined for interacting with external CLIs and handling
internal business logic errors. The non-applicability of certain patterns (like transactions or retries) is explicitly
stated, aligning with the tool's design. This ensures that `helmgraph` will be resilient, debuggable, and user-friendly
when issues arise.
