// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/B5871803/app/ent/examinationroom"
	"github.com/B5871803/app/ent/nurse"
	"github.com/B5871803/app/ent/operative"
	"github.com/B5871803/app/ent/operativerecord"
	"github.com/B5871803/app/ent/tool"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Operativerecord is the model entity for the Operativerecord schema.
type Operativerecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OperativeTime holds the value of the "OperativeTime" field.
	OperativeTime time.Time `json:"OperativeTime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OperativerecordQuery when eager-loading is set.
	Edges              OperativerecordEdges `json:"edges"`
	examinationroom_id *int
	nurse_id           *int
	operative_type     *int
	tool_name          *int
}

// OperativerecordEdges holds the relations/edges for other nodes in the graph.
type OperativerecordEdges struct {
	// Examinationroom holds the value of the examinationroom edge.
	Examinationroom *Examinationroom
	// Nurse holds the value of the nurse edge.
	Nurse *Nurse
	// Operative holds the value of the operative edge.
	Operative *Operative
	// Tool holds the value of the tool edge.
	Tool *Tool
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ExaminationroomOrErr returns the Examinationroom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperativerecordEdges) ExaminationroomOrErr() (*Examinationroom, error) {
	if e.loadedTypes[0] {
		if e.Examinationroom == nil {
			// The edge examinationroom was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: examinationroom.Label}
		}
		return e.Examinationroom, nil
	}
	return nil, &NotLoadedError{edge: "examinationroom"}
}

// NurseOrErr returns the Nurse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperativerecordEdges) NurseOrErr() (*Nurse, error) {
	if e.loadedTypes[1] {
		if e.Nurse == nil {
			// The edge nurse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: nurse.Label}
		}
		return e.Nurse, nil
	}
	return nil, &NotLoadedError{edge: "nurse"}
}

// OperativeOrErr returns the Operative value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperativerecordEdges) OperativeOrErr() (*Operative, error) {
	if e.loadedTypes[2] {
		if e.Operative == nil {
			// The edge operative was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: operative.Label}
		}
		return e.Operative, nil
	}
	return nil, &NotLoadedError{edge: "operative"}
}

// ToolOrErr returns the Tool value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OperativerecordEdges) ToolOrErr() (*Tool, error) {
	if e.loadedTypes[3] {
		if e.Tool == nil {
			// The edge tool was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tool.Label}
		}
		return e.Tool, nil
	}
	return nil, &NotLoadedError{edge: "tool"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Operativerecord) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // OperativeTime
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Operativerecord) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // examinationroom_id
		&sql.NullInt64{}, // nurse_id
		&sql.NullInt64{}, // operative_type
		&sql.NullInt64{}, // tool_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Operativerecord fields.
func (o *Operativerecord) assignValues(values ...interface{}) error {
	if m, n := len(values), len(operativerecord.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field OperativeTime", values[0])
	} else if value.Valid {
		o.OperativeTime = value.Time
	}
	values = values[1:]
	if len(values) == len(operativerecord.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field examinationroom_id", value)
		} else if value.Valid {
			o.examinationroom_id = new(int)
			*o.examinationroom_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field nurse_id", value)
		} else if value.Valid {
			o.nurse_id = new(int)
			*o.nurse_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field operative_type", value)
		} else if value.Valid {
			o.operative_type = new(int)
			*o.operative_type = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field tool_name", value)
		} else if value.Valid {
			o.tool_name = new(int)
			*o.tool_name = int(value.Int64)
		}
	}
	return nil
}

// QueryExaminationroom queries the examinationroom edge of the Operativerecord.
func (o *Operativerecord) QueryExaminationroom() *ExaminationroomQuery {
	return (&OperativerecordClient{config: o.config}).QueryExaminationroom(o)
}

// QueryNurse queries the nurse edge of the Operativerecord.
func (o *Operativerecord) QueryNurse() *NurseQuery {
	return (&OperativerecordClient{config: o.config}).QueryNurse(o)
}

// QueryOperative queries the operative edge of the Operativerecord.
func (o *Operativerecord) QueryOperative() *OperativeQuery {
	return (&OperativerecordClient{config: o.config}).QueryOperative(o)
}

// QueryTool queries the tool edge of the Operativerecord.
func (o *Operativerecord) QueryTool() *ToolQuery {
	return (&OperativerecordClient{config: o.config}).QueryTool(o)
}

// Update returns a builder for updating this Operativerecord.
// Note that, you need to call Operativerecord.Unwrap() before calling this method, if this Operativerecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Operativerecord) Update() *OperativerecordUpdateOne {
	return (&OperativerecordClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Operativerecord) Unwrap() *Operativerecord {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Operativerecord is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Operativerecord) String() string {
	var builder strings.Builder
	builder.WriteString("Operativerecord(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", OperativeTime=")
	builder.WriteString(o.OperativeTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Operativerecords is a parsable slice of Operativerecord.
type Operativerecords []*Operativerecord

func (o Operativerecords) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}