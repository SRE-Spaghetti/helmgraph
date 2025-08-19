# Project Brief: HelmGraph

## Executive Summary

This document outlines the project brief for **HelmGraph**, a service designed to convert Helm chart templates into Cypher queries for import into a Neo4j graph database. The primary problem it solves is the difficulty in visualizing and understanding the complex relationships between Kubernetes resources defined in Helm charts. The target users are DevOps engineers, Kubernetes administrators, and SREs. The key value proposition is to provide a clear, queryable graph-based representation of Helm chart deployments.

## Problem Statement

Kubernetes configurations, especially when managed through Helm, can become incredibly complex. A single Helm chart can define dozens of interconnected resources (Deployments, Services, ConfigMaps, etc.). Understanding the dependencies, relationships, and potential impact of changes within these charts is a significant challenge. Existing tools primarily offer text-based or tabular views of these resources, which fail to capture the rich, networked nature of a Kubernetes application. This lack of visibility makes debugging, security analysis, and general cluster management more difficult and time-consuming.

## Proposed Solution

The proposed solution is **HelmGraph**, a service that ingests a Helm chart and produces a set of Cypher query commands. The service will orchestrate the following workflow:

1.  Accept the name of a Helm chart as input.
2.  Use the `helm template` command to generate a consolidated Kubernetes manifest file.
3.  Parse this manifest, treating each Kubernetes resource (identified by its `kind`, `metadata.name`, and `metadata.namespace`) as a node in the graph.
4.  Analyze relationships between resources (e.g., a Deployment using a ConfigMap, a Service targeting a Deployment) and define them as relationships (edges) in the graph.
5.  Generate a script of Cypher `CREATE` and `MERGE` commands that represent these nodes and relationships.
6.  The output will be a `.cypher` file that can be directly executed against a Neo4j database to build the graph representation of the Helm chart.

This approach provides a powerful, queryable, and visual way to explore and analyze Helm chart deployments, succeeding where other tools fall short by embracing the networked nature of Kubernetes applications.

## Target Users

**Primary User Segment: DevOps Engineers / SREs**

*   **Profile:** These users are responsible for deploying, managing, and troubleshooting applications on Kubernetes. They are highly technical and comfortable with both Helm and the command line.
*   **Behaviors:** They spend a significant amount of time in their terminal, inspecting Kubernetes resources (`kubectl get`, `describe`), reading YAML files, and analyzing logs. They are often tasked with understanding why a deployment is failing or how a proposed change might impact the system.
*   **Needs & Pains:** They need to quickly understand the full scope of an application defined in a Helm chart. Their primary pain point is the time it takes to manually trace connections between different resources, which is error-prone and tedious.
*   **Goals:** Their main goal is to ensure the reliability, scalability, and security of the applications they manage. They want to reduce the mean time to resolution (MTTR) for incidents and confidently assess the impact of changes.

## Goals & Success Metrics

### Business Objectives
- **Improve Operational Efficiency:** Reduce the time DevOps and SRE teams spend manually analyzing Helm chart dependencies. Success will be measured by the time saved per analysis task, estimated through user feedback.

### User Success Metrics
- **Enhanced Understanding:** Users can successfully answer complex dependency questions about their Helm charts within minutes of using the tool (e.g., "What services would be affected if I change this ConfigMap?").
- **Increased Confidence:** Users feel more confident in deploying changes because they can clearly visualize the potential impact.

### Key Performance Indicators (KPIs)
- **Adoption Rate:** The number of unique users or teams actively using the service within the first 3 months of launch.
- **Processing Time:** The average time it takes for the service to process a standard Helm chart and generate the Cypher output. Target: < 60 seconds.

## MVP Scope

### Core Features (Must Have)
- **Helm Chart Processing:** Accept a local Helm chart path as input and execute `helm template` to generate a manifest.
- **Node Generation:** Parse the manifest to create Cypher nodes for core Kubernetes kinds: `Deployment`, `Service`, `ConfigMap`, `Secret`, and `Ingress`. Nodes should be identified by their `kind`, `name`, and `namespace`.
- **Relationship Generation:** Create Cypher relationships for common connections, such as a `Service` selecting a `Deployment`, or a `Deployment` mounting a `Secret`.
- **Cypher Script Output:** Produce a valid `.cypher` file containing all the generated `CREATE` and `MERGE` commands.

