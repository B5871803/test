// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/B5871803/app/ent/operativerecord"
	"github.com/B5871803/app/ent/tool"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ToolCreate is the builder for creating a Tool entity.
type ToolCreate struct {
	config
	mutation *ToolMutation
	hooks    []Hook
}

// SetToolName sets the toolName field.
func (tc *ToolCreate) SetToolName(s string) *ToolCreate {
	tc.mutation.SetToolName(s)
	return tc
}

// SetAmount sets the amount field.
func (tc *ToolCreate) SetAmount(s string) *ToolCreate {
	tc.mutation.SetAmount(s)
	return tc
}

// AddFromtoolIDs adds the fromtool edge to Operativerecord by ids.
func (tc *ToolCreate) AddFromtoolIDs(ids ...int) *ToolCreate {
	tc.mutation.AddFromtoolIDs(ids...)
	return tc
}

// AddFromtool adds the fromtool edges to Operativerecord.
func (tc *ToolCreate) AddFromtool(o ...*Operativerecord) *ToolCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return tc.AddFromtoolIDs(ids...)
}

// Mutation returns the ToolMutation object of the builder.
func (tc *ToolCreate) Mutation() *ToolMutation {
	return tc.mutation
}

// Save creates the Tool in the database.
func (tc *ToolCreate) Save(ctx context.Context) (*Tool, error) {
	if _, ok := tc.mutation.ToolName(); !ok {
		return nil, &ValidationError{Name: "toolName", err: errors.New("ent: missing required field \"toolName\"")}
	}
	if v, ok := tc.mutation.ToolName(); ok {
		if err := tool.ToolNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "toolName", err: fmt.Errorf("ent: validator failed for field \"toolName\": %w", err)}
		}
	}
	if _, ok := tc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if v, ok := tc.mutation.Amount(); ok {
		if err := tool.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	var (
		err  error
		node *Tool
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ToolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ToolCreate) SaveX(ctx context.Context) *Tool {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *ToolCreate) sqlSave(ctx context.Context) (*Tool, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *ToolCreate) createSpec() (*Tool, *sqlgraph.CreateSpec) {
	var (
		t     = &Tool{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tool.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.ToolName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tool.FieldToolName,
		})
		t.ToolName = value
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tool.FieldAmount,
		})
		t.Amount = value
	}
	if nodes := tc.mutation.FromtoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tool.FromtoolTable,
			Columns: []string{tool.FromtoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}