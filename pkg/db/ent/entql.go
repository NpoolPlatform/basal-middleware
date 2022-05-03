// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/service-template/pkg/db/ent/empty"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   empty.Table,
			Columns: empty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: empty.FieldID,
			},
		},
		Type:   "Empty",
		Fields: map[string]*sqlgraph.FieldSpec{},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (eq *EmptyQuery) addPredicate(pred func(s *sql.Selector)) {
	eq.predicates = append(eq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the EmptyQuery builder.
func (eq *EmptyQuery) Filter() *EmptyFilter {
	return &EmptyFilter{eq}
}

// addPredicate implements the predicateAdder interface.
func (m *EmptyMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the EmptyMutation builder.
func (m *EmptyMutation) Filter() *EmptyFilter {
	return &EmptyFilter{m}
}

// EmptyFilter provides a generic filtering capability at runtime for EmptyQuery.
type EmptyFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *EmptyFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *EmptyFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(empty.FieldID))
}
