// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/B5871803/app/ent/examinationroom"
	"github.com/B5871803/app/ent/nurse"
	"github.com/B5871803/app/ent/operative"
	"github.com/B5871803/app/ent/operativerecord"
	"github.com/B5871803/app/ent/predicate"
	"github.com/B5871803/app/ent/tool"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// OperativerecordQuery is the builder for querying Operativerecord entities.
type OperativerecordQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Operativerecord
	// eager-loading edges.
	withExaminationroom *ExaminationroomQuery
	withNurse           *NurseQuery
	withOperative       *OperativeQuery
	withTool            *ToolQuery
	withFKs             bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (oq *OperativerecordQuery) Where(ps ...predicate.Operativerecord) *OperativerecordQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OperativerecordQuery) Limit(limit int) *OperativerecordQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OperativerecordQuery) Offset(offset int) *OperativerecordQuery {
	oq.offset = &offset
	return oq
}

// Order adds an order step to the query.
func (oq *OperativerecordQuery) Order(o ...OrderFunc) *OperativerecordQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryExaminationroom chains the current query on the examinationroom edge.
func (oq *OperativerecordQuery) QueryExaminationroom() *ExaminationroomQuery {
	query := &ExaminationroomQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operativerecord.Table, operativerecord.FieldID, oq.sqlQuery()),
			sqlgraph.To(examinationroom.Table, examinationroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operativerecord.ExaminationroomTable, operativerecord.ExaminationroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNurse chains the current query on the nurse edge.
func (oq *OperativerecordQuery) QueryNurse() *NurseQuery {
	query := &NurseQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operativerecord.Table, operativerecord.FieldID, oq.sqlQuery()),
			sqlgraph.To(nurse.Table, nurse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operativerecord.NurseTable, operativerecord.NurseColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOperative chains the current query on the operative edge.
func (oq *OperativerecordQuery) QueryOperative() *OperativeQuery {
	query := &OperativeQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operativerecord.Table, operativerecord.FieldID, oq.sqlQuery()),
			sqlgraph.To(operative.Table, operative.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operativerecord.OperativeTable, operativerecord.OperativeColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTool chains the current query on the tool edge.
func (oq *OperativerecordQuery) QueryTool() *ToolQuery {
	query := &ToolQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operativerecord.Table, operativerecord.FieldID, oq.sqlQuery()),
			sqlgraph.To(tool.Table, tool.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operativerecord.ToolTable, operativerecord.ToolColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Operativerecord entity in the query. Returns *NotFoundError when no operativerecord was found.
func (oq *OperativerecordQuery) First(ctx context.Context) (*Operativerecord, error) {
	os, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(os) == 0 {
		return nil, &NotFoundError{operativerecord.Label}
	}
	return os[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OperativerecordQuery) FirstX(ctx context.Context) *Operativerecord {
	o, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return o
}

// FirstID returns the first Operativerecord id in the query. Returns *NotFoundError when no id was found.
func (oq *OperativerecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{operativerecord.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (oq *OperativerecordQuery) FirstXID(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Operativerecord entity in the query, returns an error if not exactly one entity was returned.
func (oq *OperativerecordQuery) Only(ctx context.Context) (*Operativerecord, error) {
	os, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(os) {
	case 1:
		return os[0], nil
	case 0:
		return nil, &NotFoundError{operativerecord.Label}
	default:
		return nil, &NotSingularError{operativerecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OperativerecordQuery) OnlyX(ctx context.Context) *Operativerecord {
	o, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// OnlyID returns the only Operativerecord id in the query, returns an error if not exactly one id was returned.
func (oq *OperativerecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = &NotSingularError{operativerecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OperativerecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Operativerecords.
func (oq *OperativerecordQuery) All(ctx context.Context) ([]*Operativerecord, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *OperativerecordQuery) AllX(ctx context.Context) []*Operativerecord {
	os, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return os
}

// IDs executes the query and returns a list of Operativerecord ids.
func (oq *OperativerecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(operativerecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OperativerecordQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OperativerecordQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OperativerecordQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OperativerecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OperativerecordQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OperativerecordQuery) Clone() *OperativerecordQuery {
	return &OperativerecordQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		unique:     append([]string{}, oq.unique...),
		predicates: append([]predicate.Operativerecord{}, oq.predicates...),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

//  WithExaminationroom tells the query-builder to eager-loads the nodes that are connected to
// the "examinationroom" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OperativerecordQuery) WithExaminationroom(opts ...func(*ExaminationroomQuery)) *OperativerecordQuery {
	query := &ExaminationroomQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withExaminationroom = query
	return oq
}

//  WithNurse tells the query-builder to eager-loads the nodes that are connected to
// the "nurse" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OperativerecordQuery) WithNurse(opts ...func(*NurseQuery)) *OperativerecordQuery {
	query := &NurseQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withNurse = query
	return oq
}

//  WithOperative tells the query-builder to eager-loads the nodes that are connected to
// the "operative" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OperativerecordQuery) WithOperative(opts ...func(*OperativeQuery)) *OperativerecordQuery {
	query := &OperativeQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withOperative = query
	return oq
}

//  WithTool tells the query-builder to eager-loads the nodes that are connected to
// the "tool" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OperativerecordQuery) WithTool(opts ...func(*ToolQuery)) *OperativerecordQuery {
	query := &ToolQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withTool = query
	return oq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OperativeTime time.Time `json:"OperativeTime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Operativerecord.Query().
//		GroupBy(operativerecord.FieldOperativeTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oq *OperativerecordQuery) GroupBy(field string, fields ...string) *OperativerecordGroupBy {
	group := &OperativerecordGroupBy{config: oq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		OperativeTime time.Time `json:"OperativeTime,omitempty"`
//	}
//
//	client.Operativerecord.Query().
//		Select(operativerecord.FieldOperativeTime).
//		Scan(ctx, &v)
//
func (oq *OperativerecordQuery) Select(field string, fields ...string) *OperativerecordSelect {
	selector := &OperativerecordSelect{config: oq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(), nil
	}
	return selector
}

func (oq *OperativerecordQuery) prepareQuery(ctx context.Context) error {
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OperativerecordQuery) sqlAll(ctx context.Context) ([]*Operativerecord, error) {
	var (
		nodes       = []*Operativerecord{}
		withFKs     = oq.withFKs
		_spec       = oq.querySpec()
		loadedTypes = [4]bool{
			oq.withExaminationroom != nil,
			oq.withNurse != nil,
			oq.withOperative != nil,
			oq.withTool != nil,
		}
	)
	if oq.withExaminationroom != nil || oq.withNurse != nil || oq.withOperative != nil || oq.withTool != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, operativerecord.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Operativerecord{config: oq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oq.withExaminationroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Operativerecord)
		for i := range nodes {
			if fk := nodes[i].examinationroom_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(examinationroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "examinationroom_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Examinationroom = n
			}
		}
	}

	if query := oq.withNurse; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Operativerecord)
		for i := range nodes {
			if fk := nodes[i].nurse_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(nurse.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "nurse_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Nurse = n
			}
		}
	}

	if query := oq.withOperative; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Operativerecord)
		for i := range nodes {
			if fk := nodes[i].operative_type; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(operative.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "operative_type" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Operative = n
			}
		}
	}

	if query := oq.withTool; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Operativerecord)
		for i := range nodes {
			if fk := nodes[i].tool_name; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(tool.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "tool_name" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tool = n
			}
		}
	}

	return nodes, nil
}

func (oq *OperativerecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OperativerecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (oq *OperativerecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   operativerecord.Table,
			Columns: operativerecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: operativerecord.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OperativerecordQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(operativerecord.Table)
	selector := builder.Select(t1.Columns(operativerecord.Columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(operativerecord.Columns...)...)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OperativerecordGroupBy is the builder for group-by Operativerecord entities.
type OperativerecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OperativerecordGroupBy) Aggregate(fns ...AggregateFunc) *OperativerecordGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scan the result into the given value.
func (ogb *OperativerecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OperativerecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) StringX(ctx context.Context) string {
	v, err := ogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OperativerecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) IntX(ctx context.Context) int {
	v, err := ogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OperativerecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OperativerecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ogb *OperativerecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogb *OperativerecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := ogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogb *OperativerecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ogb.sqlQuery().Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OperativerecordGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql
	columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
	columns = append(columns, ogb.fields...)
	for _, fn := range ogb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ogb.fields...)
}

// OperativerecordSelect is the builder for select fields of Operativerecord entities.
type OperativerecordSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (os *OperativerecordSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := os.path(ctx)
	if err != nil {
		return err
	}
	os.sql = query
	return os.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (os *OperativerecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := os.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OperativerecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (os *OperativerecordSelect) StringsX(ctx context.Context) []string {
	v, err := os.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = os.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (os *OperativerecordSelect) StringX(ctx context.Context) string {
	v, err := os.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OperativerecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (os *OperativerecordSelect) IntsX(ctx context.Context) []int {
	v, err := os.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = os.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (os *OperativerecordSelect) IntX(ctx context.Context) int {
	v, err := os.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OperativerecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (os *OperativerecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := os.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = os.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (os *OperativerecordSelect) Float64X(ctx context.Context) float64 {
	v, err := os.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OperativerecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (os *OperativerecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := os.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (os *OperativerecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = os.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{operativerecord.Label}
	default:
		err = fmt.Errorf("ent: OperativerecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (os *OperativerecordSelect) BoolX(ctx context.Context) bool {
	v, err := os.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (os *OperativerecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := os.sqlQuery().Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (os *OperativerecordSelect) sqlQuery() sql.Querier {
	selector := os.sql
	selector.Select(selector.Columns(os.fields...)...)
	return selector
}
