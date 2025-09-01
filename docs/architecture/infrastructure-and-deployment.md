# Infrastructure and Deployment

This section defines the deployment architecture and practices for `helmgraph`. As a command-line tool distributed as a
binary and Docker image, its deployment strategy is focused on distribution and local execution rather than continuous
hosting.

---

**Infrastructure as Code**

-   **Tool:** N/A
-   **Location:** N/A
-   **Approach:** N/A

**Rationale:** Since `helmgraph` is a standalone CLI tool and does not require its own hosted infrastructure (NFR5),
Infrastructure as Code (IaC) tools are not directly applicable for managing its operational environment. IaC would be
relevant for the Neo4j database that consumes `helmgraph`'s output, but that is external to this project's scope.

---

**Deployment Strategy**

-   **Strategy:** Binary Distribution / Docker Image Distribution
-   **CI/CD Platform:** GitHub Actions (Recommended)
-   **Pipeline Configuration:** `.github/workflows/build-and-release.yaml` (Proposed)

**Rationale:** The primary deployment strategy aligns with NFR6, distributing `helmgraph` as a standalone binary for
various operating systems (NFR2) and as a Docker image. GitHub Actions is a recommended CI/CD platform due to its
integration with GitHub repositories, ease of use, and capabilities for cross-platform compilation and Docker image
building. A dedicated workflow will automate the build, test, and release process.

---

**Environments**

-   **Development:** Local developer machines - For active coding, testing, and debugging.
-   **CI/CD:** GitHub Actions Runners - For automated builds, tests, and release artifact generation.
-   **User Machines:** End-user environments (Linux, macOS, Windows/WSL) - Where the binary or Docker image is
    downloaded and executed.

**Rationale:** Given `helmgraph`'s nature as a CLI tool, traditional "environments" like staging or production servers
are not applicable for its own hosting. Instead, environments refer to where the tool is developed, built, and
ultimately consumed by users.

---

**Environment Promotion Flow**

```text
Developer Machine (Code)
    ↓
GitHub (Push)
    ↓
GitHub Actions (Build, Test, Lint)
    ↓ (On successful build/test)
GitHub Releases (Binary & Docker Image Artifacts)
    ↓
User Machines (Download & Execute)
```

**Rationale:** This flow outlines a simple, direct path from code development to user consumption. Changes are pushed to
GitHub, automated CI/CD pipelines build and test the application, and upon success, release artifacts (binaries, Docker
images) are published to GitHub Releases, from where users can download and run the tool.

---

**Rollback Strategy**

-   **Primary Method:** Download previous release version
-   **Trigger Conditions:** Critical bug discovered in new release, unexpected behavior, performance degradation.
-   **Recovery Time Objective:** Minutes (dependent on user's download speed and execution time)

**Rationale:** For a CLI tool, rollback is straightforward: users simply download and replace the problematic version
with a previous stable release from GitHub Releases. This is a quick and effective method, as there is no complex
server-side infrastructure to manage.

---

**Rationale for Infrastructure and Deployment:**
This section details how `helmgraph` will be built, released, and consumed, aligning with its identity as a standalone
CLI tool. The focus is on efficient distribution and ease of use for the end-user, rather than complex server
deployments. GitHub Actions is proposed as the CI/CD backbone to automate the release process, ensuring consistent and
reliable delivery of the tool.
