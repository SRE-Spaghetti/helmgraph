package cypher

import (
	"fmt"
	"helmgraph/internal/parser"
	"helmgraph/internal/relations"
	"strings"
)

// Generate generates a Cypher script from a slice of resources and relationships.
func Generate(resources []*parser.Resource, relationships []*relations.Relationship) string {
	var sb strings.Builder

	// Generate constraints
	kinds := make(map[string]bool)
	for _, r := range resources {
		if _, ok := kinds[r.Kind]; !ok {
			sb.WriteString(fmt.Sprintf("CREATE CONSTRAINT IF NOT EXISTS FOR (n:%s) REQUIRE (n.name, n.namespace) IS UNIQUE;\n", r.Kind))
			kinds[r.Kind] = true
		}
	}

	// Generate nodes
	for _, r := range resources {
		sb.WriteString(fmt.Sprintf("MERGE (:%s {name: '%s', namespace: '%s', kind: '%s'});\n", r.Kind, r.Metadata.Name, r.Metadata.Namespace, r.Kind))
	}

	// Generate relationships
	for _, rel := range relationships {
		sb.WriteString(fmt.Sprintf("MATCH (a:%s {name: '%s'}), (b:%s {name: '%s'}) MERGE (a)-[:%s]->(b);\n", rel.Source.Kind, rel.Source.Metadata.Name, rel.Target.Kind, rel.Target.Metadata.Name, rel.Type))
	}

	return sb.String()
}