### Out of Scope for MVP
- Direct interaction with a live Kubernetes cluster.
- Support for remote Helm repositories or chart dependencies.
- A graphical user interface (GUI) for visualization. The tool's output is the script.
- Support for all possible Kubernetes resource types (we will focus on the most common ones first).

### MVP Success Criteria
The MVP is successful if a user can process a standard Helm chart (like `bitnami/wordpress`), load the output into Neo4j, and see an accurate graph of the application's key resources and their connections.

## Post-MVP Vision

### Phase 2 Features
- Support for remote Helm repositories (fetching charts directly from URLs or OCI registries).
- Expanded support for more Kubernetes resource types (e.g., `StatefulSet`, `DaemonSet`, `PersistentVolumeClaim`).
- Automatic analysis of Helm chart dependencies (`Chart.yaml`).

### Long-term Vision
- A web-based UI for direct visualization and interactive exploration of the generated graph.
- Integration with live Kubernetes clusters to compare the desired state (from Helm) with the actual state.
- A library of pre-built Cypher queries for common analysis tasks (e.g., "find all publicly exposed services," "show all resources with a specific label").

### Expansion Opportunities
- Support for other Kubernetes configuration tools like Kustomize or raw YAML manifests.
- Integration with CI/CD pipelines to automatically generate and store graph snapshots on every deployment, creating a historical record of architectural changes.

## Technical Considerations

### Platform Requirements
- **Target Platforms:** The service should be a command-line tool that can run on Linux, macOS, and Windows (via WSL).
- **Performance Requirements:** Should process a Helm chart with ~50 resources in under 30 seconds.

### Technology Preferences
- **Backend:** Go or Python are preferred for their strong CLI development ecosystem and YAML parsing libraries.
- **Database:** The output is for Neo4j, but the service itself does not require a database.
- **Hosting/Infrastructure:** As a CLI tool, it will be distributed as a binary. No hosting is required for the MVP.

### Architecture Considerations
- **Repository Structure:** A standard Go or Python project structure.
- **Service Architecture:** A single, stateless command-line application.
- **Integration Requirements:** Requires the `helm` CLI to be installed and available in the user's PATH.
- **Security/Compliance:** The tool only reads local files and does not transmit any data, so the initial security footprint is minimal.

## Constraints & Assumptions

### Constraints
- **Timeline:** The MVP should be developed within 4-6 weeks.
- **Resources:** The project will be developed by a single engineer.
- **Technical:** The tool's functionality is dependent on the output of the `helm template` command. Any changes or limitations in that command will directly affect HelmGraph.

### Key Assumptions
- Users have `helm` v3 installed and correctly configured on their system.
- Users have a working knowledge of Cypher and access to a Neo4j database instance to run the output script.
- The Kubernetes manifests generated by `helm template` will be well-formed YAML.
- The primary value for users is in the generated Cypher script, not in a visual UI (for the MVP).

## Risks & Open Questions

### Key Risks
- **Parsing Complexity:** Kubernetes YAML can be incredibly complex and varied. There's a risk that parsing all the different ways resources can be defined and linked will be more difficult than anticipated. (Impact: High, Probability: Medium)
- **Limited Adoption:** The target audience is niche. There's a risk that not enough users will find the tool useful to justify further development post-MVP. (Impact: Medium, Probability: Medium)

### Open Questions
- What are the most critical relationships between Kubernetes objects that users want to visualize? (e.g., network policies, service accounts, RBAC roles)
- How should the service handle custom resource definitions (CRDs)? Should they be ignored, or should there be a way to define custom mapping rules?

### Areas Needing Further Research
- A survey of existing open-source tools that perform similar Kubernetes visualization or analysis to identify best practices and potential gaps.
- Investigation into robust Go or Python libraries for parsing multi-document YAML streams and handling Kubernetes-specific object structures.

## Next Steps

This Project Brief provides the full context for HelmGraph. The next step is to hand this off to the Product Manager to create the PRD.
