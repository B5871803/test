// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/B5871803/app/ent/examinationroom"
	"github.com/B5871803/app/ent/nurse"
	"github.com/B5871803/app/ent/operative"
	"github.com/B5871803/app/ent/operativerecord"
	"github.com/B5871803/app/ent/tool"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// OperativerecordCreate is the builder for creating a Operativerecord entity.
type OperativerecordCreate struct {
	config
	mutation *OperativerecordMutation
	hooks    []Hook
}

// SetOperativeTime sets the OperativeTime field.
func (oc *OperativerecordCreate) SetOperativeTime(t time.Time) *OperativerecordCreate {
	oc.mutation.SetOperativeTime(t)
	return oc
}

// SetNillableOperativeTime sets the OperativeTime field if the given value is not nil.
func (oc *OperativerecordCreate) SetNillableOperativeTime(t *time.Time) *OperativerecordCreate {
	if t != nil {
		oc.SetOperativeTime(*t)
	}
	return oc
}

// SetExaminationroomID sets the examinationroom edge to Examinationroom by id.
func (oc *OperativerecordCreate) SetExaminationroomID(id int) *OperativerecordCreate {
	oc.mutation.SetExaminationroomID(id)
	return oc
}

// SetNillableExaminationroomID sets the examinationroom edge to Examinationroom by id if the given value is not nil.
func (oc *OperativerecordCreate) SetNillableExaminationroomID(id *int) *OperativerecordCreate {
	if id != nil {
		oc = oc.SetExaminationroomID(*id)
	}
	return oc
}

// SetExaminationroom sets the examinationroom edge to Examinationroom.
func (oc *OperativerecordCreate) SetExaminationroom(e *Examinationroom) *OperativerecordCreate {
	return oc.SetExaminationroomID(e.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (oc *OperativerecordCreate) SetNurseID(id int) *OperativerecordCreate {
	oc.mutation.SetNurseID(id)
	return oc
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (oc *OperativerecordCreate) SetNillableNurseID(id *int) *OperativerecordCreate {
	if id != nil {
		oc = oc.SetNurseID(*id)
	}
	return oc
}

// SetNurse sets the nurse edge to Nurse.
func (oc *OperativerecordCreate) SetNurse(n *Nurse) *OperativerecordCreate {
	return oc.SetNurseID(n.ID)
}

// SetOperativeID sets the operative edge to Operative by id.
func (oc *OperativerecordCreate) SetOperativeID(id int) *OperativerecordCreate {
	oc.mutation.SetOperativeID(id)
	return oc
}

// SetNillableOperativeID sets the operative edge to Operative by id if the given value is not nil.
func (oc *OperativerecordCreate) SetNillableOperativeID(id *int) *OperativerecordCreate {
	if id != nil {
		oc = oc.SetOperativeID(*id)
	}
	return oc
}

// SetOperative sets the operative edge to Operative.
func (oc *OperativerecordCreate) SetOperative(o *Operative) *OperativerecordCreate {
	return oc.SetOperativeID(o.ID)
}

// SetToolID sets the tool edge to Tool by id.
func (oc *OperativerecordCreate) SetToolID(id int) *OperativerecordCreate {
	oc.mutation.SetToolID(id)
	return oc
}

// SetNillableToolID sets the tool edge to Tool by id if the given value is not nil.
func (oc *OperativerecordCreate) SetNillableToolID(id *int) *OperativerecordCreate {
	if id != nil {
		oc = oc.SetToolID(*id)
	}
	return oc
}

// SetTool sets the tool edge to Tool.
func (oc *OperativerecordCreate) SetTool(t *Tool) *OperativerecordCreate {
	return oc.SetToolID(t.ID)
}

// Mutation returns the OperativerecordMutation object of the builder.
func (oc *OperativerecordCreate) Mutation() *OperativerecordMutation {
	return oc.mutation
}

// Save creates the Operativerecord in the database.
func (oc *OperativerecordCreate) Save(ctx context.Context) (*Operativerecord, error) {
	if _, ok := oc.mutation.OperativeTime(); !ok {
		v := operativerecord.DefaultOperativeTime()
		oc.mutation.SetOperativeTime(v)
	}
	var (
		err  error
		node *Operativerecord
	)
	if len(oc.hooks) == 0 {
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperativerecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OperativerecordCreate) SaveX(ctx context.Context) *Operativerecord {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oc *OperativerecordCreate) sqlSave(ctx context.Context) (*Operativerecord, error) {
	o, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	o.ID = int(id)
	return o, nil
}

func (oc *OperativerecordCreate) createSpec() (*Operativerecord, *sqlgraph.CreateSpec) {
	var (
		o     = &Operativerecord{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: operativerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: operativerecord.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.OperativeTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: operativerecord.FieldOperativeTime,
		})
		o.OperativeTime = value
	}
	if nodes := oc.mutation.ExaminationroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operativerecord.ExaminationroomTable,
			Columns: []string{operativerecord.ExaminationroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: examinationroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operativerecord.NurseTable,
			Columns: []string{operativerecord.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OperativeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operativerecord.OperativeTable,
			Columns: []string{operativerecord.OperativeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operative.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ToolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operativerecord.ToolTable,
			Columns: []string{operativerecord.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tool.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return o, _spec
}
