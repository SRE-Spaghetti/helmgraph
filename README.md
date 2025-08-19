# Helm Graph

## Initial Prompt

This project is for creating the HelmGraph application through the initial prompt:

```
*plan I want to develop a new project from scratch. I want to design a service that will allow me to examine a helm
chart and convert it in to cypher query commands so that is can be loaded in to Neo4j graph database. the service should
accept the name of the helm chart and use the "helm template" command to produce a manifest as an intermediate step. The
through this manifest "kind" in the helm chart should be taken as a node type and each instance of this kind should be
treated as a node (identified by label.name and label.namespace).
```

## The BMAD Method

This was developed using [the BMAD Method](https://github.com/bmad-code-org/BMAD-METHOD/blob/main/docs/user-guide.md) 

1. Ran the [Planning Workflow](https://github.com/bmad-code-org/BMAD-METHOD/blob/main/docs/user-guide.md#the-planning-workflow-web-ui-or-powerful-ide-agents)
   with [Google Gemini CLI](https://github.com/google-gemini/gemini-cli) with `google-2.5-flash` model (instead of the
   web interface as suggested by BMAD) to create:
   1. The [Product brief](docs/brief.md)
   2. The [PRD docuemnt](docs/prd.md)
1. Then hand edited the PRD document (leaving the old one as [a backup](docs/prd-backup.md)).
   1. Please compare the two to see the edits
1. I will continue to refine the plan some more
1. Then I will tackle the [Execute workflow](https://github.com/bmad-code-org/BMAD-METHOD/blob/main/docs/user-guide.md#the-core-development-cycle-ide)
   with `gemini-2.5-pro` model.


### Gemini CLI notes:

- To start Gemini CLI with the flash model run it as `gemini --model gemini-2.5-flash`
  - I ran out of tokens (`GenerateRequestsPerDayPerProjectPerModel-FreeTier > 100`) when using `gemini-2.5-pro`
    for the planning stages. The rate limits for this are about 1000 with flash model. The pro model is more suited to
    code development
- I used [Google AI Studio](https://aistudio.google.com) to connect Gemini CLI with a Google Cloud Platform **Project**.
  - In this way I was able to have billing to the GCP project rather than a [regular subscription](https://developers.google.com/program/plans-and-pricing).
