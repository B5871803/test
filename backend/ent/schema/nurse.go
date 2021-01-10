package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Nurse holds the schema definition for the Nurse entity.
type Nurse struct {
	ent.Schema
}

// Fields of the Nurse.
func (Nurse) Fields() []ent.Field {
	return []ent.Field{
		field.String("nurse_name").NotEmpty(),
		field.String("nurse_email").NotEmpty(),
		field.String("nurse_password").NotEmpty(),
		field.String("nurse_tel").NotEmpty(),
	}
}

// Edges of the Nurse.
func (Nurse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nurseop", Operativerecord.Type).StorageKey(edge.Column("nurse_id")),
	}
}
