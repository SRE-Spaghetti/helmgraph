package cypher

import (
	"helmgraph/internal/parser"
	"helmgraph/internal/relations"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	resources := []*parser.Resource{
		{
			Kind: "Service",
			Metadata: parser.Metadata{
				Name:      "my-service",
				Namespace: "default",
			},
		},
		{
			Kind: "Deployment",
			Metadata: parser.Metadata{
				Name:      "my-deployment",
				Namespace: "default",
			},
		},
	}

	relationships := []*relations.Relationship{
		{
			Source: resources[0],
			Target: resources[1],
			Type:   "SELECTS",
		},
	}

	script := Generate(resources, relationships)

	expectedConstraint1 := "CREATE CONSTRAINT IF NOT EXISTS FOR (n:Service) REQUIRE (n.name, n.namespace) IS UNIQUE;"
	expectedConstraint2 := "CREATE CONSTRAINT IF NOT EXISTS FOR (n:Deployment) REQUIRE (n.name, n.namespace) IS UNIQUE;"
	expectedNode1 := "MERGE (:Service {name: 'my-service', namespace: 'default', kind: 'Service'});"
	expectedNode2 := "MERGE (:Deployment {name: 'my-deployment', namespace: 'default', kind: 'Deployment'});"
	expectedRel := "MATCH (a:Service {name: 'my-service'}), (b:Deployment {name: 'my-deployment'}) MERGE (a)-[:SELECTS]->(b);"

	if !strings.Contains(script, expectedConstraint1) {
		t.Errorf("script does not contain expected constraint: %s", expectedConstraint1)
	}
	if !strings.Contains(script, expectedConstraint2) {
		t.Errorf("script does not contain expected constraint: %s", expectedConstraint2)
	}
	if !strings.Contains(script, expectedNode1) {
		t.Errorf("script does not contain expected node: %s", expectedNode1)
	}
	if !strings.Contains(script, expectedNode2) {
		t.Errorf("script does not contain expected node: %s", expectedNode2)
	}
	if !strings.Contains(script, expectedRel) {
		t.Errorf("script does not contain expected relationship: %s", expectedRel)
	}
}
