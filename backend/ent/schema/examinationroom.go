package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Examinationroom holds the schema definition for the Examinationroom entity.
type Examinationroom struct {
	ent.Schema
}

// Fields of the Examinationroom.
func (Examinationroom) Fields() []ent.Field {
	return []ent.Field{
		field.String("examinationroomName").NotEmpty(),
	}
}

// Edges of the Examinationroom.
func (Examinationroom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fromexaminationroom", Operativerecord.Type).StorageKey(edge.Column("examinationroom_id")),
	}
}
