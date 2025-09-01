# Introduction

This document outlines the overall project architecture for helmgraph, including backend systems, shared services, and
non-UI specific concerns. Its primary goal is to serve as the guiding architectural blueprint for AI-driven
development, ensuring consistency and adherence to chosen patterns and technologies.

**Relationship to Frontend Architecture:**
If the project includes a significant user interface, a separate Frontend Architecture Document will detail the
frontend-specific design and MUST be used in conjunction with this document. Core technology stack choices documented
herein (see "Tech Stack") are definitive for the entire project, including any frontend components.

### Starter Template or Existing Project

Before proceeding further with architecture design, check if the project is based on a starter template or existing
codebase:

1. Review the PRD and brainstorming brief for any mentions of:
- Starter templates (e.g., Create React App, Next.js, Vue CLI, Angular CLI, etc.)
- Existing projects or codebases being used as a foundation
- Boilerplate projects or scaffolding tools
- Previous projects to be cloned or adapted

2. If a starter template or existing project is mentioned:
- Ask the user to provide access via one of these methods:
  - Link to the starter template documentation
  - Upload/attach the project files (for small projects)
  - Share a link to the project repository (GitHub, GitLab, etc.)
- Analyze the starter/existing project to understand:
  - Pre-configured technology stack and versions
  - Project structure and organization patterns
  - Built-in scripts and tooling
  - Existing architectural patterns and conventions
  - Any limitations or constraints imposed by the starter
- Use this analysis to inform and align your architecture decisions

3. If no starter template is mentioned but this is a greenfield project:
- Suggest appropriate starter templates based on the tech stack preferences
- Explain the benefits (faster setup, best practices, community support)
- Let the user decide whether to use one

4. If the user confirms no starter template will be used:
- Proceed with architecture design from scratch
- Note that manual setup will be required for all tooling and configuration

Document the decision here before proceeding with the architecture design. If none, just say N/A
N/A

### Change Log

| Date | Version | Description | Author |
|---|---|---|---|