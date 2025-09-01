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
1. Then as the `*agent architect` I created the [architecture.md](docs/architecture.md)

1. For the [Execute Workflow](https://github.com/bmad-code-org/BMAD-METHOD/blob/main/docs/user-guide.md#the-core-development-cycle-ide)
   I ran it with [Google Gemini CLI](https://github.com/google-gemini/gemini-cli) with `google-2.5-pro` model (instead of the
   IDE as suggested by BMAD)
1. First I sharded the PRD and Architecture documents (strictly this is a part of the Planning Workflow but neds the bigger model)
1. Then I used the `*agent sm` (Scrum Master) to "create next story" iterating through the 5 in Epic 1 and the 2 in Epic 2
   1. I stopped after each story and committed the code for each
   2. When ever a test failed I stopped the code generation and ran the test manually to diagnose the problem
1. When all stories were finished I ran the code to test it on real Helm charts and made some adjustments
1. Finally I got Gemini to create a Makefile


### Gemini CLI notes:

- To start Gemini CLI with the flash model run it as `gemini --model gemini-2.5-flash`
  - I ran out of tokens (`GenerateRequestsPerDayPerProjectPerModel-FreeTier > 100`) when using `gemini-2.5-pro`
    for the planning stages. The rate limits for this are about 1000 with flash model. The pro model is more suited to
    code development
- I used [Google AI Studio](https://aistudio.google.com) to connect Gemini CLI with a Google Cloud Platform **Project**.
  - In this way I was able to have billing to the GCP project rather than a [regular subscription](https://developers.google.com/program/plans-and-pricing).
